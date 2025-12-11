package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/config"
)

type PandaPayCallback struct {
	service.Service
}

// HandleCallback 处理PandaPay回调
func (e *PandaPayCallback) HandleCallback(c *dto.PandaPayCallbackReq) error {
	// 1. 验证签名
	if err := e.verifySign(c); err != nil {
		e.Log.Errorf("PandaPayCallback verify sign failed: %s", err)
		return err
	}

	// 2. 查询提现订单
	var withdrawal models.HsUserWithdrawal
	err := e.Orm.Where("withdraw_no = ?", c.MerchantOrderId).First(&withdrawal).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("PandaPayCallback withdrawal not found: %s", c.MerchantOrderId)
			return errors.New("提现订单不存在")
		}
		e.Log.Errorf("PandaPayCallback query withdrawal error: %s", err)
		return errors.New("查询提现订单失败")
	}

	// 3. 检查订单状态，避免重复处理
	if withdrawal.Status == "success" || withdrawal.Status == "failed" {
		e.Log.Infof("PandaPayCallback withdrawal already processed: %s, status: %s", c.MerchantOrderId, withdrawal.Status)
		return nil // 已处理的订单直接返回成功
	}

	// 4. 根据回调状态更新订单
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		updates := make(map[string]interface{})

		// 状态映射: 0=等待, 1=成功, 2=失败, 3=下单失败, 4=处理中
		switch c.Status {
		case 1: // 成功
			updates["status"] = "success"
			updates["processed_at"] = time.Now()
			updates["channel_txn_id"] = c.SessionId

		case 2, 3: // 失败或下单失败
			updates["status"] = "failed"
			updates["processed_at"] = time.Now()
			updates["reason"] = c.Msg
			if c.SessionId != "" {
				updates["channel_txn_id"] = c.SessionId
			}

			// 失败需要退款：将扣款金额退回用户账户
			err := e.refundToUser(tx, &withdrawal, c.MerchantOrderId)
			if err != nil {
				e.Log.Errorf("PandaPayCallback refund to user error: %s", err)
				return err
			}

		case 4: // 处理中
			updates["status"] = "processing"

		case 0: // 等待
			updates["status"] = "pending"

		default:
			e.Log.Errorf("PandaPayCallback unknown status: %d", c.Status)
			return fmt.Errorf("未知的订单状态: %d", c.Status)
		}

		// 更新提现订单
		err := tx.Model(&models.HsUserWithdrawal{}).
			Where("id = ?", withdrawal.Id).
			Updates(updates).Error
		if err != nil {
			e.Log.Errorf("PandaPayCallback update withdrawal error: %s", err)
			return errors.New("更新提现订单失败")
		}

		e.Log.Infof("PandaPayCallback processed successfully: withdrawNo=%s, status=%d->%s, pandaOrderId=%d",
			c.MerchantOrderId, c.Status, updates["status"], c.OrderId)

		return nil
	})
}

// verifySign 验证签名
// 签名格式: merchantId=1001&merchantOrderId=1000&amount=100.00&appSecret=商户appSecret
// MD5加密并转小写
func (e *PandaPayCallback) verifySign(c *dto.PandaPayCallbackReq) error {
	// 获取配置的appSecret
	appSecret := config.ExtConfig.PandaPay.AppSecret
	if appSecret == "" {
		return errors.New("PandaPay AppSecret未配置")
	}

	// 构建签名字符串（注意：金额必须保留两位小数）
	signStr := fmt.Sprintf("merchantId=%d&merchantOrderId=%s&amount=%s&appSecret=%s",
		c.MerchantId, c.MerchantOrderId, c.Amount, appSecret)

	// MD5加密并转小写
	hash := md5.Sum([]byte(signStr))
	expectedSign := fmt.Sprintf("%x", hash)

	// 比较签名
	if expectedSign != c.Sign {
		e.Log.Errorf("PandaPayCallback sign mismatch: expected=%s, got=%s, signStr=%s",
			expectedSign, c.Sign, signStr)
		return errors.New("签名验证失败")
	}

	e.Log.Infof("PandaPayCallback sign verified: %s", c.Sign)
	return nil
}

