package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type OrdGiftcardWriteoffs struct {
	service.Service
}

// ============ 核销处理上下文结构体 ============

// WriteoffContext 核销处理上下文，用于在各函数间传递数据
type WriteoffContext struct {
	Order             *models.OrdUserOrders       // 订单信息
	User              *models.HsUsers             // 用户信息
	UserCurrencyCode  string                      // 用户货币代码
	IsCrypto          bool                        // 是否虚拟币
	ConfigRate        decimal.Decimal             // 配置汇率
	TargetCurrency    string                      // 目标货币
	TotalAmount       decimal.Decimal             // 核销总金额
	FinalStatus       int                         // 核销最终状态
	CreateBy          int                         // 创建人ID
	UserLevels        []models.HsConfgiUserLevels // 用户等级配置列表（预查询）
	UsdExchangeRate   decimal.Decimal             // 目标货币对美元汇率（预查询，用于经验值计算）
	FrozenLimitAmount decimal.Decimal             // 冻结限额配置（预查询，用于返利冻结计算）
}

// WriteoffItemDecimal 转换后的核销项（使用decimal类型）
type WriteoffItemDecimal struct {
	GiftCardId               int
	AdminRecognizedCode      string
	RecognizedCardValue      decimal.Decimal // 识别的卡片面值
	UserLocalCurrencyAmount  decimal.Decimal // 用户入账金额
	Status                   int
	Remark                   string
	FailureImageUrl          string
	SupplierId               int
	PlatformSettlementAmount decimal.Decimal // 平台入账金额
}

// WriteoffItemResult 单个核销项处理结果
type WriteoffItemResult struct {
	Record          models.OrdGiftcardWriteoffs
	UserLocalAmount decimal.Decimal
}

// ============ 公共方法 ============

