package jobs

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
)

// WithdrawalProcessor 提现自动处理器
type WithdrawalProcessor struct {
	db *gorm.DB
}

// NewWithdrawalProcessor 创建提现处理器
func NewWithdrawalProcessor() *WithdrawalProcessor {
	db := sdk.Runtime.GetDbByKey("*")
	return &WithdrawalProcessor{db: db}
}

// Start 启动定时任务（每5秒执行一次）
func (p *WithdrawalProcessor) Start() {
	log.Info("[WithdrawalProcessor] 提现自动处理器启动")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if err := p.Process(); err != nil {
			log.Errorf("[WithdrawalProcessor] 处理失败: %v", err)
		}
	}
}

// Process 处理待审核的提现申请
func (p *WithdrawalProcessor) Process() error {
	// 1. 查询所有 pending 状态的提现记录
	var withdrawals []models.HsUserWithdrawal
	err := p.db.Where("status = ?", "pending").Find(&withdrawals).Error
	if err != nil {
		return fmt.Errorf("查询待处理提现记录失败: %w", err)
	}

	if len(withdrawals) == 0 {
		return nil
	}

	log.Infof("[WithdrawalProcessor] 发现 %d 条待处理提现记录", len(withdrawals))

	// 2. 逐个处理
	for _, withdrawal := range withdrawals {
		if err := p.processOne(&withdrawal); err != nil {
			log.Errorf("[WithdrawalProcessor] 处理提现单 %s 失败: %v", withdrawal.WithdrawNo, err)
			continue
		}
	}

	return nil
}

// processOne 处理单个提现申请
func (p *WithdrawalProcessor) processOne(withdrawal *models.HsUserWithdrawal) error {
	// 1. 获取提现规则
	var rule models.HsConfigWithdrawRules
	err := p.db.Where("currency_code = ? AND is_active = ?", withdrawal.CurrencyCode, "1").
		First(&rule).Error
	if err != nil {
		log.Errorf("[WithdrawalProcessor] 获取提现规则失败: %v", err)
		// 如果没有找到规则，转人工审核
		return p.updateStatusToReview(withdrawal, "未找到提现规则")
	}

	// 2. 检查是否满足规则
	pass, reason, err := p.checkRules(withdrawal, &rule)
	if err != nil {
		log.Errorf("[WithdrawalProcessor] 检查规则失败: %v", err)
		return p.updateStatusToReview(withdrawal, "检查规则失败: "+err.Error())
	}
	if !pass {
		log.Infof("[WithdrawalProcessor] 提现单 %s 不满足自动处理条件: %s", withdrawal.WithdrawNo, reason)
		return p.updateStatusToReview(withdrawal, reason)
	}

	// 3. 满足规则，自动调用 PandaPay 转账
	log.Infof("[WithdrawalProcessor] 提现单 %s 满足自动处理条件，开始调用 PandaPay", withdrawal.WithdrawNo)
	return p.submitToPandaPay(withdrawal)
}