// refundToUser 退款给用户（提现失败时）
func (e *PandaPayCallback) refundToUser(tx *gorm.DB, withdrawal *models.HsUserWithdrawal, withdrawNo string) error {
	// 1. 解析提现金额（包含手续费，需要全额退回）
	amount, err := strconv.ParseFloat(withdrawal.Amount, 64)
	if err != nil {
		e.Log.Errorf("refundToUser parse amount error: %s", err)
		return errors.New("解析提现金额失败")
	}

	fee, _ := strconv.ParseFloat(withdrawal.Fee, 64)
	totalRefund := amount + fee // 退回金额 = 提现金额 + 手续费

	userId, _ := strconv.Atoi(withdrawal.UserId)
	currencyCode := withdrawal.CurrencyCode

	// 2. 查询用户当前余额和版本号
	var user models.HsUsers
	err = tx.Select("id, balance, crypto_balance, version").Where("id = ?", userId).First(&user).Error
	if err != nil {
		e.Log.Errorf("refundToUser get user error: %s", err)
		return errors.New("获取用户信息失败")
	}

	// 3. 根据提现方式选择余额字段
	var balanceField string
	var balanceBefore float64
	var decimalPlaces int
	isCrypto := withdrawal.Method == "crypto"

	if isCrypto {
		balanceField = "crypto_balance"
		balanceBefore, _ = strconv.ParseFloat(user.CryptoBalance, 64)
		decimalPlaces = 8
	} else {
		balanceField = "balance"
		balanceBefore, _ = strconv.ParseFloat(user.Balance, 64)
		decimalPlaces = 2
	}

	balanceAfter := balanceBefore + totalRefund

	// 4. 更新用户余额（使用乐观锁）
	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", userId, user.Version).
		Updates(map[string]interface{}{
			balanceField: fmt.Sprintf("%.*f", decimalPlaces, balanceAfter),
			"version":    gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		e.Log.Errorf("refundToUser update balance error: %s", result.Error)
		return errors.New("更新用户余额失败")
	}

	if result.RowsAffected == 0 {
		e.Log.Errorf("refundToUser update balance conflict for user %d", userId)
		return errors.New("余额更新冲突，请重试")
	}

	// 5. 创建流水记录
	ledger := models.HsUserLedger{
		UserId:         withdrawal.UserId,
		CurrencyCode:   currencyCode,
		Direction:      "1", // 1=入账
		Amount:         fmt.Sprintf("%.*f", decimalPlaces, totalRefund),
		BalanceBefore:  fmt.Sprintf("%.*f", decimalPlaces, balanceBefore),
		BalanceAfter:   fmt.Sprintf("%.*f", decimalPlaces, balanceAfter),
		BizType:        "withdrawal_refund",
		BizId:          strconv.Itoa(withdrawal.Id),
		IdempotencyKey: fmt.Sprintf("WITHDRAWAL_REFUND:%s:%d", withdrawNo, time.Now().UnixNano()),
		RefTable:       "hs_user_withdrawal",
		RefId:          strconv.Itoa(withdrawal.Id),
		Remark:         fmt.Sprintf("提现失败退款，提现单号: %s", withdrawNo),
		Status:         "1", // 1=已入账
	}

	err = tx.Create(&ledger).Error
	if err != nil {
		e.Log.Errorf("refundToUser create ledger error: %s", err)
		return errors.New("创建流水记录失败")
	}

	e.Log.Infof("Refund to user %d: amount=%.2f %s, balance: %.2f -> %.2f",
		userId, totalRefund, currencyCode, balanceBefore, balanceAfter)

	return nil
}
