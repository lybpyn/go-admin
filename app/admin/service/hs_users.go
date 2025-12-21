package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type HsUsers struct {
	service.Service
}

// GetPage 获取HsUsers列表
func (e *HsUsers) GetPage(c *dto.HsUsersGetPageReq, p *actions.DataPermission, list *[]models.HsUsers, count *int64) error {
	var err error
	var data models.HsUsers

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUsersService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUsers对象
func (e *HsUsers) Get(d *dto.HsUsersGetReq, p *actions.DataPermission, model *models.HsUsers) error {
	var data models.HsUsers

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUsers error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUsers对象
func (e *HsUsers) Insert(c *dto.HsUsersInsertReq) error {
    var err error
    var data models.HsUsers
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUsersService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUsers对象
func (e *HsUsers) Update(c *dto.HsUsersUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUsers{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUsersService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUsers
func (e *HsUsers) Remove(d *dto.HsUsersDeleteReq, p *actions.DataPermission) error {
	var data models.HsUsers

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUsers error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

// AdjustBalance 调整用户余额并记录流水
func (e *HsUsers) AdjustBalance(c *dto.HsUsersAdjustBalanceReq) error {
	// 解析调整金额
	adjustAmount, err := strconv.ParseFloat(c.Amount, 64)
	if err != nil {
		return fmt.Errorf("金额格式错误: %w", err)
	}
	if adjustAmount == 0 {
		return errors.New("调整金额不能为0")
	}

	// 开启事务
	tx := e.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 查询用户
	var user models.HsUsers
	if err := tx.First(&user, c.UserId).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return fmt.Errorf("查询用户失败: %w", err)
	}

	// 根据余额类型获取当前余额并更新
	var balanceBefore, balanceAfter float64
	var balanceField, bizType, currencyCode string
	var precision int

	switch c.BalanceType {
	case "balance":
		balanceField = "balance"
		bizType = "manual_adjust_balance"
		currencyCode = user.CurrencyCode
		precision = 2
		balanceBefore, _ = strconv.ParseFloat(user.Balance, 64)
	case "crypto_balance":
		balanceField = "crypto_balance"
		bizType = "manual_adjust_crypto"
		currencyCode = "USDT"
		precision = 8
		balanceBefore, _ = strconv.ParseFloat(user.CryptoBalance, 64)
	case "frozen_balance":
		balanceField = "frozen_balance"
		bizType = "manual_adjust_frozen"
		currencyCode = user.CurrencyCode
		precision = 2
		balanceBefore, _ = strconv.ParseFloat(user.FrozenBalance, 64)
	case "crypto_frozen_balance":
		balanceField = "crypto_frozen_balance"
		bizType = "manual_adjust_crypto_frozen"
		currencyCode = "USDT"
		precision = 8
		balanceBefore, _ = strconv.ParseFloat(user.CryptoFrozenBalance, 64)
	default:
		tx.Rollback()
		return fmt.Errorf("不支持的余额类型: %s", c.BalanceType)
	}

	// 计算调整后余额
	balanceAfter = balanceBefore + adjustAmount
	if balanceAfter < 0 {
		tx.Rollback()
		return fmt.Errorf("余额不足，当前余额: %."+strconv.Itoa(precision)+"f，调整金额: %."+strconv.Itoa(precision)+"f", balanceBefore, adjustAmount)
	}

	// 更新用户余额
	var balanceAfterStr string
	if precision == 2 {
		balanceAfterStr = fmt.Sprintf("%.2f", balanceAfter)
	} else {
		balanceAfterStr = fmt.Sprintf("%.8f", balanceAfter)
	}

	if err := tx.Model(&models.HsUsers{}).Where("id = ?", c.UserId).
		Update(balanceField, balanceAfterStr).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新余额失败: %w", err)
	}

	// 记录流水
	direction := "1" // 入账
	amount := adjustAmount
	if adjustAmount < 0 {
		direction = "2" // 出账
		amount = -adjustAmount
	}

	var amountStr, balanceBeforeStr string
	if precision == 2 {
		amountStr = fmt.Sprintf("%.2f", amount)
		balanceBeforeStr = fmt.Sprintf("%.2f", balanceBefore)
	} else {
		amountStr = fmt.Sprintf("%.8f", amount)
		balanceBeforeStr = fmt.Sprintf("%.8f", balanceBefore)
	}

	ledger := models.HsUserLedger{
		UserId:         strconv.Itoa(c.UserId),
		CurrencyCode:   currencyCode,
		Direction:      direction,
		Amount:         amountStr,
		BalanceBefore:  balanceBeforeStr,
		BalanceAfter:   balanceAfterStr,
		BizType:        bizType,
		BizId:          fmt.Sprintf("MANUAL_%d_%d", c.UserId, time.Now().UnixNano()),
		IdempotencyKey: fmt.Sprintf("MANUAL_ADJUST:%d:%d", c.UserId, time.Now().UnixNano()),
		Remark:         c.Remark,
		Status:         "1",
	}
	ledger.CreateBy = c.CreateBy

	if err := tx.Create(&ledger).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("记录流水失败: %w", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}

	e.Log.Infof("调整用户 %d 余额成功，类型: %s，调整前: %s，调整后: %s，备注: %s",
		c.UserId, c.BalanceType, balanceBeforeStr, balanceAfterStr, c.Remark)

	return nil
}