// GetPage 获取OrdGiftcardWriteoffs列表
func (e *OrdGiftcardWriteoffs) GetPage(c *dto.OrdGiftcardWriteoffsGetPageReq, p *actions.DataPermission, list *[]models.OrdGiftcardWriteoffs, count *int64) error {
	var err error
	var data models.OrdGiftcardWriteoffs

	err = e.Orm.Model(&data).
		Select("ord_giftcard_writeoffs.*, "+
			"ord_user_orders.order_no as order_no, "+
			"hs_users.username as user_name, "+
			"ord_giftcard_category.name as gift_card_name").
		Joins("LEFT JOIN ord_user_orders ON ord_giftcard_writeoffs.order_id = ord_user_orders.id").
		Joins("LEFT JOIN hs_users ON ord_giftcard_writeoffs.user_id = hs_users.id").
		Joins("LEFT JOIN ord_giftcard ON ord_giftcard_writeoffs.gift_card_id = ord_giftcard.id").
		Joins("LEFT JOIN ord_giftcard_region ON ord_giftcard.region_id = ord_giftcard_region.id").
		Joins("LEFT JOIN ord_giftcard_category ON ord_giftcard_region.category_id = ord_giftcard_category.id").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardWriteoffsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdGiftcardWriteoffs对象
func (e *OrdGiftcardWriteoffs) Get(d *dto.OrdGiftcardWriteoffsGetReq, p *actions.DataPermission, model *models.OrdGiftcardWriteoffs) error {
	var data models.OrdGiftcardWriteoffs

	err := e.Orm.Model(&data).
		Select("ord_giftcard_writeoffs.*, "+
			"ord_user_orders.order_no as order_no, "+
			"hs_users.username as user_name, "+
			"ord_giftcard_category.name as gift_card_name").
		Joins("LEFT JOIN ord_user_orders ON ord_giftcard_writeoffs.order_id = ord_user_orders.id").
		Joins("LEFT JOIN hs_users ON ord_giftcard_writeoffs.user_id = hs_users.id").
		Joins("LEFT JOIN ord_giftcard ON ord_giftcard_writeoffs.gift_card_id = ord_giftcard.id").
		Joins("LEFT JOIN ord_giftcard_region ON ord_giftcard.region_id = ord_giftcard_region.id").
		Joins("LEFT JOIN ord_giftcard_category ON ord_giftcard_region.category_id = ord_giftcard_category.id").
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdGiftcardWriteoffs error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrdGiftcardWriteoffs对象
func (e *OrdGiftcardWriteoffs) Insert(c *dto.OrdGiftcardWriteoffsInsertReq) error {
	var err error
	var data models.OrdGiftcardWriteoffs
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardWriteoffsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改OrdGiftcardWriteoffs对象
func (e *OrdGiftcardWriteoffs) Update(c *dto.OrdGiftcardWriteoffsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.OrdGiftcardWriteoffs{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("OrdGiftcardWriteoffsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除OrdGiftcardWriteoffs
func (e *OrdGiftcardWriteoffs) Remove(d *dto.OrdGiftcardWriteoffsDeleteReq, p *actions.DataPermission) error {
	var data models.OrdGiftcardWriteoffs

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Unscoped().Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveOrdGiftcardWriteoffs error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// ============ 批量核销主流程（重构后） ============

// BatchInsert 批量创建核销记录（带事务处理：余额增加、流水记录、账户分成）
func (e *OrdGiftcardWriteoffs) BatchInsert(c *dto.OrdGiftcardWriteoffsBatchInsertReq) error {
	// 1. 验证核销列表
	if len(c.WriteoffList) == 0 {
		return errors.New("核销记录列表不能为空")
	}

	// 2. 获取核销状态（以第一条记录为准）
	finalStatus := c.WriteoffList[0].Status
	if finalStatus != 1 && finalStatus != 2 {
		e.Log.Errorf("BatchInsert invalid writeoff status: %d", finalStatus)
		return errors.New("核销状态无效，只能是已核销(1)或失败(2)")
	}

	// 3. 根据状态分别处理
	if finalStatus == 2 {
		// 驳回状态：单独处理，不需要转换decimal
		return e.processRejectedWriteoff(c)
	}

	// 成功状态：转换decimal后处理
	return e.processSuccessWriteoff(c)
}

// ============ 驳回状态处理 ============

// processRejectedWriteoff 处理驳回状态的核销
func (e *OrdGiftcardWriteoffs) processRejectedWriteoff(c *dto.OrdGiftcardWriteoffsBatchInsertReq) error {
	// 1. 验证订单
	order, err := e.validateAndGetOrder(c.OrderId)
	if err != nil {
		return err
	}

	// 2. 构建驳回核销记录（不需要计算金额）
	writeoffRecords := make([]models.OrdGiftcardWriteoffs, 0, len(c.WriteoffList))
	for _, item := range c.WriteoffList {
		record := models.OrdGiftcardWriteoffs{
			UserId:              order.UserId,
			OrderId:             c.OrderId,
			GiftCardId:          item.GiftCardId,
			Status:              item.Status,
			Remark:              item.Remark,
			AdminRecognizedCode: item.AdminRecognizedCode,
			RecognizedCardValue: item.RecognizedCardValue,
			FailureImageUrl:     item.FailureImageUrl,
			SupplierId:          strconv.Itoa(item.SupplierId),
		}
		record.CreateBy = c.CreateBy
		writeoffRecords = append(writeoffRecords, record)
	}

	// 3. 准备订单状态更新（驳回）
	orderUpdates := map[string]interface{}{
		"status":                 4, // 订单驳回
		"processing_status":      2, // 处理取消
		"processing_started_end": time.Now(),
	}

	// 4. 执行事务
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		// 加锁查询订单
		var orderLocked models.OrdUserOrders
		err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("id = ?", c.OrderId).
			First(&orderLocked).Error
		if err != nil {
			e.Log.Errorf("processRejectedWriteoff lock order error:%s", err)
			return errors.New("获取订单锁失败")
		}

		if orderLocked.Status != 1 {
			e.Log.Errorf("processRejectedWriteoff order status changed: orderId=%d, status=%d", orderLocked.Id, orderLocked.Status)
			return errors.New("订单状态已变更，无法核销")
		}

		// 插入核销记录
		if err = tx.Create(&writeoffRecords).Error; err != nil {
			e.Log.Errorf("processRejectedWriteoff insert records error:%s", err)
			return err
		}

		// 更新订单状态
		if err = tx.Model(&models.OrdUserOrders{}).Where("id = ?", c.OrderId).Updates(orderUpdates).Error; err != nil {
			e.Log.Errorf("processRejectedWriteoff update order error:%s", err)
			return err
		}

		return nil
	})
}

// ============ 成功状态处理 ============

// SingleWriteoffResult 单个核销项处理结果
type SingleWriteoffResult struct {
	Index           int
	Success         bool
	Error           error
	UserLocalAmount decimal.Decimal
	WriteoffId      int
	ItemStatus      int // 核销项状态：1=成功，2=失败
}

// WriteoffItemBaseInfo 单个核销项的基础信息（事务外查询）
type WriteoffItemBaseInfo struct {
	Giftcard        *models.OrdGiftcard        // 礼品卡配置
	Region          *models.OrdGiftcardRegion  // 区域配置
	UserLevel       *models.HsConfgiUserLevels // 用户等级配置（可能为nil）
	RebateRate      decimal.Decimal            // 用户返利比例
	WriteoffRecord  models.OrdGiftcardWriteoffs // 预构建的核销记录
	UserLocalAmount decimal.Decimal            // 用户入账金额
}

// queryWriteoffItemBaseInfo 查询单个核销项的基础信息（事务外）
func (e *OrdGiftcardWriteoffs) queryWriteoffItemBaseInfo(item WriteoffItemDecimal, ctx *WriteoffContext, orderId int, index int) (*WriteoffItemBaseInfo, error) {
	baseInfo := &WriteoffItemBaseInfo{}

	// 1. 查询礼品卡配置
	var giftcard models.OrdGiftcard
	err := e.Orm.Where("id = ?", item.GiftCardId).First(&giftcard).Error
	if err != nil {
		e.Log.Errorf("queryWriteoffItemBaseInfo get giftcard error: giftcardId=%d, err=%s", item.GiftCardId, err)
		return nil, fmt.Errorf("第%d条核销记录的礼品卡配置不存在：ID=%d", index+1, item.GiftCardId)
	}
	baseInfo.Giftcard = &giftcard

	// 2. 查询礼品卡区域配置
	if giftcard.RegionId == "" || giftcard.RegionId == "0" {
		e.Log.Errorf("queryWriteoffItemBaseInfo giftcard region not configured: giftcardId=%d", item.GiftCardId)
		return nil, fmt.Errorf("第%d条核销记录的礼品卡未配置区域：ID=%d", index+1, item.GiftCardId)
	}

	var region models.OrdGiftcardRegion
	err = e.Orm.Where("id = ?", giftcard.RegionId).First(&region).Error
	if err != nil {
		e.Log.Errorf("queryWriteoffItemBaseInfo get region error: regionId=%s, err=%s", giftcard.RegionId, err)
		return nil, fmt.Errorf("第%d条核销记录的礼品卡区域配置不存在：RegionID=%s", index+1, giftcard.RegionId)
	}
	baseInfo.Region = &region

	// 3. 查询用户等级配置和返利比例
	baseInfo.RebateRate = decimal.Zero
	if ctx.User.LevelId != "" && ctx.User.LevelId != "0" {
		var userLevel models.HsConfgiUserLevels
		err = e.Orm.Where("id = ? AND is_active = 1", ctx.User.LevelId).First(&userLevel).Error
		if err != nil {
			e.Log.Errorf("queryWriteoffItemBaseInfo get user level error: levelId=%s, err=%s", ctx.User.LevelId, err)
			return nil, fmt.Errorf("用户等级配置不存在或未启用：LevelID=%s", ctx.User.LevelId)
		}
		baseInfo.UserLevel = &userLevel

		if userLevel.RebateRate != "" {
			rebateRate, err := decimal.NewFromString(userLevel.RebateRate)
			if err != nil {
				e.Log.Errorf("queryWriteoffItemBaseInfo parse rebate rate error: rate=%s, err=%s", userLevel.RebateRate, err)
				return nil, fmt.Errorf("用户等级返利比例格式错误：%s", userLevel.RebateRate)
			}
			if rebateRate.GreaterThan(decimal.NewFromInt(1)) {
				rebateRate = rebateRate.Div(decimal.NewFromInt(100))
			}
			baseInfo.RebateRate = rebateRate
		}
	}

	// 4. 构建核销记录
	writeoffResult, err := e.buildWriteoffRecordDecimal(item, ctx, orderId, &giftcard, &region, index)
	if err != nil {
		return nil, err
	}
	baseInfo.WriteoffRecord = writeoffResult.Record
	baseInfo.UserLocalAmount = writeoffResult.UserLocalAmount

	return baseInfo, nil
}

// processSuccessWriteoff 处理成功状态的核销（每个核销项单独事务）
func (e *OrdGiftcardWriteoffs) processSuccessWriteoff(c *dto.OrdGiftcardWriteoffsBatchInsertReq) error {
	// 1. 转换核销项为decimal类型
	decimalItems, err := e.convertToDecimalItems(c.WriteoffList)
	if err != nil {
		return err
	}

	// 2. 初始化核销上下文（不在事务内）
	ctx, err := e.initWriteoffContext(c)
	if err != nil {
		return err
	}

	// 3. 逐个处理核销项，每个项单独事务
	results := make([]SingleWriteoffResult, 0, len(decimalItems))

	for i, item := range decimalItems {
		// 3.1 查询基础信息（事务外）
		baseInfo, err := e.queryWriteoffItemBaseInfo(item, ctx, c.OrderId, i)
		if err != nil {
			results = append(results, SingleWriteoffResult{
				Index:   i,
				Success: false,
				Error:   err,
			})
			e.Log.Errorf("processSuccessWriteoff item %d query base info failed: %s", i, err)
			continue
		}

		// 3.2 执行事务（只做插入和更新）
		result := e.executeSingleWriteoffTransaction(item, ctx, baseInfo, c.OrderId, i)
		results = append(results, result)

		if !result.Success {
			e.Log.Errorf("processSuccessWriteoff item %d transaction failed: %s", i, result.Error)
		}
	}

	// 4. 检查是否有成功入账的核销项（只有Status=1且事务成功才算成功）
	successCount := 0
	for _, r := range results {
		if r.Success && r.ItemStatus == 1 {
			successCount++
		}
	}

	if successCount == 0 {
		// 所有成功状态的核销项都失败了，返回第一个错误
		for _, r := range results {
			if !r.Success && r.Error != nil {
				return r.Error
			}
		}
		return errors.New("所有核销项处理失败")
	}

	// 5. 更新订单状态（只要有一个成功，订单就是成功）
	err = e.updateOrderStatusAfterWriteoff(c.OrderId)
	if err != nil {
		e.Log.Errorf("processSuccessWriteoff update order status error: %s", err)
		return err
	}

	return nil
}

// executeSingleWriteoffTransaction 执行单个核销项事务（只做插入和更新）
func (e *OrdGiftcardWriteoffs) executeSingleWriteoffTransaction(item WriteoffItemDecimal, ctx *WriteoffContext, baseInfo *WriteoffItemBaseInfo, orderId int, index int) SingleWriteoffResult {
	result := SingleWriteoffResult{
		Index:           index,
		Success:         false,
		UserLocalAmount: baseInfo.UserLocalAmount,
		ItemStatus:      item.Status, // 记录核销项状态
	}

	err := e.Orm.Transaction(func(tx *gorm.DB) error {
		// 1. 插入核销记录
		record := baseInfo.WriteoffRecord
		if err := tx.Create(&record).Error; err != nil {
			e.Log.Errorf("executeSingleWriteoffTransaction insert record error: %s", err)
			return fmt.Errorf("第%d条核销记录插入失败", index+1)
		}
		result.WriteoffId = record.Id

		// 2. 如果是成功状态且金额大于0，处理余额、经验值、邀请分成
		if item.Status == 1 && baseInfo.UserLocalAmount.GreaterThan(decimal.Zero) {
			// 2.1 获取用户信息（加锁）- 这是唯一需要在事务内的查询
			var user models.HsUsers
			err := tx.Set("gorm:query_option", "FOR UPDATE").
				Where("id = ?", ctx.Order.UserId).
				First(&user).Error
			if err != nil {
				e.Log.Errorf("executeSingleWriteoffTransaction lock user error: %s", err)
				return errors.New("获取用户信息失败")
			}

			// 2.2 用户入账（使用预查询的返利比例和冻结限额）
			err = e.creditUserBalanceWithBaseInfo(tx, &user, baseInfo.UserLocalAmount, baseInfo.RebateRate, ctx.TargetCurrency, ctx.IsCrypto, ctx.FrozenLimitAmount, orderId, ctx.CreateBy)
			if err != nil {
				e.Log.Errorf("executeSingleWriteoffTransaction credit balance error: %s", err)
				return err
			}

			// 2.3 增加经验值
			err = e.addUserExperienceWithContext(tx, ctx, baseInfo.UserLocalAmount, orderId)
			if err != nil {
				e.Log.Errorf("executeSingleWriteoffTransaction add experience error: %s", err)
				return err
			}

			// 2.4 处理邀请分成
			amountFloat, _ := baseInfo.UserLocalAmount.Float64()
			frozenLimitFloat, _ := ctx.FrozenLimitAmount.Float64()
			err = e.processInviteCommissions(tx, ctx.Order.UserId, orderId, amountFloat, ctx.TargetCurrency, ctx.IsCrypto, frozenLimitFloat, ctx.CreateBy)
			if err != nil {
				e.Log.Errorf("executeSingleWriteoffTransaction process commissions error: %s", err)
				return err
			}
		}

		return nil
	})

	if err != nil {
		result.Error = err
		result.Success = false
	} else {
		result.Success = true
	}

	return result
}

// creditUserBalanceWithBaseInfo 用户入账（使用预查询的返利比例和冻结限额）
func (e *OrdGiftcardWriteoffs) creditUserBalanceWithBaseInfo(tx *gorm.DB, user *models.HsUsers, amount decimal.Decimal, rebateRate decimal.Decimal, currencyCode string, isCrypto bool, frozenLimit decimal.Decimal, orderId int, createBy int) error {
	// 计算返利金额和可用金额
	rebateAmount := amount.Mul(rebateRate)
	availableAmount := amount.Sub(rebateAmount)

	// 根据法币/虚拟币分别处理
	if isCrypto {
		return e.creditCryptoBalanceWithRebateDecimalTx(tx, user, availableAmount, rebateAmount, currencyCode, frozenLimit, orderId, createBy)
	}
	return e.creditFiatBalanceWithRebateDecimalTx(tx, user, availableAmount, rebateAmount, currencyCode, frozenLimit, orderId, createBy)
}

// addUserExperienceWithContext 增加用户经验值（使用预查询的上下文数据）
// amount: 用户入账金额，单位为用户本地货币（如NGN、KES等）
// 经验值计算: amount(本地货币) × UsdExchangeRate(本地货币→USD) = 美元金额 = 经验值
func (e *OrdGiftcardWriteoffs) addUserExperienceWithContext(tx *gorm.DB, ctx *WriteoffContext, amount decimal.Decimal, orderId int) error {
	// 将用户本地货币金额转换为美元金额作为经验值
	experienceAmount := amount.Mul(ctx.UsdExchangeRate)

	// 转换为 int 类型
	experienceInt := int(experienceAmount.IntPart())
	return e.addUserExperience(tx, ctx.Order.UserId, experienceInt, ctx.UserLevels, "giftcard_writeoff", orderId, ctx.CreateBy)
}

// creditUserBalanceWithRebateTx 用户入账（在事务内，使用传入的user对象）
func (e *OrdGiftcardWriteoffs) creditUserBalanceWithRebateTx(tx *gorm.DB, user *models.HsUsers, amount decimal.Decimal, currencyCode string, isCrypto bool, frozenLimit decimal.Decimal, orderId int, createBy int) error {
	// 获取用户等级返利比例
	rebateRateDecimal, err := e.getUserRebateRateDecimal(tx, user)
	if err != nil {
		return err
	}

	// 计算返利金额和可用金额
	rebateAmount := amount.Mul(rebateRateDecimal)
	availableAmount := amount.Sub(rebateAmount)

	// 根据法币/虚拟币分别处理
	if isCrypto {
		return e.creditCryptoBalanceWithRebateDecimalTx(tx, user, availableAmount, rebateAmount, currencyCode, frozenLimit, orderId, createBy)
	}
	return e.creditFiatBalanceWithRebateDecimalTx(tx, user, availableAmount, rebateAmount, currencyCode, frozenLimit, orderId, createBy)
}

// creditFiatBalanceWithRebateDecimalTx 法币入账（在事务内，使用传入的user对象）
func (e *OrdGiftcardWriteoffs) creditFiatBalanceWithRebateDecimalTx(tx *gorm.DB, user *models.HsUsers, availableAmount, rebateAmount decimal.Decimal, currencyCode string, frozenLimit decimal.Decimal, orderId int, createBy int) error {
	// 1. 获取当前余额
	balanceBefore, _ := decimal.NewFromString(user.Balance)
	frozenBalanceBefore, _ := decimal.NewFromString(user.FrozenBalance)

	// 2. 处理可用金额入账
	balanceAfter := balanceBefore.Add(availableAmount)

	// 3. 处理返利金额到冻结余额（根据冻结配置）
	// rebateAmount 全部先进入冻结，达到阈值后 overflowAmount 转出到可用
	frozenDelta, overflowAmount := e.calculateFrozenAllocationDecimal(frozenBalanceBefore, rebateAmount, frozenLimit)
	frozenBalanceAfter := frozenBalanceBefore.Add(frozenDelta)
	balanceAfter = balanceAfter.Add(overflowAmount)

	// 4. 更新用户余额（使用乐观锁）
	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", user.Id, user.Version).
		Updates(map[string]interface{}{
			"balance":        balanceAfter.StringFixed(2),
			"frozen_balance": frozenBalanceAfter.StringFixed(2),
			"version":        gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		e.Log.Errorf("creditFiatBalanceWithRebateDecimalTx update balance error:%s", result.Error)
		return errors.New("更新用户法币余额失败")
	}

	if result.RowsAffected == 0 {
		e.Log.Errorf("creditFiatBalanceWithRebateDecimalTx update balance conflict for user %d", user.Id)
		return errors.New("法币余额更新冲突，请重试")
	}

	// 5. 创建流水记录（分开记录）
	nanoTime := time.Now().UnixNano()
	balanceRunning := balanceBefore
	frozenBalanceRunning := frozenBalanceBefore

	// 5.1 可用余额流水1：核销金额入账
	if availableAmount.GreaterThan(decimal.Zero) {
		newBalance := balanceRunning.Add(availableAmount)
		err := e.createFiatLedgerDecimal(tx, user.Id, availableAmount, balanceRunning, newBalance,
			currencyCode, "giftcard_writeoff_fiat", orderId, nanoTime, "礼品卡核销到账", createBy)
		if err != nil {
			return err
		}
		balanceRunning = newBalance
	}

	// 5.2 冻结余额流水1：返利金额全部进入冻结
	if rebateAmount.GreaterThan(decimal.Zero) {
		newFrozenBalance := frozenBalanceRunning.Add(rebateAmount)
		err := e.createFiatFrozenLedgerDecimal(tx, user.Id, rebateAmount, frozenBalanceRunning, newFrozenBalance,
			currencyCode, "giftcard_writeoff_rebate_frozen_in", orderId, nanoTime+1, "礼品卡核销返利（转入冻结）", createBy)
		if err != nil {
			return err
		}
		frozenBalanceRunning = newFrozenBalance
	}

	// 5.3 如果有达到阈值转出的金额
	if overflowAmount.GreaterThan(decimal.Zero) {
		// 冻结余额流水2：达到阈值后从冻结转出
		newFrozenBalance := frozenBalanceRunning.Sub(overflowAmount)
		err := e.createFiatFrozenLedgerOutDecimal(tx, user.Id, overflowAmount, frozenBalanceRunning, newFrozenBalance,
			currencyCode, "giftcard_writeoff_rebate_frozen_out", orderId, nanoTime+2, "礼品卡返利达到阈值（转出冻结）", createBy)
		if err != nil {
			return err
		}

		// 可用余额流水2：从冻结余额转入
		newBalance := balanceRunning.Add(overflowAmount)
		err = e.createFiatLedgerDecimal(tx, user.Id, overflowAmount, balanceRunning, newBalance,
			currencyCode, "giftcard_writeoff_rebate_unfrozen", orderId, nanoTime+3, "礼品卡返利（冻结转入可用）", createBy)
		if err != nil {
			return err
		}
	}

	e.Log.Infof("User %d fiat credited: available=%s, rebate=%s, overflow=%s, currency=%s",
		user.Id, availableAmount.String(), rebateAmount.String(), overflowAmount.String(), currencyCode)

	return nil
}

// creditCryptoBalanceWithRebateDecimalTx 虚拟币入账（在事务内，使用传入的user对象）
func (e *OrdGiftcardWriteoffs) creditCryptoBalanceWithRebateDecimalTx(tx *gorm.DB, user *models.HsUsers, availableAmount, rebateAmount decimal.Decimal, currencyCode string, frozenLimit decimal.Decimal, orderId int, createBy int) error {
	// 1. 获取当前余额
	balanceBefore, _ := decimal.NewFromString(user.CryptoBalance)
	frozenBalanceBefore, _ := decimal.NewFromString(user.CryptoFrozenBalance)

	// 2. 处理可用金额入账
	balanceAfter := balanceBefore.Add(availableAmount)

	// 3. 处理返利金额到冻结余额（根据冻结配置）
	// rebateAmount 全部先进入冻结，达到阈值后 overflowAmount 转出到可用
	frozenDelta, overflowAmount := e.calculateFrozenAllocationDecimal(frozenBalanceBefore, rebateAmount, frozenLimit)
	frozenBalanceAfter := frozenBalanceBefore.Add(frozenDelta)
	balanceAfter = balanceAfter.Add(overflowAmount)

	// 4. 更新用户余额（使用乐观锁）
	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", user.Id, user.Version).
		Updates(map[string]interface{}{
			"crypto_balance":        balanceAfter.StringFixed(8),
			"crypto_frozen_balance": frozenBalanceAfter.StringFixed(8),
			"version":               gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		e.Log.Errorf("creditCryptoBalanceWithRebateDecimalTx update balance error:%s", result.Error)
		return errors.New("更新用户虚拟币余额失败")
	}

	if result.RowsAffected == 0 {
		e.Log.Errorf("creditCryptoBalanceWithRebateDecimalTx update balance conflict for user %d", user.Id)
		return errors.New("虚拟币余额更新冲突，请重试")
	}

	// 5. 创建流水记录（分开记录）
	nanoTime := time.Now().UnixNano()
	balanceRunning := balanceBefore
	frozenBalanceRunning := frozenBalanceBefore

	// 5.1 可用余额流水1：核销金额入账
	if availableAmount.GreaterThan(decimal.Zero) {
		newBalance := balanceRunning.Add(availableAmount)
		err := e.createCryptoLedgerDecimal(tx, user.Id, availableAmount, balanceRunning, newBalance,
			currencyCode, "giftcard_writeoff_crypto", orderId, nanoTime, "礼品卡核销到账", createBy)
		if err != nil {
			return err
		}
		balanceRunning = newBalance
	}

	// 5.2 冻结余额流水1：返利金额全部进入冻结
	if rebateAmount.GreaterThan(decimal.Zero) {
		newFrozenBalance := frozenBalanceRunning.Add(rebateAmount)
		err := e.createCryptoFrozenLedgerDecimal(tx, user.Id, rebateAmount, frozenBalanceRunning, newFrozenBalance,
			currencyCode, "giftcard_writeoff_rebate_frozen_in", orderId, nanoTime+1, "礼品卡核销返利（转入冻结）", createBy)
		if err != nil {
			return err
		}
		frozenBalanceRunning = newFrozenBalance
	}

	// 5.3 如果有达到阈值转出的金额
	if overflowAmount.GreaterThan(decimal.Zero) {
		// 冻结余额流水2：达到阈值后从冻结转出
		newFrozenBalance := frozenBalanceRunning.Sub(overflowAmount)
		err := e.createCryptoFrozenLedgerOutDecimal(tx, user.Id, overflowAmount, frozenBalanceRunning, newFrozenBalance,
			currencyCode, "giftcard_writeoff_rebate_frozen_out", orderId, nanoTime+2, "礼品卡返利达到阈值（转出冻结）", createBy)
		if err != nil {
			return err
		}

		// 可用余额流水2：从冻结余额转入
		newBalance := balanceRunning.Add(overflowAmount)
		err = e.createCryptoLedgerDecimal(tx, user.Id, overflowAmount, balanceRunning, newBalance,
			currencyCode, "giftcard_writeoff_rebate_unfrozen", orderId, nanoTime+3, "礼品卡返利（冻结转入可用）", createBy)
		if err != nil {
			return err
		}
	}

	e.Log.Infof("User %d crypto credited: available=%s, rebate=%s, overflow=%s, currency=%s",
		user.Id, availableAmount.String(), rebateAmount.String(), overflowAmount.String(), currencyCode)

	return nil
}

// updateOrderStatusAfterWriteoff 核销完成后更新订单状态（只要有一个成功，订单就是成功）
func (e *OrdGiftcardWriteoffs) updateOrderStatusAfterWriteoff(orderId int) error {
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		// 加锁查询订单
		var order models.OrdUserOrders
		err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("id = ?", orderId).
			First(&order).Error
		if err != nil {
			e.Log.Errorf("updateOrderStatusAfterWriteoff lock order error:%s", err)
			return errors.New("获取订单锁失败")
		}

		// 准备更新数据：订单完成
		orderUpdates := map[string]interface{}{
			"status":                 2, // 订单完成
			"processing_status":      3, // 处理完成
			"processing_started_end": time.Now(),
			"completed_at":           time.Now(),
		}

		// 更新订单状态
		if err = tx.Model(&models.OrdUserOrders{}).Where("id = ?", orderId).Updates(orderUpdates).Error; err != nil {
			e.Log.Errorf("updateOrderStatusAfterWriteoff update order error:%s", err)
			return err
		}

		return nil
	})
}

// convertToDecimalItems 将请求中的核销项转换为decimal类型
func (e *OrdGiftcardWriteoffs) convertToDecimalItems(items []dto.OrdGiftcardWriteoffsBatchItem) ([]WriteoffItemDecimal, error) {
	result := make([]WriteoffItemDecimal, 0, len(items))

	for i, item := range items {
		decimalItem := WriteoffItemDecimal{
			GiftCardId:          item.GiftCardId,
			AdminRecognizedCode: item.AdminRecognizedCode,
			Status:              item.Status,
			Remark:              item.Remark,
			FailureImageUrl:     item.FailureImageUrl,
			SupplierId:          item.SupplierId,
		}

		// 转换 RecognizedCardValue
		if item.RecognizedCardValue != "" && item.RecognizedCardValue != "0" {
			val, err := decimal.NewFromString(item.RecognizedCardValue)
			if err != nil {
				e.Log.Errorf("convertToDecimalItems parse RecognizedCardValue error: index=%d, value=%s, err=%s", i, item.RecognizedCardValue, err)
				return nil, fmt.Errorf("第%d条核销记录的卡片面值格式错误：%s", i+1, item.RecognizedCardValue)
			}
			decimalItem.RecognizedCardValue = val
		}

		// 转换 UserLocalCurrencyAmount
		if item.UserLocalCurrencyAmount != "" && item.UserLocalCurrencyAmount != "0" {
			val, err := decimal.NewFromString(item.UserLocalCurrencyAmount)
			if err != nil {
				e.Log.Errorf("convertToDecimalItems parse UserLocalCurrencyAmount error: index=%d, value=%s, err=%s", i, item.UserLocalCurrencyAmount, err)
				return nil, fmt.Errorf("第%d条核销记录的用户入账金额格式错误：%s", i+1, item.UserLocalCurrencyAmount)
			}
			decimalItem.UserLocalCurrencyAmount = val
		}

		// 转换 PlatformSettlementAmount
		if item.PlatformSettlementAmount != "" && item.PlatformSettlementAmount != "0" {
			val, err := decimal.NewFromString(item.PlatformSettlementAmount)
			if err != nil {
				e.Log.Errorf("convertToDecimalItems parse PlatformSettlementAmount error: index=%d, value=%s, err=%s", i, item.PlatformSettlementAmount, err)
				return nil, fmt.Errorf("第%d条核销记录的平台入账金额格式错误：%s", i+1, item.PlatformSettlementAmount)
			}
			decimalItem.PlatformSettlementAmount = val
		}

		result = append(result, decimalItem)
	}

	return result, nil
}

// initWriteoffContext 初始化核销上下文
func (e *OrdGiftcardWriteoffs) initWriteoffContext(c *dto.OrdGiftcardWriteoffsBatchInsertReq) (*WriteoffContext, error) {
	ctx := &WriteoffContext{
		CreateBy:    c.CreateBy,
		FinalStatus: 1, // 成功状态
		TotalAmount: decimal.Zero,
	}

	// 1. 查询并验证订单
	order, err := e.validateAndGetOrder(c.OrderId)
	if err != nil {
		return nil, err
	}
	ctx.Order = order
	ctx.IsCrypto = order.BalanceType == 2

	// 2. 获取用户货币信息
	user, userCurrencyCode, err := e.getUserCurrencyInfo(order.UserId)
	if err != nil {
		return nil, err
	}
	ctx.User = user
	ctx.UserCurrencyCode = userCurrencyCode

	// 3. 计算配置汇率
	configRate, targetCurrency, err := e.calculateConfigRateDecimal(order, userCurrencyCode)
	if err != nil {
		return nil, err
	}
	ctx.ConfigRate = configRate
	ctx.TargetCurrency = targetCurrency

	// 4. 预查询用户等级配置列表（用于升级检查）
	userLevels, err := e.queryUserLevels()
	if err != nil {
		e.Log.Errorf("initWriteoffContext query user levels error: %s", err)
		// 等级配置查询失败不阻塞核销流程，但会导致无法升级
		ctx.UserLevels = nil
	} else {
		ctx.UserLevels = userLevels
	}

	// 5. 预查询用户本地货币对美元汇率（用于经验值计算）
	// 场景示例：礼品卡CAD加元 → 用户NGN尼日利亚元 → 经验值USD美元
	// UserLocalAmount 已是用户本地货币(NGN)金额，只需 NGN→USD 汇率转换
	// 经验值 = UserLocalAmount(本地货币) × UsdExchangeRate = USD金额(取整)
	if !ctx.IsCrypto && targetCurrency != "USD" {
		usdRate, err := e.getUsdExchangeRate(targetCurrency)
		if err != nil {
			e.Log.Errorf("initWriteoffContext get USD rate error: targetCurrency=%s, err=%s", targetCurrency, err)
			return nil, fmt.Errorf("获取货币 %s 对美元汇率失败，请先配置汇率", targetCurrency)
		}
		ctx.UsdExchangeRate = usdRate
	} else {
		// 虚拟币(USDT视为USD)或用户本地货币已是USD，汇率为1
		ctx.UsdExchangeRate = decimal.NewFromInt(1)
	}

	// 6. 预查询冻结限额配置（用于返利冻结计算）
	frozenLimit, err := e.queryFrozenLimitAmount(targetCurrency, ctx.IsCrypto)
	if err != nil {
		e.Log.Errorf("initWriteoffContext query frozen limit error: %s", err)
		// 冻结限额查询失败，默认无限额（全部进入冻结）
		ctx.FrozenLimitAmount = decimal.Zero
	} else {
		ctx.FrozenLimitAmount = frozenLimit
	}

	return ctx, nil
}

// queryUserLevels 查询所有启用的用户等级配置
func (e *OrdGiftcardWriteoffs) queryUserLevels() ([]models.HsConfgiUserLevels, error) {
	var levels []models.HsConfgiUserLevels
	err := e.Orm.Where("is_active = ?", "1").Order("sort_order ASC").Find(&levels).Error
	if err != nil {
		return nil, fmt.Errorf("查询等级配置失败: %w", err)
	}
	return levels, nil
}

// getUsdExchangeRate 获取目标货币对美元的汇率
func (e *OrdGiftcardWriteoffs) getUsdExchangeRate(targetCurrency string) (decimal.Decimal, error) {
	rate, err := e.getCurrencyRateFromDB(targetCurrency, "USD")
	if err != nil {
		return decimal.Zero, err
	}
	rateDecimal, err := decimal.NewFromString(rate)
	if err != nil {
		return decimal.Zero, fmt.Errorf("汇率格式错误：%s", rate)
	}
	return rateDecimal, nil
}

// getCurrencyRateFromDB 从数据库获取货币汇率（不使用事务）
func (e *OrdGiftcardWriteoffs) getCurrencyRateFromDB(fromCurrency, toCurrency string) (string, error) {
	if fromCurrency == toCurrency {
		return "1.00000000", nil
	}

	var rateRecord models.OrdConfigCurrencyRates
	err := e.Orm.Where("base_currency_code = ? AND quote_currency_code = ? AND status = 1",
		fromCurrency, toCurrency).
		First(&rateRecord).Error
	if err != nil {
		return "", fmt.Errorf("获取货币 %s 到 %s 的汇率失败", fromCurrency, toCurrency)
	}
	return rateRecord.Rate, nil
}

// queryFrozenLimitAmount 查询冻结限额配置（不使用事务）
func (e *OrdGiftcardWriteoffs) queryFrozenLimitAmount(currencyCode string, isCrypto bool) (decimal.Decimal, error) {
	currencyType := "fiat"
	if isCrypto {
		currencyType = "crypto"
	}

	var frozenLimitConfig models.HsConfigFrozenLimit
	err := e.Orm.Where("currency_type = ? AND currency_code = ? AND is_active = 1", currencyType, currencyCode).
		First(&frozenLimitConfig).Error

	if err != nil {
		// 没有配置限额，返回0表示无限额
		return decimal.Zero, nil
	}

	frozenLimit, err := decimal.NewFromString(frozenLimitConfig.FrozenLimitAmount)
	if err != nil {
		return decimal.Zero, fmt.Errorf("冻结限额格式错误：%s", frozenLimitConfig.FrozenLimitAmount)
	}

	return frozenLimit, nil
}

// validateAndGetOrder 验证并获取订单
func (e *OrdGiftcardWriteoffs) validateAndGetOrder(orderId int) (*models.OrdUserOrders, error) {
	var order models.OrdUserOrders
	err := e.Orm.Where("id = ?", orderId).First(&order).Error
	if err != nil {
		e.Log.Errorf("BatchInsert get order error:%s", err)
		return nil, errors.New("获取订单信息失败")
	}

	if order.Status != 1 {
		e.Log.Errorf("BatchInsert order status invalid: orderId=%d, status=%d", order.Id, order.Status)
		return nil, errors.New("订单状态不正确，只有已接单的订单才能核销")
	}

	return &order, nil
}

// getUserCurrencyInfo 获取用户货币信息
func (e *OrdGiftcardWriteoffs) getUserCurrencyInfo(userId int) (*models.HsUsers, string, error) {
	var user models.HsUsers
	err := e.Orm.Select("id, currency_code, level_id").Where("id = ?", userId).First(&user).Error
	if err != nil {
		e.Log.Errorf("BatchInsert get user error:%s", err)
		return nil, "", errors.New("获取用户信息失败")
	}

	if user.CurrencyCode == "" {
		e.Log.Errorf("BatchInsert user currency code is empty for user %d", userId)
		return nil, "", errors.New("用户未设置货币代码")
	}

	return &user, user.CurrencyCode, nil
}

// calculateConfigRateDecimal 计算配置汇率（返回decimal类型）
func (e *OrdGiftcardWriteoffs) calculateConfigRateDecimal(order *models.OrdUserOrders, userCurrencyCode string) (decimal.Decimal, string, error) {
	isCrypto := order.BalanceType == 2
	var configRate decimal.Decimal
	var targetCurrency string

	if isCrypto {
		targetCurrency = "USD"
		if order.CurrencyCode == "USD" {
			configRate = decimal.NewFromInt(1)
		} else {
			rateStr, err := e.getCurrencyRate(e.Orm, order.CurrencyCode, "USD")
			if err != nil {
				return decimal.Zero, "", err
			}
			configRate, err = decimal.NewFromString(rateStr)
			if err != nil {
				e.Log.Errorf("calculateConfigRateDecimal parse rate error: rate=%s, err=%s", rateStr, err)
				return decimal.Zero, "", fmt.Errorf("汇率格式错误：%s", rateStr)
			}
		}
	} else {
		targetCurrency = userCurrencyCode
		if order.CurrencyCode == userCurrencyCode {
			configRate = decimal.NewFromInt(1)
		} else {
			rateStr, err := e.getCurrencyRateViaUSD(e.Orm, order.CurrencyCode, userCurrencyCode)
			if err != nil {
				return decimal.Zero, "", err
			}
			configRate, err = decimal.NewFromString(rateStr)
			if err != nil {
				e.Log.Errorf("calculateConfigRateDecimal parse rate error: rate=%s, err=%s", rateStr, err)
				return decimal.Zero, "", fmt.Errorf("汇率格式错误：%s", rateStr)
			}
		}
	}

	return configRate, targetCurrency, nil
}

// buildWriteoffRecordDecimal 构建单条核销记录（使用decimal类型）
func (e *OrdGiftcardWriteoffs) buildWriteoffRecordDecimal(item WriteoffItemDecimal, ctx *WriteoffContext, orderId int, giftcard *models.OrdGiftcard, region *models.OrdGiftcardRegion, index int) (*WriteoffItemResult, error) {
	result := &WriteoffItemResult{}

	// 1. 获取折扣率
	var platformSaleRate string
	var discountRate decimal.Decimal
	if giftcard != nil && giftcard.DiscountRate != "" {
		platformSaleRate = giftcard.DiscountRate
		var err error
		discountRate, err = decimal.NewFromString(platformSaleRate)
		if err != nil {
			e.Log.Errorf("buildWriteoffRecordDecimal parse discount rate error: rate=%s, err=%s", platformSaleRate, err)
			return nil, fmt.Errorf("第%d条核销记录的折扣率格式错误：%s", index+1, platformSaleRate)
		}
	}

	// 2. 用户本地货币金额（必传参数）
	if item.UserLocalCurrencyAmount.IsZero() {
		return nil, fmt.Errorf("第%d条核销记录缺少用户入账金额（userLocalCurrencyAmount）", index+1)
	}

	// 使用传入的用户入账金额，向下取整
	userLocalAmount := item.UserLocalCurrencyAmount.Floor()
	convertedAmount := userLocalAmount.StringFixed(8)

	// 校验面额规则
	if giftcard != nil && giftcard.ValuesConfig != "" {
		if err := e.validateDenominationDecimal(userLocalAmount, giftcard.ValuesConfig); err != nil {
			return nil, fmt.Errorf("第%d条核销记录金额 %s 不符合入账面额规则：%s", index+1, userLocalAmount.String(), err.Error())
		}
	}

	result.UserLocalAmount = userLocalAmount

	// 3. 计算平台入账金额和货币
	var platformSettlementAmount string
	platformSettlementCurrency := ctx.Order.CurrencyCode
	if region != nil && region.CurrencyCode != "" {
		platformSettlementCurrency = region.CurrencyCode
	}

	// 计算平台货币对美元汇率
	var platformToUsdRate string
	if platformSettlementCurrency == "USD" {
		platformToUsdRate = "1.00000000"
	} else {
		rate, err := e.getCurrencyRate(e.Orm, platformSettlementCurrency, "USD")
		if err != nil {
			e.Log.Errorf("buildWriteoffRecordDecimal get platform currency rate error: currency=%s, err=%s", platformSettlementCurrency, err)
			return nil, fmt.Errorf("第%d条核销记录获取平台货币 %s 对美元汇率失败", index+1, platformSettlementCurrency)
		}
		platformToUsdRate = rate
	}

	// 平台入账金额
	if !item.PlatformSettlementAmount.IsZero() {
		platformSettlementAmount = item.PlatformSettlementAmount.StringFixed(8)
	} else if item.RecognizedCardValue.GreaterThan(decimal.Zero) && discountRate.GreaterThan(decimal.Zero) {
		// 自动计算：卡片面值 × 折扣率
		platformCost := item.RecognizedCardValue.Mul(discountRate)
		platformSettlementAmount = platformCost.StringFixed(8)
	}

	// 4. 构建核销记录
	result.Record = models.OrdGiftcardWriteoffs{
		UserId:                     ctx.Order.UserId,
		OrderId:                    orderId,
		GiftCardId:                 item.GiftCardId,
		Status:                     item.Status,
		Remark:                     item.Remark,
		AdminRecognizedCode:        item.AdminRecognizedCode,
		PlatformSaleRate:           platformSaleRate,
		RecognizedCardValue:        item.RecognizedCardValue.StringFixed(8),
		FailureImageUrl:            item.FailureImageUrl,
		SupplierId:                 strconv.Itoa(item.SupplierId),
		ConfigRate:                 ctx.ConfigRate.StringFixed(8),
		UserLocalCurrencyAmount:    convertedAmount,
		UserCurrencyCode:           ctx.TargetCurrency,
		PlatformSettlementAmount:   platformSettlementAmount,
		PlatformSettlementCurrency: platformSettlementCurrency,
		PlatformToUsdRate:          platformToUsdRate,
	}
	result.Record.CreateBy = ctx.CreateBy

	return result, nil
}


// prepareOrderUpdates 准备订单状态更新数据
func (e *OrdGiftcardWriteoffs) prepareOrderUpdates(finalStatus int) map[string]interface{} {
	var orderStatus int
	var processingStatus int
	if finalStatus == 1 {
		orderStatus = 2      // 订单完成
		processingStatus = 3 // 处理完成
	} else {
		orderStatus = 4      // 订单驳回
		processingStatus = 2 // 处理取消
	}

	orderUpdates := map[string]interface{}{
		"status":                 orderStatus,
		"processing_status":      processingStatus,
		"processing_started_end": time.Now(),
	}

	if orderStatus == 2 {
		orderUpdates["completed_at"] = time.Now()
	}

	return orderUpdates
}

// calculateExperienceAmount 计算经验值金额（转换为USD）
func (e *OrdGiftcardWriteoffs) calculateExperienceAmount(totalAmount decimal.Decimal, targetCurrency string, isCrypto bool, tx *gorm.DB) (decimal.Decimal, error) {
	if isCrypto {
		return totalAmount, nil
	}

	if targetCurrency == "USD" {
		return totalAmount, nil
	}

	rate, err := e.getCurrencyRate(tx, targetCurrency, "USD")
	if err != nil {
		e.Log.Errorf("calculateExperienceAmount get currency rate error: currency=%s, err=%s", targetCurrency, err)
		return decimal.Zero, fmt.Errorf("获取货币 %s 对美元汇率失败", targetCurrency)
	}

	rateDecimal, err := decimal.NewFromString(rate)
	if err != nil {
		e.Log.Errorf("calculateExperienceAmount parse rate error: rate=%s, err=%s", rate, err)
		return decimal.Zero, fmt.Errorf("汇率格式错误：%s", rate)
	}

	return totalAmount.Mul(rateDecimal), nil
}

// ============ 用户余额入账（含会员等级返利） ============

// creditUserBalanceWithRebate 用户入账（含会员等级返利到冻结余额）
func (e *OrdGiftcardWriteoffs) creditUserBalanceWithRebate(tx *gorm.DB, userId int, amount decimal.Decimal, currencyCode string, isCrypto bool, frozenLimit decimal.Decimal, orderId int, createBy int) error {
	// 1. 获取用户信息及等级返利配置
	var user models.HsUsers
	err := tx.Where("id = ?", userId).First(&user).Error
	if err != nil {
		e.Log.Errorf("creditUserBalanceWithRebate get user error:%s", err)
		return errors.New("获取用户信息失败")
	}

	// 2. 获取用户等级返利比例
	rebateRateDecimal, err := e.getUserRebateRateDecimal(tx, &user)
	if err != nil {
		return err
	}

	// 3. 计算返利金额和可用金额
	rebateAmount := amount.Mul(rebateRateDecimal)
	availableAmount := amount.Sub(rebateAmount)

	// 4. 根据法币/虚拟币分别处理
	if isCrypto {
		return e.creditCryptoBalanceWithRebateDecimal(tx, &user, availableAmount, rebateAmount, currencyCode, frozenLimit, orderId, createBy)
	}
	return e.creditFiatBalanceWithRebateDecimal(tx, &user, availableAmount, rebateAmount, currencyCode, frozenLimit, orderId, createBy)
}

// getUserRebateRate 获取用户等级返利比例
func (e *OrdGiftcardWriteoffs) getUserRebateRate(tx *gorm.DB, user *models.HsUsers) (float64, error) {
	// 用户没有等级，返利比例为0
	if user.LevelId == "" || user.LevelId == "0" {
		return 0, nil
	}

	// 用户有等级，必须找到等级配置
	var userLevel models.HsConfgiUserLevels
	err := tx.Where("id = ? AND is_active = 1", user.LevelId).First(&userLevel).Error
	if err != nil {
		e.Log.Errorf("getUserRebateRate get user level error: levelId=%s, err=%s", user.LevelId, err)
		return 0, fmt.Errorf("用户等级配置不存在或未启用：LevelID=%s", user.LevelId)
	}

	if userLevel.RebateRate == "" {
		return 0, nil
	}

	rebateRate, err := strconv.ParseFloat(userLevel.RebateRate, 64)
	if err != nil {
		e.Log.Errorf("getUserRebateRate parse rebate rate error: rate=%s, err=%s", userLevel.RebateRate, err)
		return 0, fmt.Errorf("用户等级返利比例格式错误：%s", userLevel.RebateRate)
	}

	// 如果返利比例大于1，说明是百分比形式，需要转换
	if rebateRate > 1 {
		rebateRate = rebateRate / 100.0
	}

	return rebateRate, nil
}

// getUserRebateRateDecimal 获取用户等级返利比例（返回decimal类型）
func (e *OrdGiftcardWriteoffs) getUserRebateRateDecimal(tx *gorm.DB, user *models.HsUsers) (decimal.Decimal, error) {
	// 用户没有等级，返利比例为0
	if user.LevelId == "" || user.LevelId == "0" {
		return decimal.Zero, nil
	}

	// 用户有等级，必须找到等级配置
	var userLevel models.HsConfgiUserLevels
	err := tx.Where("id = ? AND is_active = 1", user.LevelId).First(&userLevel).Error
	if err != nil {
		e.Log.Errorf("getUserRebateRateDecimal get user level error: levelId=%s, err=%s", user.LevelId, err)
		return decimal.Zero, fmt.Errorf("用户等级配置不存在或未启用：LevelID=%s", user.LevelId)
	}

	if userLevel.RebateRate == "" {
		return decimal.Zero, nil
	}

	rebateRate, err := decimal.NewFromString(userLevel.RebateRate)
	if err != nil {
		e.Log.Errorf("getUserRebateRateDecimal parse rebate rate error: rate=%s, err=%s", userLevel.RebateRate, err)
		return decimal.Zero, fmt.Errorf("用户等级返利比例格式错误：%s", userLevel.RebateRate)
	}

	// 如果返利比例大于1，说明是百分比形式，需要转换
	if rebateRate.GreaterThan(decimal.NewFromInt(1)) {
		rebateRate = rebateRate.Div(decimal.NewFromInt(100))
	}

	return rebateRate, nil
}

// ============ 法币入账处理 ============

// creditFiatBalanceWithRebate 法币入账（含返利到冻结余额）
func (e *OrdGiftcardWriteoffs) creditFiatBalanceWithRebate(tx *gorm.DB, user *models.HsUsers, availableAmount, rebateAmount float64, currencyCode string, frozenLimit float64, orderId int, createBy int) error {
	const decimalPlaces = 2

	// 1. 获取当前余额
	balanceBefore, _ := strconv.ParseFloat(user.Balance, 64)
	frozenBalanceBefore, _ := strconv.ParseFloat(user.FrozenBalance, 64)

	// 2. 处理可用金额入账
	balanceAfter := balanceBefore + availableAmount

	// 3. 处理返利金额到冻结余额（根据冻结配置）
	frozenAmount, overflowAmount := e.calculateFrozenAllocation(frozenBalanceBefore, rebateAmount, frozenLimit)
	frozenBalanceAfter := frozenBalanceBefore + frozenAmount
	balanceAfter += overflowAmount // 超出冻结限额的部分进入可用余额

	// 4. 更新用户余额（使用乐观锁）
	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", user.Id, user.Version).
		Updates(map[string]interface{}{
			"balance":        fmt.Sprintf("%.2f", balanceAfter),
			"frozen_balance": fmt.Sprintf("%.2f", frozenBalanceAfter),
			"version":        gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		e.Log.Errorf("creditFiatBalanceWithRebate update balance error:%s", result.Error)
		return errors.New("更新用户法币余额失败")
	}

	if result.RowsAffected == 0 {
		e.Log.Errorf("creditFiatBalanceWithRebate update balance conflict for user %d", user.Id)
		return errors.New("法币余额更新冲突，请重试")
	}

	// 5. 创建流水记录
	nanoTime := time.Now().UnixNano()

	// 5.1 可用余额流水（包含溢出的冻结金额）
	totalAvailable := availableAmount + overflowAmount
	if totalAvailable > 0 {
		err := e.createFiatLedger(tx, user.Id, totalAvailable, balanceBefore, balanceBefore+totalAvailable,
			currencyCode, "giftcard_writeoff_fiat", orderId, nanoTime, "礼品卡核销到账", createBy)
		if err != nil {
			return err
		}
	}

	// 5.2 冻结余额流水（返利部分）
	if frozenAmount > 0 {
		err := e.createFiatFrozenLedger(tx, user.Id, frozenAmount, frozenBalanceBefore, frozenBalanceAfter,
			currencyCode, "giftcard_writeoff_rebate_frozen", orderId, nanoTime+1, "礼品卡核销返利（冻结）", createBy)
		if err != nil {
			return err
		}
	}

	e.Log.Infof("User %d fiat credited: available=%.2f, rebate_frozen=%.2f, overflow=%.2f, currency=%s",
		user.Id, availableAmount, frozenAmount, overflowAmount, currencyCode)

	return nil
}

// createFiatLedger 创建法币流水记录
func (e *OrdGiftcardWriteoffs) createFiatLedger(tx *gorm.DB, userId int, amount, balanceBefore, balanceAfter float64, currencyCode, bizType string, orderId int, nanoTime int64, remark string, createBy int) error {
	ledger := models.HsUserLedger{
		UserId:         strconv.Itoa(userId),
		CurrencyCode:   currencyCode,
		Direction:      "1",
		Amount:         fmt.Sprintf("%.2f", amount),
		BalanceBefore:  fmt.Sprintf("%.2f", balanceBefore),
		BalanceAfter:   fmt.Sprintf("%.2f", balanceAfter),
		BizType:        bizType,
		BizId:          strconv.Itoa(orderId),
		IdempotencyKey: fmt.Sprintf("%s:%d:%d", bizType, orderId, nanoTime),
		RefTable:       "ord_giftcard_writeoffs",
		RefId:          strconv.Itoa(orderId),
		Remark:         fmt.Sprintf("%s，订单号: %d", remark, orderId),
		Status:         "1",
	}
	ledger.CreateBy = createBy

	if err := tx.Create(&ledger).Error; err != nil {
		e.Log.Errorf("createFiatLedger error:%s", err)
		return errors.New("创建法币流水记录失败")
	}
	return nil
}

// createFiatFrozenLedger 创建法币冻结流水记录
func (e *OrdGiftcardWriteoffs) createFiatFrozenLedger(tx *gorm.DB, userId int, amount, balanceBefore, balanceAfter float64, currencyCode, bizType string, orderId int, nanoTime int64, remark string, createBy int) error {
	ledger := models.HsUserFrozenLedger{
		UserId:         strconv.Itoa(userId),
		CurrencyCode:   currencyCode,
		Direction:      "1",
		Amount:         fmt.Sprintf("%.2f", amount),
		FrozenBefore:   fmt.Sprintf("%.2f", balanceBefore),
		FrozenAfter:    fmt.Sprintf("%.2f", balanceAfter),
		BizType:        bizType,
		BizId:          strconv.Itoa(orderId),
		IdempotencyKey: fmt.Sprintf("%s:%d:%d", bizType, orderId, nanoTime),
		Remark:         fmt.Sprintf("%s，订单号: %d，关联表: ord_giftcard_writeoffs", remark, orderId),
		Status:         "1",
	}
	ledger.CreateBy = createBy

	if err := tx.Create(&ledger).Error; err != nil {
		e.Log.Errorf("createFiatFrozenLedger error:%s", err)
		return errors.New("创建法币冻结流水记录失败")
	}
	return nil
}

// ============ 虚拟币入账处理 ============

// creditCryptoBalanceWithRebate 虚拟币入账（含返利到冻结余额）
func (e *OrdGiftcardWriteoffs) creditCryptoBalanceWithRebate(tx *gorm.DB, user *models.HsUsers, availableAmount, rebateAmount float64, currencyCode string, frozenLimit float64, orderId int, createBy int) error {
	const decimalPlaces = 8

	// 1. 获取当前余额
	balanceBefore, _ := strconv.ParseFloat(user.CryptoBalance, 64)
	frozenBalanceBefore, _ := strconv.ParseFloat(user.CryptoFrozenBalance, 64)

	// 2. 处理可用金额入账
	balanceAfter := balanceBefore + availableAmount

	// 3. 处理返利金额到冻结余额（根据冻结配置）
	frozenAmount, overflowAmount := e.calculateFrozenAllocation(frozenBalanceBefore, rebateAmount, frozenLimit)
	frozenBalanceAfter := frozenBalanceBefore + frozenAmount
	balanceAfter += overflowAmount // 超出冻结限额的部分进入可用余额

	// 4. 更新用户余额（使用乐观锁）
	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", user.Id, user.Version).
		Updates(map[string]interface{}{
			"crypto_balance":        fmt.Sprintf("%.8f", balanceAfter),
			"crypto_frozen_balance": fmt.Sprintf("%.8f", frozenBalanceAfter),
			"version":               gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		e.Log.Errorf("creditCryptoBalanceWithRebate update balance error:%s", result.Error)
		return errors.New("更新用户虚拟币余额失败")
	}

	if result.RowsAffected == 0 {
		e.Log.Errorf("creditCryptoBalanceWithRebate update balance conflict for user %d", user.Id)
		return errors.New("虚拟币余额更新冲突，请重试")
	}

	// 5. 创建流水记录
	nanoTime := time.Now().UnixNano()

	// 5.1 可用余额流水（包含溢出的冻结金额）
	totalAvailable := availableAmount + overflowAmount
	if totalAvailable > 0 {
		err := e.createCryptoLedger(tx, user.Id, totalAvailable, balanceBefore, balanceBefore+totalAvailable,
			currencyCode, "giftcard_writeoff_crypto", orderId, nanoTime, "礼品卡核销到账", createBy)
		if err != nil {
			return err
		}
	}

	// 5.2 冻结余额流水（返利部分）
	if frozenAmount > 0 {
		err := e.createCryptoFrozenLedger(tx, user.Id, frozenAmount, frozenBalanceBefore, frozenBalanceAfter,
			currencyCode, "giftcard_writeoff_rebate_frozen", orderId, nanoTime+1, "礼品卡核销返利（冻结）", createBy)
		if err != nil {
			return err
		}
	}

	e.Log.Infof("User %d crypto credited: available=%.8f, rebate_frozen=%.8f, overflow=%.8f, currency=%s",
		user.Id, availableAmount, frozenAmount, overflowAmount, currencyCode)

	return nil
}

// createCryptoLedger 创建虚拟币流水记录
func (e *OrdGiftcardWriteoffs) createCryptoLedger(tx *gorm.DB, userId int, amount, balanceBefore, balanceAfter float64, currencyCode, bizType string, orderId int, nanoTime int64, remark string, createBy int) error {
	ledger := models.HsUserLedger{
		UserId:         strconv.Itoa(userId),
		CurrencyCode:   currencyCode,
		Direction:      "1",
		Amount:         fmt.Sprintf("%.8f", amount),
		BalanceBefore:  fmt.Sprintf("%.8f", balanceBefore),
		BalanceAfter:   fmt.Sprintf("%.8f", balanceAfter),
		BizType:        bizType,
		BizId:          strconv.Itoa(orderId),
		IdempotencyKey: fmt.Sprintf("%s:%d:%d", bizType, orderId, nanoTime),
		RefTable:       "ord_giftcard_writeoffs",
		RefId:          strconv.Itoa(orderId),
		Remark:         fmt.Sprintf("%s，订单号: %d", remark, orderId),
		Status:         "1",
	}
	ledger.CreateBy = createBy

	if err := tx.Create(&ledger).Error; err != nil {
		e.Log.Errorf("createCryptoLedger error:%s", err)
		return errors.New("创建虚拟币流水记录失败")
	}
	return nil
}

// createCryptoFrozenLedger 创建虚拟币冻结流水记录
func (e *OrdGiftcardWriteoffs) createCryptoFrozenLedger(tx *gorm.DB, userId int, amount, balanceBefore, balanceAfter float64, currencyCode, bizType string, orderId int, nanoTime int64, remark string, createBy int) error {
	ledger := models.HsUserFrozenLedger{
		UserId:         strconv.Itoa(userId),
		CurrencyCode:   currencyCode,
		Direction:      "1",
		Amount:         fmt.Sprintf("%.8f", amount),
		FrozenBefore:   fmt.Sprintf("%.8f", balanceBefore),
		FrozenAfter:    fmt.Sprintf("%.8f", balanceAfter),
		BizType:        bizType,
		BizId:          strconv.Itoa(orderId),
		IdempotencyKey: fmt.Sprintf("%s:%d:%d", bizType, orderId, nanoTime),
		Remark:         fmt.Sprintf("%s，订单号: %d，关联表: ord_giftcard_writeoffs", remark, orderId),
		Status:         "1",
	}
	ledger.CreateBy = createBy

	if err := tx.Create(&ledger).Error; err != nil {
		e.Log.Errorf("createCryptoFrozenLedger error:%s", err)
		return errors.New("创建虚拟币冻结流水记录失败")
	}
	return nil
}

// ============ 法币/虚拟币入账处理（Decimal版本） ============

// creditFiatBalanceWithRebateDecimal 法币入账（使用decimal类型，含返利到冻结余额）
func (e *OrdGiftcardWriteoffs) creditFiatBalanceWithRebateDecimal(tx *gorm.DB, user *models.HsUsers, availableAmount, rebateAmount decimal.Decimal, currencyCode string, frozenLimit decimal.Decimal, orderId int, createBy int) error {
	// 1. 获取当前余额
	balanceBefore, _ := decimal.NewFromString(user.Balance)
	frozenBalanceBefore, _ := decimal.NewFromString(user.FrozenBalance)

	// 2. 处理可用金额入账
	balanceAfter := balanceBefore.Add(availableAmount)

	// 3. 处理返利金额到冻结余额（根据冻结配置）
	frozenAmount, overflowAmount := e.calculateFrozenAllocationDecimal(frozenBalanceBefore, rebateAmount, frozenLimit)
	frozenBalanceAfter := frozenBalanceBefore.Add(frozenAmount)
	balanceAfter = balanceAfter.Add(overflowAmount) // 超出冻结限额的部分进入可用余额

	// 4. 更新用户余额（使用乐观锁）
	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", user.Id, user.Version).
		Updates(map[string]interface{}{
			"balance":        balanceAfter.StringFixed(2),
			"frozen_balance": frozenBalanceAfter.StringFixed(2),
			"version":        gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		e.Log.Errorf("creditFiatBalanceWithRebateDecimal update balance error:%s", result.Error)
		return errors.New("更新用户法币余额失败")
	}

	if result.RowsAffected == 0 {
		e.Log.Errorf("creditFiatBalanceWithRebateDecimal update balance conflict for user %d", user.Id)
		return errors.New("法币余额更新冲突，请重试")
	}

	// 5. 创建流水记录
	nanoTime := time.Now().UnixNano()

	// 5.1 可用余额流水（包含溢出的冻结金额）
	totalAvailable := availableAmount.Add(overflowAmount)
	if totalAvailable.GreaterThan(decimal.Zero) {
		err := e.createFiatLedgerDecimal(tx, user.Id, totalAvailable, balanceBefore, balanceBefore.Add(totalAvailable),
			currencyCode, "giftcard_writeoff_fiat", orderId, nanoTime, "礼品卡核销到账", createBy)
		if err != nil {
			return err
		}
	}

	// 5.2 冻结余额流水（返利部分）
	if frozenAmount.GreaterThan(decimal.Zero) {
		err := e.createFiatFrozenLedgerDecimal(tx, user.Id, frozenAmount, frozenBalanceBefore, frozenBalanceAfter,
			currencyCode, "giftcard_writeoff_rebate_frozen", orderId, nanoTime+1, "礼品卡核销返利（冻结）", createBy)
		if err != nil {
			return err
		}
	}

	e.Log.Infof("User %d fiat credited: available=%s, rebate_frozen=%s, overflow=%s, currency=%s",
		user.Id, availableAmount.String(), frozenAmount.String(), overflowAmount.String(), currencyCode)

	return nil
}

// createFiatLedgerDecimal 创建法币流水记录（使用decimal类型）
func (e *OrdGiftcardWriteoffs) createFiatLedgerDecimal(tx *gorm.DB, userId int, amount, balanceBefore, balanceAfter decimal.Decimal, currencyCode, bizType string, orderId int, nanoTime int64, remark string, createBy int) error {
	ledger := models.HsUserLedger{
		UserId:         strconv.Itoa(userId),
		CurrencyCode:   currencyCode,
		Direction:      "1",
		Amount:         amount.StringFixed(2),
		BalanceBefore:  balanceBefore.StringFixed(2),
		BalanceAfter:   balanceAfter.StringFixed(2),
		BizType:        bizType,
		BizId:          strconv.Itoa(orderId),
		IdempotencyKey: fmt.Sprintf("%s:%d:%d", bizType, orderId, nanoTime),
		RefTable:       "ord_giftcard_writeoffs",
		RefId:          strconv.Itoa(orderId),
		Remark:         fmt.Sprintf("%s，订单号: %d", remark, orderId),
		Status:         "1",
	}
	ledger.CreateBy = createBy

	if err := tx.Create(&ledger).Error; err != nil {
		e.Log.Errorf("createFiatLedgerDecimal error:%s", err)
		return errors.New("创建法币流水记录失败")
	}
	return nil
}

// createFiatFrozenLedgerDecimal 创建法币冻结流水记录（使用decimal类型）
func (e *OrdGiftcardWriteoffs) createFiatFrozenLedgerDecimal(tx *gorm.DB, userId int, amount, balanceBefore, balanceAfter decimal.Decimal, currencyCode, bizType string, orderId int, nanoTime int64, remark string, createBy int) error {
	ledger := models.HsUserFrozenLedger{
		UserId:         strconv.Itoa(userId),
		CurrencyCode:   currencyCode,
		Direction:      "1",
		Amount:         amount.StringFixed(2),
		FrozenBefore:   balanceBefore.StringFixed(2),
		FrozenAfter:    balanceAfter.StringFixed(2),
		BizType:        bizType,
		BizId:          strconv.Itoa(orderId),
		IdempotencyKey: fmt.Sprintf("%s:%d:%d", bizType, orderId, nanoTime),
		Remark:         fmt.Sprintf("%s，订单号: %d，关联表: ord_giftcard_writeoffs", remark, orderId),
		Status:         "1",
	}
	ledger.CreateBy = createBy

	if err := tx.Create(&ledger).Error; err != nil {
		e.Log.Errorf("createFiatFrozenLedgerDecimal error:%s", err)
		return errors.New("创建法币冻结流水记录失败")
	}
	return nil
}

// createFiatFrozenLedgerOutDecimal 创建法币冻结余额转出流水记录（使用decimal类型）
func (e *OrdGiftcardWriteoffs) createFiatFrozenLedgerOutDecimal(tx *gorm.DB, userId int, amount, balanceBefore, balanceAfter decimal.Decimal, currencyCode, bizType string, orderId int, nanoTime int64, remark string, createBy int) error {
	ledger := models.HsUserFrozenLedger{
		UserId:         strconv.Itoa(userId),
		CurrencyCode:   currencyCode,
		Direction:      "2", // 转出
		Amount:         amount.StringFixed(2),
		FrozenBefore:   balanceBefore.StringFixed(2),
		FrozenAfter:    balanceAfter.StringFixed(2),
		BizType:        bizType,
		BizId:          strconv.Itoa(orderId),
		IdempotencyKey: fmt.Sprintf("%s:%d:%d", bizType, orderId, nanoTime),
		Remark:         fmt.Sprintf("%s，订单号: %d，关联表: ord_giftcard_writeoffs", remark, orderId),
		Status:         "1",
	}
	ledger.CreateBy = createBy

	if err := tx.Create(&ledger).Error; err != nil {
		e.Log.Errorf("createFiatFrozenLedgerOutDecimal error:%s", err)
		return errors.New("创建法币冻结转出流水记录失败")
	}
	return nil
}

// creditCryptoBalanceWithRebateDecimal 虚拟币入账（使用decimal类型，含返利到冻结余额）
func (e *OrdGiftcardWriteoffs) creditCryptoBalanceWithRebateDecimal(tx *gorm.DB, user *models.HsUsers, availableAmount, rebateAmount decimal.Decimal, currencyCode string, frozenLimit decimal.Decimal, orderId int, createBy int) error {
	// 1. 获取当前余额
	balanceBefore, _ := decimal.NewFromString(user.CryptoBalance)
	frozenBalanceBefore, _ := decimal.NewFromString(user.CryptoFrozenBalance)

	// 2. 处理可用金额入账
	balanceAfter := balanceBefore.Add(availableAmount)

	// 3. 处理返利金额到冻结余额（根据冻结配置）
	frozenAmount, overflowAmount := e.calculateFrozenAllocationDecimal(frozenBalanceBefore, rebateAmount, frozenLimit)
	frozenBalanceAfter := frozenBalanceBefore.Add(frozenAmount)
	balanceAfter = balanceAfter.Add(overflowAmount) // 超出冻结限额的部分进入可用余额

	// 4. 更新用户余额（使用乐观锁）
	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", user.Id, user.Version).
		Updates(map[string]interface{}{
			"crypto_balance":        balanceAfter.StringFixed(8),
			"crypto_frozen_balance": frozenBalanceAfter.StringFixed(8),
			"version":               gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		e.Log.Errorf("creditCryptoBalanceWithRebateDecimal update balance error:%s", result.Error)
		return errors.New("更新用户虚拟币余额失败")
	}

	if result.RowsAffected == 0 {
		e.Log.Errorf("creditCryptoBalanceWithRebateDecimal update balance conflict for user %d", user.Id)
		return errors.New("虚拟币余额更新冲突，请重试")
	}

	// 5. 创建流水记录
	nanoTime := time.Now().UnixNano()

	// 5.1 可用余额流水（包含溢出的冻结金额）
	totalAvailable := availableAmount.Add(overflowAmount)
	if totalAvailable.GreaterThan(decimal.Zero) {
		err := e.createCryptoLedgerDecimal(tx, user.Id, totalAvailable, balanceBefore, balanceBefore.Add(totalAvailable),
			currencyCode, "giftcard_writeoff_crypto", orderId, nanoTime, "礼品卡核销到账", createBy)
		if err != nil {
			return err
		}
	}

	// 5.2 冻结余额流水（返利部分）
	if frozenAmount.GreaterThan(decimal.Zero) {
		err := e.createCryptoFrozenLedgerDecimal(tx, user.Id, frozenAmount, frozenBalanceBefore, frozenBalanceAfter,
			currencyCode, "giftcard_writeoff_rebate_frozen", orderId, nanoTime+1, "礼品卡核销返利（冻结）", createBy)
		if err != nil {
			return err
		}
	}

	e.Log.Infof("User %d crypto credited: available=%s, rebate_frozen=%s, overflow=%s, currency=%s",
		user.Id, availableAmount.String(), frozenAmount.String(), overflowAmount.String(), currencyCode)

	return nil
}

// createCryptoLedgerDecimal 创建虚拟币流水记录（使用decimal类型）
func (e *OrdGiftcardWriteoffs) createCryptoLedgerDecimal(tx *gorm.DB, userId int, amount, balanceBefore, balanceAfter decimal.Decimal, currencyCode, bizType string, orderId int, nanoTime int64, remark string, createBy int) error {
	ledger := models.HsUserLedger{
		UserId:         strconv.Itoa(userId),
		CurrencyCode:   currencyCode,
		Direction:      "1",
		Amount:         amount.StringFixed(8),
		BalanceBefore:  balanceBefore.StringFixed(8),
		BalanceAfter:   balanceAfter.StringFixed(8),
		BizType:        bizType,
		BizId:          strconv.Itoa(orderId),
		IdempotencyKey: fmt.Sprintf("%s:%d:%d", bizType, orderId, nanoTime),
		RefTable:       "ord_giftcard_writeoffs",
		RefId:          strconv.Itoa(orderId),
		Remark:         fmt.Sprintf("%s，订单号: %d", remark, orderId),
		Status:         "1",
	}
	ledger.CreateBy = createBy

	if err := tx.Create(&ledger).Error; err != nil {
		e.Log.Errorf("createCryptoLedgerDecimal error:%s", err)
		return errors.New("创建虚拟币流水记录失败")
	}
	return nil
}

// createCryptoFrozenLedgerDecimal 创建虚拟币冻结流水记录（使用decimal类型）
func (e *OrdGiftcardWriteoffs) createCryptoFrozenLedgerDecimal(tx *gorm.DB, userId int, amount, balanceBefore, balanceAfter decimal.Decimal, currencyCode, bizType string, orderId int, nanoTime int64, remark string, createBy int) error {
	ledger := models.HsUserFrozenLedger{
		UserId:         strconv.Itoa(userId),
		CurrencyCode:   currencyCode,
		Direction:      "1",
		Amount:         amount.StringFixed(8),
		FrozenBefore:   balanceBefore.StringFixed(8),
		FrozenAfter:    balanceAfter.StringFixed(8),
		BizType:        bizType,
		BizId:          strconv.Itoa(orderId),
		IdempotencyKey: fmt.Sprintf("%s:%d:%d", bizType, orderId, nanoTime),
		Remark:         fmt.Sprintf("%s，订单号: %d，关联表: ord_giftcard_writeoffs", remark, orderId),
		Status:         "1",
	}
	ledger.CreateBy = createBy

	if err := tx.Create(&ledger).Error; err != nil {
		e.Log.Errorf("createCryptoFrozenLedgerDecimal error:%s", err)
		return errors.New("创建虚拟币冻结流水记录失败")
	}
	return nil
}

// createCryptoFrozenLedgerOutDecimal 创建虚拟币冻结余额转出流水记录（使用decimal类型）
func (e *OrdGiftcardWriteoffs) createCryptoFrozenLedgerOutDecimal(tx *gorm.DB, userId int, amount, balanceBefore, balanceAfter decimal.Decimal, currencyCode, bizType string, orderId int, nanoTime int64, remark string, createBy int) error {
	ledger := models.HsUserFrozenLedger{
		UserId:         strconv.Itoa(userId),
		CurrencyCode:   currencyCode,
		Direction:      "2", // 转出
		Amount:         amount.StringFixed(8),
		FrozenBefore:   balanceBefore.StringFixed(8),
		FrozenAfter:    balanceAfter.StringFixed(8),
		BizType:        bizType,
		BizId:          strconv.Itoa(orderId),
		IdempotencyKey: fmt.Sprintf("%s:%d:%d", bizType, orderId, nanoTime),
		Remark:         fmt.Sprintf("%s，订单号: %d，关联表: ord_giftcard_writeoffs", remark, orderId),
		Status:         "1",
	}
	ledger.CreateBy = createBy

	if err := tx.Create(&ledger).Error; err != nil {
		e.Log.Errorf("createCryptoFrozenLedgerOutDecimal error:%s", err)
		return errors.New("创建虚拟币冻结转出流水记录失败")
	}
	return nil
}

// ============ 冻结余额配置处理 ============

// calculateFrozenAllocation 根据冻结配置计算冻结分配
// 返回：frozenAmount（实际冻结金额）, overflowAmount（超出限额进入可用的金额）
func (e *OrdGiftcardWriteoffs) calculateFrozenAllocation(currentFrozenBalance, amount, frozenLimit float64) (float64, float64) {
	if amount <= 0 {
		return 0, 0
	}

	if frozenLimit <= 0 {
		// 无限额配置，全部进入冻结余额
		return amount, 0
	}

	// 使用取模计算：冻结余额 = (当前冻结 + 新增) % 限额
	totalFrozen := currentFrozenBalance + amount
	newFrozenBalance := math.Mod(totalFrozen, frozenLimit)
	frozenDelta := newFrozenBalance - currentFrozenBalance
	toAvailableAmount := amount - frozenDelta

	e.Log.Infof("Frozen allocation: current=%.8f, add=%.8f, limit=%.8f -> frozenDelta=%.8f, toAvailable=%.8f",
		currentFrozenBalance, amount, frozenLimit, frozenDelta, toAvailableAmount)

	return frozenDelta, toAvailableAmount
}

// calculateFrozenAllocationDecimal 根据冻结配置计算冻结分配（使用decimal类型）
// 逻辑：返利先进入冻结余额，当冻结余额达到配置阈值时，将配置金额转入可用余额
// 示例：配置阈值3000，当前冻结0，新增10000 → 冻结剩余1000(10000%3000)，可用增加9000
// 返回：frozenDelta（冻结余额变化量，可为负）, toAvailableAmount（转入可用余额的金额）
func (e *OrdGiftcardWriteoffs) calculateFrozenAllocationDecimal(currentFrozenBalance, amount, frozenLimit decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	if amount.LessThanOrEqual(decimal.Zero) {
		return decimal.Zero, decimal.Zero
	}

	if frozenLimit.LessThanOrEqual(decimal.Zero) {
		// 无限额配置，全部进入冻结余额
		return amount, decimal.Zero
	}

	// 1. 计算总冻结余额
	totalFrozen := currentFrozenBalance.Add(amount)

	// 2. 使用取模计算最终冻结余额（totalFrozen % frozenLimit）
	newFrozenBalance := totalFrozen.Mod(frozenLimit)

	// 3. 计算冻结余额变化量（可能为负，表示从冻结转出到可用）
	// 例：当前2500，返利1000，阈值3000 → 最终500，变化量=-2000
	frozenDelta := newFrozenBalance.Sub(currentFrozenBalance)

	// 4. 转入可用余额的金额 = 返利金额 - 冻结变化量
	// 例：返利1000 - (-2000) = 3000（包含返利+释放的冻结）
	toAvailableAmount := amount.Sub(frozenDelta)

	e.Log.Infof("Frozen allocation: current=%s, add=%s, limit=%s -> newFrozen=%s, frozenDelta=%s, toAvailable=%s",
		currentFrozenBalance.String(), amount.String(), frozenLimit.String(),
		newFrozenBalance.String(), frozenDelta.String(), toAvailableAmount.String())

	return frozenDelta, toAvailableAmount
}

// ============ 汇率查询 ============

// getCurrencyRate 获取货币汇率
func (e *OrdGiftcardWriteoffs) getCurrencyRate(tx *gorm.DB, fromCurrency, toCurrency string) (string, error) {
	if fromCurrency == toCurrency {
		return "1.00000000", nil
	}

	var rateRecord models.OrdConfigCurrencyRates
	err := tx.Where("base_currency_code = ? AND quote_currency_code = ? AND status = 1",
		fromCurrency, toCurrency).
		First(&rateRecord).Error
	if err != nil {
		e.Log.Errorf("Get currency rate %s to %s error:%s", fromCurrency, toCurrency, err)
		return "", fmt.Errorf("未找到 %s 到 %s 的有效汇率配置", fromCurrency, toCurrency)
	}
	return rateRecord.Rate, nil
}

// getCurrencyRateViaUSD 直接获取货币汇率
func (e *OrdGiftcardWriteoffs) getCurrencyRateViaUSD(tx *gorm.DB, fromCurrency, toCurrency string) (string, error) {
	if fromCurrency == toCurrency {
		return "1.00000000", nil
	}
	return e.getCurrencyRate(tx, fromCurrency, toCurrency)
}

// ============ 邀请分成处理 ============

// InviterCommissionConfig 邀请人分成配置
type InviterCommissionConfig struct {
	CommissionRate float64 // 分成比例
	FrozenRate     float64 // 进入冻结余额的比例
}

// processInviteCommissions 处理邀请分成
func (e *OrdGiftcardWriteoffs) processInviteCommissions(tx *gorm.DB, userId, orderId int, amount float64, sourceCurrency string, isCrypto bool, frozenLimit float64, createBy int) error {
	orderIdStr := strconv.Itoa(orderId)

	// 1. 查询邀请关系
	var inviteRelation models.HsInviteRelations
	err := tx.Where("user_id = ?", userId).First(&inviteRelation).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	// 2. 处理一级邀请人分成
	if inviteRelation.Level1InviterId != "" && inviteRelation.Level1InviterId != "0" {
		err = e.processInviterCommission(tx, inviteRelation.Level1InviterId, 1, amount, sourceCurrency, isCrypto, frozenLimit, orderIdStr, createBy)
		if err != nil {
			return err
		}
	}

	// 3. 处理二级邀请人分成
	if inviteRelation.Level2InviterId != "" && inviteRelation.Level2InviterId != "0" {
		err = e.processInviterCommission(tx, inviteRelation.Level2InviterId, 2, amount, sourceCurrency, isCrypto, frozenLimit, orderIdStr, createBy)
		if err != nil {
			return err
		}
	}

	return nil
}

// processInviterCommission 处理单个邀请人的分成
func (e *OrdGiftcardWriteoffs) processInviterCommission(tx *gorm.DB, inviterId string, level int, amount float64, sourceCurrency string, isCrypto bool, frozenLimit float64, orderIdStr string, createBy int) error {
	config, err := e.getInviterCommissionConfig(tx, inviterId, level)
	if err != nil {
		e.Log.Errorf("Get level%d inviter commission config error:%s", level, err)
		return err
	}

	if config.CommissionRate <= 0 {
		return nil
	}

	commissionAmount := amount * config.CommissionRate

	// 创建分成记录
	commission := models.HsInviteCommissions{
		OrderId:          orderIdStr,
		UserId:           inviterId,
		CommissionLevel:  strconv.Itoa(level),
		CommissionRate:   fmt.Sprintf("%.2f", config.CommissionRate),
		CommissionAmount: fmt.Sprintf("%.8f", commissionAmount),
		Status:           "1",
	}
	commission.CreateBy = createBy

	if err = tx.Create(&commission).Error; err != nil {
		e.Log.Errorf("Create level%d commission error:%s", level, err)
		return fmt.Errorf("创建%d级分成记录失败", level)
	}

	// 根据frozen_rate分配
	frozenAmount := commissionAmount * config.FrozenRate
	availableAmount := commissionAmount * (1 - config.FrozenRate)
	orderId, _ := strconv.Atoi(orderIdStr)

	if frozenAmount > 0 {
		err = e.updateInviterFrozenBalance(tx, inviterId, frozenAmount, sourceCurrency, isCrypto, frozenLimit, orderId, strconv.Itoa(level), createBy)
		if err != nil {
			return err
		}
	}

	if availableAmount > 0 {
		err = e.updateInviterAvailableBalance(tx, inviterId, availableAmount, sourceCurrency, isCrypto, orderId, strconv.Itoa(level), createBy)
		if err != nil {
			return err
		}
	}

	return nil
}

// getInviterCommissionConfig 获取邀请人分成配置
func (e *OrdGiftcardWriteoffs) getInviterCommissionConfig(tx *gorm.DB, inviterId string, level int) (*InviterCommissionConfig, error) {
	var commissionConfig models.HsConfigInviteCommission
	err := tx.Where("config_name = ? AND status = 1", "invite_config").First(&commissionConfig).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("邀请分成配置不存在或未启用")
		}
		return nil, fmt.Errorf("获取分成配置失败: %s", err.Error())
	}

	var commissionRateStr string
	if level == 1 {
		commissionRateStr = commissionConfig.FirstLevelRate
	} else if level == 2 {
		commissionRateStr = commissionConfig.SecondLevelRate
	} else {
		return nil, fmt.Errorf("无效的分成层级: %d", level)
	}

	if commissionRateStr == "" {
		return nil, fmt.Errorf("分成比例配置为空（%d级）", level)
	}

	commissionRate, err := strconv.ParseFloat(commissionRateStr, 64)
	if err != nil {
		e.Log.Errorf("getInviterCommissionConfig parse commission rate error: rate=%s, err=%s", commissionRateStr, err)
		return nil, fmt.Errorf("分成比例格式错误：%s", commissionRateStr)
	}
	commissionRate = commissionRate / 100.0

	// 获取邀请人信息
	var inviter models.HsUsers
	err = tx.Where("id = ?", inviterId).First(&inviter).Error
	if err != nil {
		e.Log.Errorf("getInviterCommissionConfig get inviter error: inviterId=%s, err=%s", inviterId, err)
		return nil, fmt.Errorf("获取邀请人信息失败：InviterID=%s", inviterId)
	}

	// 获取邀请人等级返利比例
	var frozenRate float64 = 1.0
	if inviter.LevelId != "" && inviter.LevelId != "0" {
		var userLevel models.HsConfgiUserLevels
		err = tx.Where("id = ? AND is_active = 1", inviter.LevelId).First(&userLevel).Error
		if err != nil {
			e.Log.Errorf("getInviterCommissionConfig get inviter level error: levelId=%s, err=%s", inviter.LevelId, err)
			return nil, fmt.Errorf("邀请人等级配置不存在或未启用：LevelID=%s", inviter.LevelId)
		}
		if userLevel.RebateRate != "" {
			frozenRate, err = strconv.ParseFloat(userLevel.RebateRate, 64)
			if err != nil {
				e.Log.Errorf("getInviterCommissionConfig parse rebate rate error: rate=%s, err=%s", userLevel.RebateRate, err)
				return nil, fmt.Errorf("邀请人等级返利比例格式错误：%s", userLevel.RebateRate)
			}
		}
	}

	return &InviterCommissionConfig{
		CommissionRate: commissionRate,
		FrozenRate:     frozenRate,
	}, nil
}

// updateInviterFrozenBalance 更新邀请人冻结余额（法币/虚拟币分离）
func (e *OrdGiftcardWriteoffs) updateInviterFrozenBalance(tx *gorm.DB, inviterId string, sourceAmount float64, sourceCurrency string, isCrypto bool, frozenLimit float64, orderId int, level string, createBy int) error {
	if isCrypto {
		return e.updateInviterCryptoFrozenBalance(tx, inviterId, sourceAmount, sourceCurrency, frozenLimit, orderId, level, createBy)
	}
	return e.updateInviterFiatFrozenBalance(tx, inviterId, sourceAmount, sourceCurrency, frozenLimit, orderId, level, createBy)
}

// updateInviterFiatFrozenBalance 更新邀请人法币冻结余额
func (e *OrdGiftcardWriteoffs) updateInviterFiatFrozenBalance(tx *gorm.DB, inviterId string, amount float64, currency string, frozenLimit float64, orderId int, level string, createBy int) error {
	var inviter models.HsUsers
	err := tx.Where("id = ?", inviterId).First(&inviter).Error
	if err != nil {
		return fmt.Errorf("获取邀请人信息失败")
	}

	frozenBalanceBefore, _ := strconv.ParseFloat(inviter.FrozenBalance, 64)
	availableBalanceBefore, _ := strconv.ParseFloat(inviter.Balance, 64)

	// 计算冻结分配
	frozenAmount, overflowAmount := e.calculateFrozenAllocation(frozenBalanceBefore, amount, frozenLimit)
	frozenBalanceAfter := frozenBalanceBefore + frozenAmount
	availableBalanceAfter := availableBalanceBefore + overflowAmount

	// 更新余额
	updateFields := map[string]interface{}{"version": gorm.Expr("version + 1")}
	if frozenAmount > 0 {
		updateFields["frozen_balance"] = fmt.Sprintf("%.2f", frozenBalanceAfter)
	}
	if overflowAmount > 0 {
		updateFields["balance"] = fmt.Sprintf("%.2f", availableBalanceAfter)
	}

	result := tx.Model(&models.HsUsers{}).Where("id = ? AND version = ?", inviterId, inviter.Version).Updates(updateFields)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("更新邀请人法币余额失败")
	}

	// 创建流水
	nanoTime := time.Now().UnixNano()
	orderIdStr := strconv.Itoa(orderId)

	if frozenAmount > 0 {
		ledger := models.HsUserLedger{
			UserId: inviterId, CurrencyCode: currency, Direction: "1",
			Amount: fmt.Sprintf("%.2f", frozenAmount), BalanceBefore: fmt.Sprintf("%.2f", frozenBalanceBefore),
			BalanceAfter: fmt.Sprintf("%.2f", frozenBalanceAfter), BizType: "invite_commission_frozen_fiat",
			BizId: orderIdStr, IdempotencyKey: fmt.Sprintf("INVITE_FROZEN_L%s:%d:%d", level, orderId, nanoTime),
			RefTable: "hs_invite_commissions", RefId: orderIdStr,
			Remark: fmt.Sprintf("邀请分成冻结（%s级），订单号: %d", level, orderId), Status: "1",
		}
		ledger.CreateBy = createBy
		tx.Create(&ledger)
	}

	if overflowAmount > 0 {
		ledger := models.HsUserLedger{
			UserId: inviterId, CurrencyCode: currency, Direction: "1",
			Amount: fmt.Sprintf("%.2f", overflowAmount), BalanceBefore: fmt.Sprintf("%.2f", availableBalanceBefore),
			BalanceAfter: fmt.Sprintf("%.2f", availableBalanceAfter), BizType: "invite_commission_available_fiat",
			BizId: orderIdStr, IdempotencyKey: fmt.Sprintf("INVITE_AVAILABLE_L%s:%d:%d", level, orderId, nanoTime+1),
			RefTable: "hs_invite_commissions", RefId: orderIdStr,
			Remark: fmt.Sprintf("邀请分成可用（%s级，超冻结限额），订单号: %d", level, orderId), Status: "1",
		}
		ledger.CreateBy = createBy
		tx.Create(&ledger)
	}

	return nil
}

// updateInviterCryptoFrozenBalance 更新邀请人虚拟币冻结余额
func (e *OrdGiftcardWriteoffs) updateInviterCryptoFrozenBalance(tx *gorm.DB, inviterId string, amount float64, currency string, frozenLimit float64, orderId int, level string, createBy int) error {
	var inviter models.HsUsers
	err := tx.Where("id = ?", inviterId).First(&inviter).Error
	if err != nil {
		return fmt.Errorf("获取邀请人信息失败")
	}

	frozenBalanceBefore, _ := strconv.ParseFloat(inviter.CryptoFrozenBalance, 64)
	availableBalanceBefore, _ := strconv.ParseFloat(inviter.CryptoBalance, 64)

	// 计算冻结分配
	frozenAmount, overflowAmount := e.calculateFrozenAllocation(frozenBalanceBefore, amount, frozenLimit)
	frozenBalanceAfter := frozenBalanceBefore + frozenAmount
	availableBalanceAfter := availableBalanceBefore + overflowAmount

	// 更新余额
	updateFields := map[string]interface{}{"version": gorm.Expr("version + 1")}
	if frozenAmount > 0 {
		updateFields["crypto_frozen_balance"] = fmt.Sprintf("%.8f", frozenBalanceAfter)
	}
	if overflowAmount > 0 {
		updateFields["crypto_balance"] = fmt.Sprintf("%.8f", availableBalanceAfter)
	}

	result := tx.Model(&models.HsUsers{}).Where("id = ? AND version = ?", inviterId, inviter.Version).Updates(updateFields)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("更新邀请人虚拟币余额失败")
	}

	// 创建流水
	nanoTime := time.Now().UnixNano()
	orderIdStr := strconv.Itoa(orderId)

	if frozenAmount > 0 {
		ledger := models.HsUserLedger{
			UserId: inviterId, CurrencyCode: currency, Direction: "1",
			Amount: fmt.Sprintf("%.8f", frozenAmount), BalanceBefore: fmt.Sprintf("%.8f", frozenBalanceBefore),
			BalanceAfter: fmt.Sprintf("%.8f", frozenBalanceAfter), BizType: "invite_commission_frozen_crypto",
			BizId: orderIdStr, IdempotencyKey: fmt.Sprintf("INVITE_FROZEN_L%s:%d:%d", level, orderId, nanoTime),
			RefTable: "hs_invite_commissions", RefId: orderIdStr,
			Remark: fmt.Sprintf("邀请分成冻结（%s级），订单号: %d", level, orderId), Status: "1",
		}
		ledger.CreateBy = createBy
		tx.Create(&ledger)
	}

	if overflowAmount > 0 {
		ledger := models.HsUserLedger{
			UserId: inviterId, CurrencyCode: currency, Direction: "1",
			Amount: fmt.Sprintf("%.8f", overflowAmount), BalanceBefore: fmt.Sprintf("%.8f", availableBalanceBefore),
			BalanceAfter: fmt.Sprintf("%.8f", availableBalanceAfter), BizType: "invite_commission_available_crypto",
			BizId: orderIdStr, IdempotencyKey: fmt.Sprintf("INVITE_AVAILABLE_L%s:%d:%d", level, orderId, nanoTime+1),
			RefTable: "hs_invite_commissions", RefId: orderIdStr,
			Remark: fmt.Sprintf("邀请分成可用（%s级，超冻结限额），订单号: %d", level, orderId), Status: "1",
		}
		ledger.CreateBy = createBy
		tx.Create(&ledger)
	}

	return nil
}

// updateInviterAvailableBalance 更新邀请人可用余额（法币/虚拟币分离）
func (e *OrdGiftcardWriteoffs) updateInviterAvailableBalance(tx *gorm.DB, inviterId string, sourceAmount float64, sourceCurrency string, isCrypto bool, orderId int, level string, createBy int) error {
	if isCrypto {
		return e.updateInviterCryptoAvailableBalance(tx, inviterId, sourceAmount, sourceCurrency, orderId, level, createBy)
	}
	return e.updateInviterFiatAvailableBalance(tx, inviterId, sourceAmount, sourceCurrency, orderId, level, createBy)
}

// updateInviterFiatAvailableBalance 更新邀请人法币可用余额
func (e *OrdGiftcardWriteoffs) updateInviterFiatAvailableBalance(tx *gorm.DB, inviterId string, amount float64, currency string, orderId int, level string, createBy int) error {
	var inviter models.HsUsers
	err := tx.Where("id = ?", inviterId).First(&inviter).Error
	if err != nil {
		return fmt.Errorf("获取邀请人信息失败")
	}

	balanceBefore, _ := strconv.ParseFloat(inviter.Balance, 64)
	balanceAfter := balanceBefore + amount

	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", inviterId, inviter.Version).
		Updates(map[string]interface{}{
			"balance": fmt.Sprintf("%.2f", balanceAfter),
			"version": gorm.Expr("version + 1"),
		})

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("更新邀请人法币可用余额失败")
	}

	orderIdStr := strconv.Itoa(orderId)
	ledger := models.HsUserLedger{
		UserId: inviterId, CurrencyCode: currency, Direction: "1",
		Amount: fmt.Sprintf("%.2f", amount), BalanceBefore: fmt.Sprintf("%.2f", balanceBefore),
		BalanceAfter: fmt.Sprintf("%.2f", balanceAfter), BizType: "invite_commission_available_fiat",
		BizId: orderIdStr, IdempotencyKey: fmt.Sprintf("INVITE_AVAILABLE_L%s:%d:%d", level, orderId, time.Now().UnixNano()),
		RefTable: "hs_invite_commissions", RefId: orderIdStr,
		Remark: fmt.Sprintf("邀请分成可用（%s级），订单号: %d", level, orderId), Status: "1",
	}
	ledger.CreateBy = createBy
	return tx.Create(&ledger).Error
}

// updateInviterCryptoAvailableBalance 更新邀请人虚拟币可用余额
func (e *OrdGiftcardWriteoffs) updateInviterCryptoAvailableBalance(tx *gorm.DB, inviterId string, amount float64, currency string, orderId int, level string, createBy int) error {
	var inviter models.HsUsers
	err := tx.Where("id = ?", inviterId).First(&inviter).Error
	if err != nil {
		return fmt.Errorf("获取邀请人信息失败")
	}

	balanceBefore, _ := strconv.ParseFloat(inviter.CryptoBalance, 64)
	balanceAfter := balanceBefore + amount

	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", inviterId, inviter.Version).
		Updates(map[string]interface{}{
			"crypto_balance": fmt.Sprintf("%.8f", balanceAfter),
			"version":        gorm.Expr("version + 1"),
		})

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("更新邀请人虚拟币可用余额失败")
	}

	orderIdStr := strconv.Itoa(orderId)
	ledger := models.HsUserLedger{
		UserId: inviterId, CurrencyCode: currency, Direction: "1",
		Amount: fmt.Sprintf("%.8f", amount), BalanceBefore: fmt.Sprintf("%.8f", balanceBefore),
		BalanceAfter: fmt.Sprintf("%.8f", balanceAfter), BizType: "invite_commission_available_crypto",
		BizId: orderIdStr, IdempotencyKey: fmt.Sprintf("INVITE_AVAILABLE_L%s:%d:%d", level, orderId, time.Now().UnixNano()),
		RefTable: "hs_invite_commissions", RefId: orderIdStr,
		Remark: fmt.Sprintf("邀请分成可用（%s级），订单号: %d", level, orderId), Status: "1",
	}
	ledger.CreateBy = createBy
	return tx.Create(&ledger).Error
}

// ============ 面额校验 ============

// ValuesConfigStruct 面额配置结构
type ValuesConfigStruct struct {
	Fixed []float64 `json:"fixed"`
	Range *struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	} `json:"range"`
}

// validateDenomination 校验金额是否符合入账面额规则
func (e *OrdGiftcardWriteoffs) validateDenomination(amount float64, valuesConfig string) error {
	if valuesConfig == "" {
		return nil
	}

	var config ValuesConfigStruct
	if err := json.Unmarshal([]byte(valuesConfig), &config); err != nil {
		return nil
	}

	// 校验固定面额
	if len(config.Fixed) > 0 {
		for _, fixedValue := range config.Fixed {
			if fmt.Sprintf("%.2f", amount) == fmt.Sprintf("%.2f", fixedValue) {
				return nil
			}
		}
	}

	// 校验面额区间
	if config.Range != nil {
		if amount >= config.Range.Min && amount <= config.Range.Max {
			return nil
		}
	}

	if len(config.Fixed) == 0 && config.Range == nil {
		return nil
	}

	// 构建错误信息
	var validValues []string
	if len(config.Fixed) > 0 {
		fixedStrs := make([]string, len(config.Fixed))
		for i, v := range config.Fixed {
			fixedStrs[i] = fmt.Sprintf("%.2f", v)
		}
		validValues = append(validValues, fmt.Sprintf("固定面额: %v", fixedStrs))
	}
	if config.Range != nil {
		validValues = append(validValues, fmt.Sprintf("区间: %.2f - %.2f", config.Range.Min, config.Range.Max))
	}

	return fmt.Errorf("允许的面额为 %s", strings.Join(validValues, " 或 "))
}

// validateDenominationDecimal 校验金额是否符合入账面额规则（使用decimal类型）
func (e *OrdGiftcardWriteoffs) validateDenominationDecimal(amount decimal.Decimal, valuesConfig string) error {
	if valuesConfig == "" {
		return nil
	}

	var config ValuesConfigStruct
	if err := json.Unmarshal([]byte(valuesConfig), &config); err != nil {
		return nil
	}

	// 校验固定面额
	if len(config.Fixed) > 0 {
		for _, fixedValue := range config.Fixed {
			fixedDecimal := decimal.NewFromFloat(fixedValue)
			if amount.Equal(fixedDecimal) {
				return nil
			}
		}
	}

	// 校验面额区间
	if config.Range != nil {
		minDecimal := decimal.NewFromFloat(config.Range.Min)
		maxDecimal := decimal.NewFromFloat(config.Range.Max)
		if amount.GreaterThanOrEqual(minDecimal) && amount.LessThanOrEqual(maxDecimal) {
			return nil
		}
	}

	if len(config.Fixed) == 0 && config.Range == nil {
		return nil
	}

	// 构建错误信息
	var validValues []string
	if len(config.Fixed) > 0 {
		fixedStrs := make([]string, len(config.Fixed))
		for i, v := range config.Fixed {
			fixedStrs[i] = fmt.Sprintf("%.2f", v)
		}
		validValues = append(validValues, fmt.Sprintf("固定面额: %v", fixedStrs))
	}
	if config.Range != nil {
		validValues = append(validValues, fmt.Sprintf("区间: %.2f - %.2f", config.Range.Min, config.Range.Max))
	}

	return fmt.Errorf("允许的面额为 %s", strings.Join(validValues, " 或 "))
}

// ============ 经验值处理 ============

// addUserExperience 增加用户经验值（使用预查询的等级配置）
func (e *OrdGiftcardWriteoffs) addUserExperience(tx *gorm.DB, userId int, experienceChange int, levels []models.HsConfgiUserLevels, sourceType string, orderId int, createBy int) error {
	if experienceChange <= 0 {
		return nil
	}

	var user models.HsUsers
	err := tx.Select("id, level_id, experience, total_experience, version").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return errors.New("获取用户信息失败")
	}

	currentLevelId, _ := strconv.Atoi(user.LevelId)
	experienceBefore, _ := strconv.Atoi(user.Experience)
	totalExperienceBefore, _ := strconv.Atoi(user.TotalExperience)
	experienceAfter := experienceBefore + experienceChange
	totalExperienceAfter := totalExperienceBefore + experienceChange

	// 检查是否需要升级（使用预查询的等级配置）
	newLevelId := e.checkUserLevelUpgrade(currentLevelId, totalExperienceAfter, levels)

	// 构建更新字段
	updates := map[string]interface{}{
		"experience":       strconv.Itoa(experienceAfter),
		"total_experience": strconv.Itoa(totalExperienceAfter),
		"version":          gorm.Expr("version + 1"),
	}

	// 如果需要升级，同时更新等级
	if newLevelId != currentLevelId {
		updates["level_id"] = strconv.Itoa(newLevelId)
		e.Log.Infof("User %d level upgrade: %d -> %d (totalExp: %d)", userId, currentLevelId, newLevelId, totalExperienceAfter)
	}

	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", userId, user.Version).
		Updates(updates)

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("更新用户经验值失败")
	}

	experienceLog := models.HsUserExperienceLogs{
		UserId:           strconv.Itoa(userId),
		ExperienceChange: strconv.Itoa(experienceChange),
		ExperienceBefore: strconv.Itoa(experienceBefore),
		ExperienceAfter:  strconv.Itoa(experienceAfter),
		SourceType:       sourceType,
		SourceId:         strconv.Itoa(orderId),
		Description:      fmt.Sprintf("礼品卡核销获得经验值，订单号: %d", orderId),
	}
	experienceLog.CreateBy = createBy

	return tx.Create(&experienceLog).Error
}

