package service

import (
	"encoding/json"
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
	"go-admin/common/pandapay"
)

type HsUserWithdrawal struct {
	service.Service
}

// buildReasonJSON 将原因字符串转换为JSON格式
func buildReasonJSON(message string, extras ...map[string]interface{}) string {
	data := map[string]interface{}{
		"message": message,
		"time":    time.Now().Format("2006-01-02 15:04:05"),
	}
	// 合并额外的字段
	for _, extra := range extras {
		for k, v := range extra {
			data[k] = v
		}
	}
	jsonBytes, _ := json.Marshal(data)
	return string(jsonBytes)
}

// GetPage 获取HsUserWithdrawal列表
func (e *HsUserWithdrawal) GetPage(c *dto.HsUserWithdrawalGetPageReq, p *actions.DataPermission, list *[]models.HsUserWithdrawal, count *int64) error {
	var err error
	var data models.HsUserWithdrawal

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).Order("created_at DESC").
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawalService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetPageWithStats 获取带统计信息的HsUserWithdrawal列表
func (e *HsUserWithdrawal) GetPageWithStats(c *dto.HsUserWithdrawalGetPageReq, p *actions.DataPermission, list *[]dto.HsUserWithdrawalWithStats, count *int64) error {
	var err error
	var data models.HsUserWithdrawal
	var withdrawalList []models.HsUserWithdrawal

	// 先获取基础的提现记录列表
	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Order("created_at DESC").Find(&withdrawalList).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawalService GetPageWithStats error:%s \r\n", err)
		return err
	}

	// 为每个提现记录计算统计信息
	for _, withdrawal := range withdrawalList {
		statsItem := dto.HsUserWithdrawalWithStats{
			HsUserWithdrawal: withdrawal,
		}

		// 计算累计提现金额和笔数（只统计成功的提现）
		var withdrawStats struct {
			TotalAmount string
			Count       int64
		}
		err = e.Orm.Model(&models.HsUserWithdrawal{}).
			Select("COALESCE(SUM(CAST(amount AS DECIMAL(18,2))), 0) as total_amount, COUNT(*) as count").
			Where("user_id = ? AND status = 'success'", withdrawal.UserId).
			Scan(&withdrawStats).Error
		if err != nil {
			e.Log.Errorf("Failed to calculate withdrawal stats for user %s: %s", withdrawal.UserId, err)
			withdrawStats.TotalAmount = "0.00"
			withdrawStats.Count = 0
		}

		// 计算累计售卖金额和笔数（只统计已完成的订单）
		var orderStats struct {
			TotalAmount string
			Count       int64
		}
		err = e.Orm.Model(&models.OrdUserOrders{}).
			Select("COALESCE(SUM(CAST(balance AS DECIMAL(20,8))), 0) as total_amount, COUNT(*) as count").
			Where("user_id = ? AND status = 3", withdrawal.UserId). // status=3表示已完成
			Scan(&orderStats).Error
		if err != nil {
			e.Log.Errorf("Failed to calculate order stats for user %s: %s", withdrawal.UserId, err)
			orderStats.TotalAmount = "0.00"
			orderStats.Count = 0
		}

		statsItem.TotalWithdrawn = withdrawStats.TotalAmount
		statsItem.WithdrawCount = withdrawStats.Count
		statsItem.TotalSales = orderStats.TotalAmount
		statsItem.OrderCount = orderStats.Count

		*list = append(*list, statsItem)
	}

	return nil
}