// checkRules 检查是否满足自动处理规则
func (p *WithdrawalProcessor) checkRules(withdrawal *models.HsUserWithdrawal, rule *models.HsConfigWithdrawRules) (bool, string, error) {
	// 1. 检查每日累计提现金额
	if rule.DailyLimitAmount != "" && rule.DailyLimitAmount != "0" {
		dailyLimit, err := strconv.ParseFloat(rule.DailyLimitAmount, 64)
		if err != nil {
			return false, "", fmt.Errorf("解析每日限额失败: %w", err)
		}

		// 查询今日该用户的累计提现金额（成功的）
		var todayTotal float64
		today := time.Now().Format("2006-01-02")
		if err := p.db.Model(&models.HsUserWithdrawal{}).
			Select("COALESCE(SUM(CAST(amount AS DECIMAL(18,2))), 0)").
			Where("user_id = ? AND currency_code = ?  AND DATE(requested_at) = ?",
				withdrawal.UserId, withdrawal.CurrencyCode, today).
			Scan(&todayTotal).Error; err != nil {
			return false, "", fmt.Errorf("查询今日提现总额失败: %w", err)
		}

		// 加上当前这笔提现金额
		currentAmount, err := strconv.ParseFloat(withdrawal.Amount, 64)
		if err != nil {
			return false, "", fmt.Errorf("解析提现金额失败: %w", err)
		}

		if todayTotal+currentAmount > dailyLimit {
			return false, fmt.Sprintf("超出每日提现限额：已提现 %.2f，限额 %.2f", todayTotal, dailyLimit), nil
		}
	}

	// 2. 检查每日提现次数
	if rule.DailyLimitCount != "" && rule.DailyLimitCount != "0" {
		dailyCount, err := strconv.Atoi(rule.DailyLimitCount)
		if err != nil {
			return false, "", fmt.Errorf("解析每日提现次数限制失败: %w", err)
		}

		// 查询今日该用户的提现次数
		var todayCount int64
		today := time.Now().Format("2006-01-02")
		if err := p.db.Model(&models.HsUserWithdrawal{}).
			Where("user_id = ? AND currency_code = ?  AND DATE(requested_at) = ?",
				withdrawal.UserId, withdrawal.CurrencyCode, today).
			Count(&todayCount).Error; err != nil {
			return false, "", fmt.Errorf("查询今日提现次数失败: %w", err)
		}

		// 加上当前这笔
		if int(todayCount)+1 > dailyCount {
			return false, fmt.Sprintf("超出每日提现次数限制：已提现 %d 次，限额 %d 次", todayCount, dailyCount), nil
		}
	}

	return true, "", nil
}

// updateStatusToReview 更新状态为人工审核
func (p *WithdrawalProcessor) updateStatusToReview(withdrawal *models.HsUserWithdrawal, reason string) error {
	reasonJSON, _ := json.Marshal(map[string]string{
		"auto_review_reason": reason,
		"time":               time.Now().Format("2006-01-02 15:04:05"),
	})

	return p.db.Model(&models.HsUserWithdrawal{}).
		Where("id = ? AND status = ?", withdrawal.Id, "pending").
		Updates(map[string]interface{}{
			"status": "review",
			"reason": string(reasonJSON),
		}).Error
}

// submitToPandaPay 提交到 PandaPay 进行转账（使用与approve相同的逻辑）
func (p *WithdrawalProcessor) submitToPandaPay(withdrawal *models.HsUserWithdrawal) error {
	// 1. 创建 service 实例
	svc := service.HsUserWithdrawal{}
	svc.Orm = p.db
	svc.Log = log.NewHelper(log.DefaultLogger)

	// 2. 调用公共转账方法
	result, err := svc.SubmitWithdrawalPayout(withdrawal)
	if err != nil {
		// 转账失败，转人工审核
		return p.handlePayoutError(withdrawal, err.Error())
	}

	// 3. 根据结果更新提现状态
	updateData := map[string]interface{}{
		"channel_txn_id": result.ChannelTxnId,
		"status":         result.Status,
		"processed_at":   time.Now(),
	}

	// 如果有失败原因，也更新
	if result.Reason != "" {
		updateData["reason"] = result.Reason
	}

	// 4. 更新数据库
	if err := p.db.Model(&models.HsUserWithdrawal{}).
		Where("id = ?", withdrawal.Id).
		Updates(updateData).Error; err != nil {
		log.Errorf("[WithdrawalProcessor] 更新提现记录失败: %v", err)
		return err
	}

	log.Infof("[WithdrawalProcessor] 提现单 %s 已提交到 PandaPay，通道订单号: %s，状态: %s",
		withdrawal.WithdrawNo, result.ChannelTxnId, result.Status)

	return nil
}

// handlePayoutError 处理代付失败，转人工审核
func (p *WithdrawalProcessor) handlePayoutError(withdrawal *models.HsUserWithdrawal, errMsg string) error {
	log.Errorf("[WithdrawalProcessor] 提现单 %s 自动处理失败: %s", withdrawal.WithdrawNo, errMsg)

	reasonJSON, _ := json.Marshal(map[string]string{
		"auto_payout_error": errMsg,
		"time":              time.Now().Format("2006-01-02 15:04:05"),
	})

	return p.db.Model(&models.HsUserWithdrawal{}).
		Where("id = ?", withdrawal.Id).
		Updates(map[string]interface{}{
			"status": "review",
			"reason": string(reasonJSON),
		}).Error
}