// checkUserLevelUpgrade 检查用户是否需要升级（使用预查询的等级配置）
// 返回用户应该达到的等级ID
func (e *OrdGiftcardWriteoffs) checkUserLevelUpgrade(currentLevelId int, totalExperience int, levels []models.HsConfgiUserLevels) int {
	if len(levels) == 0 {
		return currentLevelId
	}

	// 找到用户应该达到的最高等级
	// 遍历等级列表，找到 totalExperience >= up_experience 的最高等级
	targetLevelId := currentLevelId
	for _, level := range levels {
		upExperience, _ := strconv.Atoi(level.UpExperience)
		if totalExperience >= upExperience {
			targetLevelId = level.Id
		} else {
			// 经验不足以升到这个等级，后续等级也不用检查了
			break
		}
	}

	return targetLevelId
}

// ============ 计算用户入账金额（辅助接口） ============

// CalculateUserLocalCurrency 计算用户入账金额（辅助接口）
func (e *OrdGiftcardWriteoffs) CalculateUserLocalCurrency(c *dto.OrdGiftcardWriteoffsCalculateReq) (*dto.OrdGiftcardWriteoffsCalculateResp, error) {
	// 1. 查询订单信息
	var order models.OrdUserOrders
	err := e.Orm.Where("id = ?", c.OrderId).First(&order).Error
	if err != nil {
		return nil, errors.New("获取订单信息失败")
	}

	userId := order.UserId

	// 2. 查询用户货币代码
	var user models.HsUsers
	err = e.Orm.Select("currency_code").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, errors.New("获取用户信息失败")
	}

	userCurrencyCode := user.CurrencyCode
	if userCurrencyCode == "" {
		return nil, errors.New("用户未设置货币代码")
	}

	// 3. 获取礼品卡配置
	if c.GiftCardId <= 0 {
		return nil, errors.New("必须提供礼品卡ID")
	}

	var giftcard models.OrdGiftcard
	err = e.Orm.Where("id = ?", c.GiftCardId).First(&giftcard).Error
	if err != nil {
		return nil, errors.New("查询礼品卡配置失败")
	}

	// 确定折扣率
	usedDiscountRate := c.DiscountRate
	if usedDiscountRate == "" {
		usedDiscountRate = giftcard.DiscountRate
	}

	// 查询区域货币
	if giftcard.RegionId == "" || giftcard.RegionId == "0" {
		return nil, errors.New("礼品卡未配置区域")
	}

	var region models.OrdGiftcardRegion
	err = e.Orm.Where("id = ?", giftcard.RegionId).First(&region).Error
	if err != nil {
		return nil, errors.New("查询礼品卡区域失败")
	}

	if region.CurrencyCode == "" {
		return nil, errors.New("礼品卡区域未配置货币代码")
	}

	sourceCurrencyCode := region.CurrencyCode

	// 4. 计算配置汇率
	isCrypto := order.BalanceType == 2
	var configRate, targetCurrency string

	if isCrypto {
		targetCurrency = "USD"
		if sourceCurrencyCode == "USD" {
			configRate = "1.00000000"
		} else {
			configRate, err = e.getCurrencyRate(e.Orm, sourceCurrencyCode, "USD")
			if err != nil {
				return nil, err
			}
		}
	} else {
		targetCurrency = userCurrencyCode
		if sourceCurrencyCode == userCurrencyCode {
			configRate = "1.00000000"
		} else {
			configRate, err = e.getCurrencyRateViaUSD(e.Orm, sourceCurrencyCode, userCurrencyCode)
			if err != nil {
				return nil, err
			}
		}
	}

	// 5. 计算用户本地货币金额
	cardValueFloat, _ := strconv.ParseFloat(c.RecognizedCardValue, 64)
	discountRateFloat := 1.0
	if usedDiscountRate != "" && usedDiscountRate != "0" {
		if parsed, err := strconv.ParseFloat(usedDiscountRate, 64); err == nil && parsed > 0 {
			discountRateFloat = parsed
		}
	}

	configRateFloat, _ := strconv.ParseFloat(configRate, 64)
	userLocalAmount := math.Floor(cardValueFloat * discountRateFloat * configRateFloat)

	// 6. 面额校验
	var denominationValidation *dto.OrdGiftcardWriteoffsDenominationValidation
	if giftcard.ValuesConfig != "" {
		validationErr := e.validateDenomination(userLocalAmount, giftcard.ValuesConfig)
		denominationValidation = &dto.OrdGiftcardWriteoffsDenominationValidation{
			IsValid: validationErr == nil,
		}
		if validationErr != nil {
			denominationValidation.ErrorMessage = validationErr.Error()
		}

		var config ValuesConfigStruct
		if json.Unmarshal([]byte(giftcard.ValuesConfig), &config) == nil {
			if len(config.Fixed) > 0 {
				denominationValidation.AllowedFixed = make([]string, len(config.Fixed))
				for i, v := range config.Fixed {
					denominationValidation.AllowedFixed[i] = fmt.Sprintf("%.2f", v)
				}
			}
			if config.Range != nil {
				denominationValidation.AllowedRange = &struct {
					Min string `json:"min" comment:"最小值"`
					Max string `json:"max" comment:"最大值"`
				}{
					Min: fmt.Sprintf("%.2f", config.Range.Min),
					Max: fmt.Sprintf("%.2f", config.Range.Max),
				}
			}
		}
	}

	return &dto.OrdGiftcardWriteoffsCalculateResp{
		UserLocalCurrencyAmount: fmt.Sprintf("%.8f", userLocalAmount),
		UserCurrencyCode:        targetCurrency,
		ConfigRate:              configRate,
		DiscountRate:            usedDiscountRate,
		IsCrypto:                isCrypto,
		OrderCurrencyCode:       sourceCurrencyCode,
		DenominationValidation:  denominationValidation,
	}, nil
}