// Get 获取HsUserWithdrawal对象
func (e *HsUserWithdrawal) Get(d *dto.HsUserWithdrawalGetReq, p *actions.DataPermission, model *models.HsUserWithdrawal) error {
	var data models.HsUserWithdrawal

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserWithdrawal error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserWithdrawal对象
func (e *HsUserWithdrawal) Insert(c *dto.HsUserWithdrawalInsertReq) error {
    var err error
    var data models.HsUserWithdrawal
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawalService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserWithdrawal对象
func (e *HsUserWithdrawal) Update(c *dto.HsUserWithdrawalUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserWithdrawal{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserWithdrawalService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserWithdrawal
func (e *HsUserWithdrawal) Remove(d *dto.HsUserWithdrawalDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserWithdrawal

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveHsUserWithdrawal error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// BankAccountInfo 银行账户信息结构（从accountInfo JSON解析）
// 支持新格式：{"bank_id": 1, "bank_name": "Bank of America", "card_type": 1, "branch_name": "", "card_number": "****3659", "card_holder_name": "wwww"}
type BankAccountInfo struct {
	// 新格式字段
	BankId         int    `json:"bank_id"`
	BankName       string `json:"bank_name"`
	CardType       int    `json:"card_type"`
	BranchName     string `json:"branch_name"`
	CardNumber     string `json:"card_number"`
	CardHolderName string `json:"card_holder_name"`

	// 旧格式字段（兼容）
	AccountName   string `json:"accountName"`
	BankCode      string `json:"bankCode"`
	AccountNumber string `json:"accountNumber"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
	Address       string `json:"address"`
}

// WithdrawalPayoutResult 提现转账结果
type WithdrawalPayoutResult struct {
	ChannelTxnId string // 通道交易ID
	Status       string // 提现状态：processing/success/failed
	Reason       string // 失败原因
}

// SubmitWithdrawalPayout 提交提现转账（公共方法，供approve和自动处理器使用）
func (e *HsUserWithdrawal) SubmitWithdrawalPayout(withdrawal *models.HsUserWithdrawal) (*WithdrawalPayoutResult, error) {
	result := &WithdrawalPayoutResult{}

	// 根据提现方式处理
	if withdrawal.Method == "bank" {
		// 银行卡提现，调用PandaPay代付接口
		e.Log.Infof("开始处理银行卡提现，提现单号: %s", withdrawal.WithdrawNo)

		// 解析账户信息
		var accountInfo BankAccountInfo
		if err := json.Unmarshal([]byte(withdrawal.AccountInfo), &accountInfo); err != nil {
			return nil, fmt.Errorf("解析账户信息失败: %w", err)
		}

		// 判断是新格式还是旧格式，并统一处理
		var accountName, bankCode, accountNumber, email, mobile, address string

		if accountInfo.BankId > 0 {
			// 新格式：{"bank_id": 1, "bank_name": "...", "card_holder_name": "...", "card_number": "..."}
			e.Log.Infof("使用新格式账户信息，BankId: %d", accountInfo.BankId)

			// 根据bank_id查询银行编码
			var bank models.HsBanks
			err := e.Orm.Where("id = ? AND status = 1", accountInfo.BankId).First(&bank).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, fmt.Errorf("银行信息不存在或已禁用，BankId: %d", accountInfo.BankId)
				}
				e.Log.Errorf("查询银行信息失败: %s", err)
				return nil, fmt.Errorf("查询银行信息失败: %w", err)
			}

			// 映射新格式字段
			accountName = accountInfo.CardHolderName
			bankCode = bank.BankCode
			accountNumber = accountInfo.CardNumber
			email = ""    // 新格式没有email
			mobile = ""   // 新格式没有mobile
			address = ""  // 可以使用分行名称
			if accountInfo.BranchName != "" {
				address = accountInfo.BranchName
			}

			e.Log.Infof("银行信息: BankName=%s, BankCode=%s, CardHolder=%s", bank.BankName, bankCode, accountName)
		} else {
			// 旧格式：{"accountName": "...", "bankCode": "...", "accountNumber": "..."}
			e.Log.Infof("使用旧格式账户信息")
			accountName = accountInfo.AccountName
			bankCode = accountInfo.BankCode
			accountNumber = accountInfo.AccountNumber
			email = accountInfo.Email
			mobile = accountInfo.Mobile
			address = accountInfo.Address
		}

		// 验证必填字段
		if accountName == "" || bankCode == "" || accountNumber == "" {
			return nil, errors.New("银行账户信息不完整：缺少持卡人姓名、银行编码或卡号")
		}

		// 解析金额
		amount, err := pandapay.ParseAmount(withdrawal.NetAmount)
		if err != nil {
			return nil, fmt.Errorf("金额解析失败: %w", err)
		}

		// 构建PandaPay请求
		client := pandapay.NewClient()
		payoutReq := pandapay.BuildBankPayoutRequest(
			withdrawal.WithdrawNo,
			amount,
			withdrawal.CurrencyCode,
			accountName,
			bankCode,
			accountNumber,
			email,
			mobile,
			address,
		)

		// 调用PandaPay代付接口
		e.Log.Infof("调用PandaPay代付接口，订单号: %s, 金额: %.2f", withdrawal.WithdrawNo, amount)
		resp, err := client.SubmitPayout(payoutReq)
		if err != nil {
			e.Log.Errorf("PandaPay代付失败: %s", err)
			result.Status = "failed"
			result.Reason = buildReasonJSON(fmt.Sprintf("代付失败: %s", err.Error()))
			return result, err
		}

		// 代付成功
		result.ChannelTxnId = resp.Data.ChannelTxnId
		e.Log.Infof("PandaPay代付成功，通道流水号: %s, 状态: %d", result.ChannelTxnId, resp.Data.Status)

		// 根据PandaPay返回的状态设置提现状态
		// 0=等待, 1=成功, 2=失败, 4=处理中
		switch resp.Data.Status {
		case 1:
			result.Status = "success"
		case 0, 4:
			result.Status = "processing"
		case 2, 3:
			result.Status = "failed"
			result.Reason = buildReasonJSON(fmt.Sprintf("PandaPay返回失败状态: %d", resp.Data.Status))
		default:
			result.Status = "processing"
		}
	} else if withdrawal.Method == "crypto" {
		// 加密货币提现，不调用PandaPay，直接标记为成功
		e.Log.Infof("加密货币提现，提现单号: %s, 无需调用代付接口", withdrawal.WithdrawNo)
		result.Status = "success"
	} else {
		return nil, fmt.Errorf("不支持的提现方式: %s", withdrawal.Method)
	}

	return result, nil
}

// Approve 审核通过提现申请，并自动转账（bank方式调用PandaPay）
func (e *HsUserWithdrawal) Approve(c *dto.HsUserWithdrawalApproveReq, p *actions.DataPermission) error {
	var withdrawal models.HsUserWithdrawal

	// 查询提现记录
	err := e.Orm.Model(&withdrawal).
		Scopes(
			actions.Permission(withdrawal.TableName(), p),
		).
		First(&withdrawal, c.GetId()).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("提现记录不存在或无权操作")
		}
		e.Log.Errorf("查询提现记录失败: %s", err)
		return err
	}

	// 检查状态，只有review状态才能审核通过
	if withdrawal.Status != "review" && withdrawal.Status != "pending" {
		return fmt.Errorf("提现记录状态不正确，当前状态: %s，只能审核review或pending状态", withdrawal.Status)
	}

	// 调用公共转账方法
	result, payoutErr := e.SubmitWithdrawalPayout(&withdrawal)

	var channelTxnId string
	if result != nil {
		channelTxnId = result.ChannelTxnId
		withdrawal.Status = result.Status
		withdrawal.Reason = result.Reason
	} else {
		// 转账失败
		withdrawal.Status = "failed"
		if payoutErr != nil {
			withdrawal.Reason = buildReasonJSON(payoutErr.Error())
		}
	}

	// 更新提现记录
	withdrawal.ChannelTxnId = channelTxnId
	withdrawal.ProcessedAt = time.Now()
	withdrawal.UpdateBy = c.UpdateBy

	// 构建更新数据
	updateData := map[string]interface{}{
		"channel_txn_id": channelTxnId,
		"processed_at":   time.Now(),
		"update_by":      c.UpdateBy,
		"status":         withdrawal.Status,
	}

	// 只有在失败状态下才设置reason字段
	if withdrawal.Status == "failed" {
		reasonValue := withdrawal.Reason
		// 如果有备注，添加到reason字段
		if c.Remark != "" {
			if reasonValue != "" {
				// 尝试解析现有的JSON，添加备注
				var existingData map[string]interface{}
				if err := json.Unmarshal([]byte(reasonValue), &existingData); err == nil {
					existingData["remark"] = c.Remark
					jsonBytes, _ := json.Marshal(existingData)
					reasonValue = string(jsonBytes)
				} else {
					// 如果不是有效JSON，创建新的JSON
					reasonValue = buildReasonJSON(reasonValue, map[string]interface{}{"remark": c.Remark})
				}
			} else {
				reasonValue = buildReasonJSON(c.Remark)
			}
		}
		updateData["reason"] = reasonValue
	}
	// 非失败状态不更新 reason 字段，保持数据库中的原值

	// 保存更新
	db := e.Orm.Model(&withdrawal).Where("id = ?", withdrawal.Id).Updates(updateData)
	if err := db.Error; err != nil {
		e.Log.Errorf("更新提现记录失败: %s", err)
		return fmt.Errorf("更新提现记录失败: %w", err)
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	// 如果代付失败，返回错误
	if payoutErr != nil {
		return fmt.Errorf("审核通过但代付失败: %w", payoutErr)
	}

	return nil
}

// Reject 拒绝提现申请
func (e *HsUserWithdrawal) Reject(c *dto.HsUserWithdrawalRejectReq, p *actions.DataPermission) error {
	var withdrawal models.HsUserWithdrawal

	// 查询提现记录
	err := e.Orm.Model(&withdrawal).
		Scopes(
			actions.Permission(withdrawal.TableName(), p),
		).
		First(&withdrawal, c.GetId()).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("提现记录不存在或无权操作")
		}
		e.Log.Errorf("查询提现记录失败: %s", err)
		return err
	}

	// 检查状态
	if withdrawal.Status != "review" && withdrawal.Status != "pending" {
		return fmt.Errorf("提现记录状态不正确，当前状态: %s，只能拒绝review或pending状态", withdrawal.Status)
	}

	// 更新为失败状态
	withdrawal.Status = "failed"
	withdrawal.Reason = buildReasonJSON(c.Reason, map[string]interface{}{"type": "rejected"})
	withdrawal.ProcessedAt = time.Now()
	withdrawal.UpdateBy = c.UpdateBy

	// 保存更新
	db := e.Orm.Save(&withdrawal)
	if err := db.Error; err != nil {
		e.Log.Errorf("更新提现记录失败: %s", err)
		return fmt.Errorf("更新提现记录失败: %w", err)
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	return nil
}

// ManualTransfer 手动处理提现转账
func (e *HsUserWithdrawal) ManualTransfer(c *dto.HsUserWithdrawalManualTransferReq, p *actions.DataPermission) error {
	var withdrawal models.HsUserWithdrawal

	// 查询提现记录
	err := e.Orm.Model(&withdrawal).
		Scopes(
			actions.Permission(withdrawal.TableName(), p),
		).
		First(&withdrawal, c.GetId()).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("提现记录不存在或无权操作")
		}
		e.Log.Errorf("查询提现记录失败: %s", err)
		return err
	}

	// 检查状态，只有 review 或 processing 状态才能手动处理
	if withdrawal.Status != "review" && withdrawal.Status != "processing" {
		return fmt.Errorf("提现记录状态不正确，当前状态: %s，只能手动处理review或processing状态", withdrawal.Status)
	}

	// 根据是否成功设置状态
	if c.Success {
		withdrawal.Status = "success"

		// 手动处理成功，更新用户的总提现金额
		userId, _ := strconv.Atoi(withdrawal.UserId)
		netAmount := withdrawal.NetAmount

		// 1. 查询用户当前版本号
		var user models.HsUsers
		if err := e.Orm.Select("id, version").Where("id = ?", userId).First(&user).Error; err != nil {
			e.Log.Errorf("ManualTransfer get user error: %s", err)
			return errors.New("获取用户信息失败")
		}

		// 2. 根据提现方式选择更新的字段
		var updateField string
		if withdrawal.Method == "crypto" {
			updateField = "total_withdraw_crypto"
		} else {
			updateField = "total_withdraw"
		}

		// 3. 使用乐观锁更新
		result := e.Orm.Model(&models.HsUsers{}).
			Where("id = ? AND version = ?", userId, user.Version).
			Updates(map[string]interface{}{
				updateField: gorm.Expr("COALESCE("+updateField+", 0) + ?", netAmount),
				"version":   gorm.Expr("version + 1"),
			})
		if result.Error != nil {
			e.Log.Errorf("ManualTransfer update %s error: %s", updateField, result.Error)
			return errors.New("更新用户总提现金额失败")
		}
		if result.RowsAffected == 0 {
			e.Log.Errorf("ManualTransfer version conflict for user %d", userId)
			return errors.New("用户数据更新冲突，请重试")
		}
		e.Log.Infof("Updated user %d %s: +%s", userId, updateField, netAmount)
	} else {
		withdrawal.Status = "failed"
	}

	// 构建备注信息
	reasonData := map[string]interface{}{
		"manual_transfer": true,
		"remark":          c.Remark,
		"time":            time.Now().Format("2006-01-02 15:04:05"),
	}
	reasonJSON, _ := json.Marshal(reasonData)

	withdrawal.Reason = string(reasonJSON)
	withdrawal.TransferImage = c.TransferImage
	withdrawal.ProcessedAt = time.Now()
	withdrawal.UpdateBy = c.UpdateBy

	// 保存更新
	db := e.Orm.Save(&withdrawal)
	if err := db.Error; err != nil {
		e.Log.Errorf("更新提现记录失败: %s", err)
		return fmt.Errorf("更新提现记录失败: %w", err)
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	e.Log.Infof("手动处理提现单 %s，结果: %v", withdrawal.WithdrawNo, c.Success)
	return nil
}
