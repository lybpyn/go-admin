package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type OrdGiftcardWriteoffs struct {
	service.Service
}

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

// BatchInsert 批量创建核销记录（带事务处理：余额增加、流水记录、账户分成）
func (e *OrdGiftcardWriteoffs) BatchInsert(c *dto.OrdGiftcardWriteoffsBatchInsertReq) error {
	var err error

	// ============ 事务外：所有只读查询和数据准备 ============

	// 1. 查询订单信息
	var order models.OrdUserOrders
	err = e.Orm.Where("id = ?", c.OrderId).First(&order).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert get order error:%s", err)
		return errors.New("获取订单信息失败")
	}

	// 预检查订单状态，必须是已经接单状态（status=1）才能核销（快速失败）
	if order.Status != 1 {
		e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert order status invalid: orderId=%d, status=%d", order.Id, order.Status)
		return errors.New("订单状态不正确，只有已接单的订单才能核销")
	}

	// 从订单中获取用户ID
	userId := order.UserId

	// 2. 查询用户货币代码（直接从用户表读取）
	var user models.HsUsers
	err = e.Orm.Select("currency_code").Where("id = ?", userId).First(&user).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert get user error:%s", err)
		return errors.New("获取用户信息失败")
	}

	userCurrencyCode := user.CurrencyCode
	if userCurrencyCode == "" {
		e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert user currency code is empty for user %d", userId)
		return errors.New("用户未设置货币代码")
	}

	// 3. 计算配置汇率和目标货币
	isCrypto := order.BalanceType == 2
	var configRate string
	var targetCurrency string

	if isCrypto {
		targetCurrency = "USD"
		if order.CurrencyCode == "USD" {
			configRate = "1.00000000"
		} else {
			rate, err := e.getCurrencyRate(e.Orm, order.CurrencyCode, "USD")
			if err != nil {
				return err
			}
			configRate = rate
		}
	} else {
		targetCurrency = userCurrencyCode
		if order.CurrencyCode == userCurrencyCode {
			configRate = "1.00000000"
		} else {
			rate, err := e.getCurrencyRateViaUSD(e.Orm, order.CurrencyCode, userCurrencyCode)
			if err != nil {
				return err
			}
			configRate = rate
		}
	}

	// 4. 批量查询礼品卡配置
	giftcardIdSet := make(map[int]bool)
	for _, item := range c.WriteoffList {
		if item.GiftCardId > 0 {
			giftcardIdSet[item.GiftCardId] = true
		}
	}

	// 批量查询礼品卡配置（获取面额规则和折扣率）
	giftcardCache := make(map[int]*models.OrdGiftcard)
	regionIdSet := make(map[string]bool)
	if len(giftcardIdSet) > 0 {
		giftcardIds := make([]int, 0, len(giftcardIdSet))
		for id := range giftcardIdSet {
			giftcardIds = append(giftcardIds, id)
		}

		var giftcards []models.OrdGiftcard
		err = e.Orm.Where("id IN ?", giftcardIds).Find(&giftcards).Error
		if err != nil {
			e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert get giftcard configs error:%s", err)
			return errors.New("查询礼品卡配置失败")
		}

		for i := range giftcards {
			giftcardCache[giftcards[i].Id] = &giftcards[i]
			// 收集区域ID用于查询货币代码
			if giftcards[i].RegionId != "" && giftcards[i].RegionId != "0" {
				regionIdSet[giftcards[i].RegionId] = true
			}
		}
	}

	// 5. 批量查询礼品卡区域配置（获取货币代码）
	regionCache := make(map[string]*models.OrdGiftcardRegion)
	if len(regionIdSet) > 0 {
		regionIds := make([]string, 0, len(regionIdSet))
		for id := range regionIdSet {
			regionIds = append(regionIds, id)
		}

		var regions []models.OrdGiftcardRegion
		err = e.Orm.Where("id IN ?", regionIds).Find(&regions).Error
		if err != nil {
			e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert get region configs error:%s", err)
			return errors.New("查询礼品卡区域配置失败")
		}

		for i := range regions {
			regionCache[strconv.Itoa(regions[i].Id)] = &regions[i]
		}
	}

	// 6. 预转换配置汇率
	configRateFloat, _ := strconv.ParseFloat(configRate, 64)

	// 7. 验证核销状态并构建批量插入数据
	if len(c.WriteoffList) == 0 {
		return errors.New("核销记录列表不能为空")
	}

	// 获取第一条记录的状态作为订单最终状态
	finalStatus := c.WriteoffList[0].Status

	// 核销状态只能是成功(1)或失败(2)
	if finalStatus != 1 && finalStatus != 2 {
		e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert invalid writeoff status: %d", finalStatus)
		return errors.New("核销状态无效，只能是已核销(1)或失败(2)")
	}

	writeoffRecords := make([]models.OrdGiftcardWriteoffs, 0, len(c.WriteoffList))
	var totalAmount float64

	for _, item := range c.WriteoffList {
		// 使用int类型的GiftCardId
		giftcardId := item.GiftCardId

		// 从数据库配置中获取折扣率（平台售卡汇率）
		var platformSaleRate string
		var discountRateFloat float64
		if giftcardId > 0 {
			if giftcard, exists := giftcardCache[giftcardId]; exists {
				platformSaleRate = giftcard.DiscountRate
				discountRateFloat, _ = strconv.ParseFloat(platformSaleRate, 64)
			} else {
				e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert giftcard config not found: giftcardId=%d", giftcardId)
				return fmt.Errorf("礼品卡配置不存在：ID=%d", giftcardId)
			}
		} else {
			// 如果没有提供礼品卡ID，使用默认值
			platformSaleRate = ""
			discountRateFloat = 0
		}

		// 获取卡片面值
		var cardValueFloat float64
		if item.RecognizedCardValue != "" {
			cardValueFloat, _ = strconv.ParseFloat(item.RecognizedCardValue, 64)
		}

		// 获取用户本地货币金额
		var userLocalAmount float64
		var convertedAmount string

		if item.UserLocalCurrencyAmount != "" && item.UserLocalCurrencyAmount != "0" {
			// 使用前端传入的用户本地货币金额
			userLocalAmount, _ = strconv.ParseFloat(item.UserLocalCurrencyAmount, 64)
			convertedAmount = item.UserLocalCurrencyAmount

			// 校验面额规则（如果有对应的礼品卡配置）
			if giftcardId > 0 {
				if giftcard, exists := giftcardCache[giftcardId]; exists && giftcard.ValuesConfig != "" {
					// 解析面额配置并校验
					if err := e.validateDenomination(userLocalAmount, giftcard.ValuesConfig); err != nil {
						e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert denomination validation failed: amount=%.8f, config=%s, error=%s",
							userLocalAmount, giftcard.ValuesConfig, err.Error())
						return fmt.Errorf("金额 %.2f 不符合入账面额规则：%s", userLocalAmount, err.Error())
					}
				}
			}
		} else {
			// 计算转换后的金额（给用户的金额）
			// 用户金额 = 卡片面值 × 货币转换汇率
			convertedAmount = "0.00000000"
			if cardValueFloat > 0 {
				userLocalAmount = cardValueFloat * configRateFloat
				convertedAmount = fmt.Sprintf("%.8f", userLocalAmount)
			}
		}

		// 只累加已核销状态的金额
		if item.Status == 1 && userLocalAmount > 0 {
			totalAmount += userLocalAmount
		}

		// 计算平台入账金额（如果未提供）
		// 平台入账金额 = 卡片面值 × 折扣率
		platformSettlementAmount := item.PlatformSettlementAmount

		// 根据礼品卡区域配置获取货币代码（以管理员核销选择为准）
		platformSettlementCurrency := order.CurrencyCode // 默认使用订单货币
		var platformToUsdRate string

		if giftcardId > 0 {
			if giftcard, exists := giftcardCache[giftcardId]; exists && giftcard.RegionId != "" {
				if region, regionExists := regionCache[giftcard.RegionId]; regionExists {
					platformSettlementCurrency = region.CurrencyCode
				}
			}
		}

		// 计算平台货币对美元汇率
		if platformSettlementCurrency == "USD" {
			platformToUsdRate = "1.00000000"
		} else {
			rate, err := e.getCurrencyRate(e.Orm, platformSettlementCurrency, "USD")
			if err != nil {
				e.Log.Warnf("Get platform currency %s to USD rate error:%s, use 1", platformSettlementCurrency, err)
				platformToUsdRate = "1.00000000"
			} else {
				platformToUsdRate = rate
			}
		}

		if platformSettlementAmount == "" || platformSettlementAmount == "0" {
			// 自动计算平台入账金额
			if cardValueFloat > 0 && discountRateFloat > 0 {
				// 平台收卡成本 = 卡片面值 × 折扣率
				platformCost := cardValueFloat * discountRateFloat
				platformSettlementAmount = fmt.Sprintf("%.8f", platformCost)
			}
		}

		record := models.OrdGiftcardWriteoffs{
			UserId:                     userId,
			OrderId:                    c.OrderId,
			GiftCardId:                 giftcardId,
			GiftCardDiscountId:         0, // 不再使用折扣表，设置为0
			Status:                     item.Status,
			Remark:                     item.Remark,
			AdminRecognizedCode:        item.AdminRecognizedCode,
			PlatformSaleRate:           platformSaleRate,
			RecognizedCardValue:        item.RecognizedCardValue,
			FailureImageUrl:            item.FailureImageUrl,
			SupplierId:                 strconv.Itoa(item.SupplierId),
			ConfigRate:                 configRate,
			UserLocalCurrencyAmount:    convertedAmount,
			UserCurrencyCode:           targetCurrency,
			PlatformSettlementAmount:   platformSettlementAmount,
			PlatformSettlementCurrency: platformSettlementCurrency,
			PlatformToUsdRate:          platformToUsdRate,
		}
		record.CreateBy = c.CreateBy
		writeoffRecords = append(writeoffRecords, record)
	}

	// 8. 准备订单状态更新数据
	// 核销成功(1) → 订单完成(2)，处理状态完成(3)
	// 核销失败(2) → 订单驳回(4)，处理状态取消(2)
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

	// ============ 事务内：只执行写操作 ============
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		// 1. 加锁查询订单，再次验证状态（防止并发核销）
		var orderLocked models.OrdUserOrders
		err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("id = ?", c.OrderId).
			First(&orderLocked).Error
		if err != nil {
			e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert lock order error:%s", err)
			return errors.New("获取订单锁失败")
		}

		// 再次检查状态（防止并发修改）
		if orderLocked.Status != 1 {
			e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert order status changed: orderId=%d, status=%d", orderLocked.Id, orderLocked.Status)
			return errors.New("订单状态已变更，无法核销")
		}

		// 2. 批量插入核销记录
		if err = tx.Create(&writeoffRecords).Error; err != nil {
			e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert error:%s", err)
			return err
		}

		// 3. 更新订单状态
		e.Log.Infof("Updating order %d: status=%d, processing_status=%d (writeoff status: %d)", c.OrderId, orderStatus, processingStatus, finalStatus)

		if err = tx.Model(&models.OrdUserOrders{}).Where("id = ?", c.OrderId).Updates(orderUpdates).Error; err != nil {
			e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert update order status error:%s", err)
			return err
		}

		// 4. 如果核销成功且总金额大于0，需要进行余额增加、流水记录和账户分成
		if finalStatus == 1 && totalAmount > 0 {
			// 用户入账（独立封装方法）
			bizType := "giftcard_writeoff_fiat"
			if isCrypto {
				bizType = "giftcard_writeoff_crypto"
			}

			remark := fmt.Sprintf("礼品卡核销到账，订单号: %d", c.OrderId)
			err = e.creditUserBalance(tx, userId, totalAmount, targetCurrency, isCrypto, c.OrderId, bizType, remark, c.CreateBy)
			if err != nil {
				e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert credit user balance error:%s", err)
				return err
			}

			// 处理邀请分成（分成进入冻结余额）
			err = e.processInviteCommissions(tx, userId, c.OrderId, totalAmount, targetCurrency, isCrypto, c.CreateBy)
			if err != nil {
				e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert process commissions error:%s", err)
				return err
			}
		}

		return nil
	})
}

// creditUserBalance 用户入账（独立封装）
// 参数：
//   - userId: 用户ID
//   - amount: 入账金额
//   - currencyCode: 货币代码
//   - isCrypto: 是否虚拟币（true=虚拟币，false=法币）
//   - orderId: 订单ID
//   - bizType: 业务类型（如：giftcard_writeoff_fiat、giftcard_writeoff_crypto）
//   - remark: 备注信息
//   - createBy: 创建人ID
func (e *OrdGiftcardWriteoffs) creditUserBalance(tx *gorm.DB, userId int, amount float64, currencyCode string, isCrypto bool, orderId int, bizType string, remark string, createBy int) error {
	// 1. 查询用户当前余额和版本号
	var user models.HsUsers
	err := tx.Select("id, balance, crypto_balance, version").Where("id = ?", userId).First(&user).Error
	if err != nil {
		e.Log.Errorf("creditUserBalance get user balance error:%s", err)
		return errors.New("获取用户余额失败")
	}

	// 2. 根据类型选择余额字段和精度
	var balanceBefore float64
	var balanceField string
	var decimalPlaces int

	if isCrypto {
		balanceField = "crypto_balance"
		decimalPlaces = 8
		balanceBefore, _ = strconv.ParseFloat(user.CryptoBalance, 64)
	} else {
		balanceField = "balance"
		decimalPlaces = 2
		balanceBefore, _ = strconv.ParseFloat(user.Balance, 64)
	}

	balanceAfter := balanceBefore + amount

	// 3. 更新用户余额（使用乐观锁）
	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", userId, user.Version).
		Updates(map[string]interface{}{
			balanceField: fmt.Sprintf("%.*f", decimalPlaces, balanceAfter),
			"version":    gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		e.Log.Errorf("creditUserBalance update balance error:%s", result.Error)
		return errors.New("更新用户余额失败")
	}

	if result.RowsAffected == 0 {
		e.Log.Errorf("creditUserBalance update balance conflict for user %d", userId)
		return errors.New("余额更新冲突，请重试")
	}

	// 4. 创建流水记录
	orderIdStr := strconv.Itoa(orderId)
	userIdStr := strconv.Itoa(userId)

	ledger := models.HsUserLedger{
		UserId:         userIdStr,
		CurrencyCode:   currencyCode,
		Direction:      "1", // 1=入账
		Amount:         fmt.Sprintf("%.*f", decimalPlaces, amount),
		BalanceBefore:  fmt.Sprintf("%.*f", decimalPlaces, balanceBefore),
		BalanceAfter:   fmt.Sprintf("%.*f", decimalPlaces, balanceAfter),
		BizType:        bizType,
		BizId:          orderIdStr,
		IdempotencyKey: fmt.Sprintf("%s:%d:%d", bizType, orderId, time.Now().UnixNano()),
		RefTable:       "ord_giftcard_writeoffs",
		RefId:          orderIdStr,
		Remark:         remark,
		Status:         "1", // 1=已入账
	}
	ledger.CreateBy = createBy

	err = tx.Create(&ledger).Error
	if err != nil {
		e.Log.Errorf("creditUserBalance create ledger error:%s", err)
		return errors.New("创建流水记录失败")
	}

	e.Log.Infof("User %d credited %.*f %s, balance: %.*f -> %.*f",
		userId, decimalPlaces, amount, currencyCode,
		decimalPlaces, balanceBefore, decimalPlaces, balanceAfter)

	return nil
}

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
		e.Log.Errorf("Get currency rate %s to %s error:%s \r\n", fromCurrency, toCurrency, err)
		return "", fmt.Errorf("未找到 %s 到 %s 的有效汇率配置", fromCurrency, toCurrency)
	}
	return rateRecord.Rate, nil
}

// getCurrencyRateViaUSD 直接获取货币汇率，不使用USD中转
func (e *OrdGiftcardWriteoffs) getCurrencyRateViaUSD(tx *gorm.DB, fromCurrency, toCurrency string) (string, error) {
	if fromCurrency == toCurrency {
		return "1.00000000", nil
	}

	// 直接查询汇率配置
	return e.getCurrencyRate(tx, fromCurrency, toCurrency)
}

// processInviteCommissions 处理邀请分成（一级5%、二级3%），分成进入冻结余额
func (e *OrdGiftcardWriteoffs) processInviteCommissions(tx *gorm.DB, userId, orderId int, amount float64, sourceCurrency string, isCrypto bool, createBy int) error {
	orderIdStr := strconv.Itoa(orderId)

	// 1. 查询邀请关系
	var inviteRelation models.HsInviteRelations
	err := tx.Where("user_id = ?", userId).First(&inviteRelation).Error
	if err != nil {
		// 如果没有邀请关系，不报错，直接返回
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Infof("User %s has no invite relations, skip commission", userId)
			return nil
		}
		return err
	}

	// 2. 处理一级邀请人分成（根据邀请人等级获取分成配置）
	if inviteRelation.Level1InviterId != "" && inviteRelation.Level1InviterId != "0" {
		// 获取一级邀请人的等级和分成配置
		config, err := e.getInviterCommissionConfig(tx, inviteRelation.Level1InviterId, 1)
		if err != nil {
			e.Log.Errorf("Get level1 inviter commission config error:%s \r\n", err)
			return err
		}

		// 如果分成比例大于0，才进行分成处理
		if config.CommissionRate > 0 {
			commissionAmount := amount * config.CommissionRate

			// 创建分成记录
			commission := models.HsInviteCommissions{
				OrderId:          orderIdStr,
				UserId:           inviteRelation.Level1InviterId,
				CommissionLevel:  "1",
				CommissionRate:   fmt.Sprintf("%.2f", config.CommissionRate),
				CommissionAmount: fmt.Sprintf("%.8f", commissionAmount),
				Status:           "1", // 1=已结算
			}
			commission.CreateBy = createBy

			err = tx.Create(&commission).Error
			if err != nil {
				e.Log.Errorf("Create level1 commission error:%s \r\n", err)
				return errors.New("创建一级分成记录失败")
			}

			// 根据frozen_rate分配到冻结余额和可用余额
			frozenAmount := commissionAmount * config.FrozenRate
			availableAmount := commissionAmount * (1 - config.FrozenRate)

			// 更新一级邀请人冻结余额
			if frozenAmount > 0 {
				err = e.updateInviterFrozenBalance(tx, inviteRelation.Level1InviterId, frozenAmount, sourceCurrency, isCrypto, orderId, "1", createBy)
				if err != nil {
					return err
				}
			}

			// 更新一级邀请人可用余额
			if availableAmount > 0 {
				err = e.updateInviterAvailableBalance(tx, inviteRelation.Level1InviterId, availableAmount, sourceCurrency, isCrypto, orderId, "1", createBy)
				if err != nil {
					return err
				}
			}
		}
	}

	// 3. 处理二级邀请人分成（根据邀请人等级获取分成配置）
	if inviteRelation.Level2InviterId != "" && inviteRelation.Level2InviterId != "0" {
		// 获取二级邀请人的等级和分成配置
		config, err := e.getInviterCommissionConfig(tx, inviteRelation.Level2InviterId, 2)
		if err != nil {
			e.Log.Errorf("Get level2 inviter commission config error:%s \r\n", err)
			return err
		}

		// 如果分成比例大于0，才进行分成处理
		if config.CommissionRate > 0 {
			commissionAmount := amount * config.CommissionRate

			// 创建分成记录
			commission := models.HsInviteCommissions{
				OrderId:          orderIdStr,
				UserId:           inviteRelation.Level2InviterId,
				CommissionLevel:  "2",
				CommissionRate:   fmt.Sprintf("%.2f", config.CommissionRate),
				CommissionAmount: fmt.Sprintf("%.8f", commissionAmount),
				Status:           "1", // 1=已结算
			}
			commission.CreateBy = createBy

			err = tx.Create(&commission).Error
			if err != nil {
				e.Log.Errorf("Create level2 commission error:%s \r\n", err)
				return errors.New("创建二级分成记录失败")
			}

			// 根据frozen_rate分配到冻结余额和可用余额
			frozenAmount := commissionAmount * config.FrozenRate
			availableAmount := commissionAmount * (1 - config.FrozenRate)

			// 更新二级邀请人冻结余额
			if frozenAmount > 0 {
				err = e.updateInviterFrozenBalance(tx, inviteRelation.Level2InviterId, frozenAmount, sourceCurrency, isCrypto, orderId, "2", createBy)
				if err != nil {
					return err
				}
			}

			// 更新二级邀请人可用余额
			if availableAmount > 0 {
				err = e.updateInviterAvailableBalance(tx, inviteRelation.Level2InviterId, availableAmount, sourceCurrency, isCrypto, orderId, "2", createBy)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// InviterCommissionConfig 邀请人分成配置
type InviterCommissionConfig struct {
	CommissionRate float64 // 分成比例
	FrozenRate     float64 // 进入冻结余额的比例
}

// getInviterCommissionConfig 获取邀请人的分成配置（从hs_config_invite_commission表）
func (e *OrdGiftcardWriteoffs) getInviterCommissionConfig(tx *gorm.DB, inviterId string, level int) (*InviterCommissionConfig, error) {
	// 1. 获取名为 'invite_config' 的启用配置
	var commissionConfig models.HsConfigInviteCommission
	err := tx.Where("config_name = ? AND status = 1", "invite_config").
		First(&commissionConfig).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("Commission config 'invite_config' not found or not active")
			return nil, errors.New("邀请分成配置不存在或未启用，请联系管理员配置")
		}
		e.Log.Errorf("Get commission config error:%s \r\n", err)
		return nil, fmt.Errorf("获取分成配置失败: %s", err.Error())
	}

	// 2. 根据level参数获取对应的分成比例
	var commissionRateStr string
	if level == 1 {
		commissionRateStr = commissionConfig.FirstLevelRate
	} else if level == 2 {
		commissionRateStr = commissionConfig.SecondLevelRate
	} else {
		return nil, fmt.Errorf("无效的分成层级: %d", level)
	}

	// 3. 验证分成比例不能为空
	if commissionRateStr == "" {
		e.Log.Errorf("Commission rate is empty for level %d in config 'invite_config'", level)
		return nil, fmt.Errorf("分成比例配置为空（%d级），请联系管理员配置", level)
	}

	// 4. 转换为float64
	commissionRate, err := strconv.ParseFloat(commissionRateStr, 64)
	if err != nil {
		e.Log.Errorf("Parse commission rate %s error:%s \r\n", commissionRateStr, err)
		return nil, fmt.Errorf("分成比例格式错误: %s", commissionRateStr)
	}

	// 5. 转换为小数（数据库存储的是百分比，如5.0000表示5%）
	commissionRate = commissionRate / 100.0

	// 6. 获取邀请人信息以获取等级特权
	var inviter models.HsUsers
	err = tx.Where("id = ?", inviterId).First(&inviter).Error
	if err != nil {
		e.Log.Errorf("Get inviter %s error:%s \r\n", inviterId, err)
		return nil, fmt.Errorf("获取邀请人信息失败")
	}

	// 7. 获取邀请人的等级配置，使用rebate_rate作为返利比例
	var frozenRate float64 = 1.0 // 默认100%进入冻结余额
	if inviter.LevelId != "" && inviter.LevelId != "0" {
		var userLevel models.HsConfgiUserLevels
		err = tx.Where("id = ? AND is_active = 1", inviter.LevelId).First(&userLevel).Error
		if err == nil && userLevel.RebateRate != "" {
			// 直接使用rebate_rate字段
			frozenRate, _ = strconv.ParseFloat(userLevel.RebateRate, 64)
		}
	}

	return &InviterCommissionConfig{
		CommissionRate: commissionRate,
		FrozenRate:     frozenRate,
	}, nil
}

// updateInviterFrozenBalance 更新邀请人冻结余额和流水（支持多货币汇率转换）
func (e *OrdGiftcardWriteoffs) updateInviterFrozenBalance(tx *gorm.DB, inviterId string, sourceAmount float64, sourceCurrency string, isCrypto bool, orderId int, level string, createBy int) error {
	orderIdStr := strconv.Itoa(orderId)

	// 1. 获取邀请人信息（直接从用户表读取货币代码）
	var inviter models.HsUsers
	err := tx.Where("id = ?", inviterId).First(&inviter).Error
	if err != nil {
		e.Log.Errorf("Get inviter %s error:%s \r\n", inviterId, err)
		return fmt.Errorf("获取邀请人信息失败")
	}

	// 2. 使用订单用户的货币，不进行汇率转换
	convertedAmount := sourceAmount
	targetCurrency := sourceCurrency

	// 3. 检查冻结余额限制（根据CurrencyType分别处理）
	var currencyType string
	if isCrypto {
		currencyType = "crypto"
	} else {
		currencyType = "fiat"
	}

	var frozenLimitConfig models.HsConfigFrozenLimit
	err = tx.Where("currency_type = ? AND currency_code = ? AND is_active = 1", currencyType, targetCurrency).
		First(&frozenLimitConfig).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		e.Log.Errorf("Get frozen limit config error:%s \r\n", err)
		return fmt.Errorf("获取冻结余额限制配置失败")
	}

	// 4. 准备余额字段和精度
	var frozenBalanceBefore float64
	var frozenBalanceField string
	var availableBalanceBefore float64
	var availableBalanceField string
	var decimalPlaces int

	if isCrypto {
		// 虚拟币冻结余额（8位小数）
		frozenBalanceField = "crypto_frozen_balance"
		availableBalanceField = "crypto_balance"
		decimalPlaces = 8
		frozenBalanceBefore, _ = strconv.ParseFloat(inviter.CryptoFrozenBalance, 64)
		availableBalanceBefore, _ = strconv.ParseFloat(inviter.CryptoBalance, 64)
	} else {
		// 法币冻结余额（2位小数）
		frozenBalanceField = "frozen_balance"
		availableBalanceField = "balance"
		decimalPlaces = 2
		frozenBalanceBefore, _ = strconv.ParseFloat(inviter.FrozenBalance, 64)
		availableBalanceBefore, _ = strconv.ParseFloat(inviter.Balance, 64)
	}

	// 5. 检查是否超过限制，计算冻结部分和可用部分
	var frozenAmount float64
	var availableAmount float64

	if frozenLimitConfig.Id > 0 {
		frozenLimit, _ := strconv.ParseFloat(frozenLimitConfig.FrozenLimitAmount, 64)
		potentialFrozenBalance := frozenBalanceBefore + convertedAmount

		if potentialFrozenBalance > frozenLimit {
			// 超过限制，分配到冻结和可用两部分
			frozenAmount = frozenLimit - frozenBalanceBefore
			if frozenAmount < 0 {
				frozenAmount = 0
			}
			availableAmount = convertedAmount - frozenAmount

			e.Log.Infof("Inviter %s frozen balance would exceed limit: current=%.8f, add=%.8f, limit=%.8f, currency=%s, type=%s. Split: frozen=%.8f, available=%.8f",
				inviterId, frozenBalanceBefore, convertedAmount, frozenLimit, targetCurrency, currencyType, frozenAmount, availableAmount)
		} else {
			// 未超过限制，全部增加到冻结余额
			frozenAmount = convertedAmount
			availableAmount = 0
		}
	} else {
		// 没有限制配置，全部增加到冻结余额
		frozenAmount = convertedAmount
		availableAmount = 0
	}

	// 6. 更新邀请人余额（使用乐观锁）
	frozenBalanceAfter := frozenBalanceBefore + frozenAmount
	availableBalanceAfter := availableBalanceBefore + availableAmount

	updateFields := map[string]interface{}{
		"version": gorm.Expr("version + 1"),
	}

	if frozenAmount > 0 {
		updateFields[frozenBalanceField] = fmt.Sprintf("%.*f", decimalPlaces, frozenBalanceAfter)
	}

	if availableAmount > 0 {
		updateFields[availableBalanceField] = fmt.Sprintf("%.*f", decimalPlaces, availableBalanceAfter)
	}

	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", inviterId, inviter.Version).
		Updates(updateFields)

	if result.Error != nil {
		e.Log.Errorf("Update inviter %s balance error:%s \r\n", inviterId, result.Error)
		return errors.New("更新邀请人余额失败")
	}

	if result.RowsAffected == 0 {
		e.Log.Errorf("Update inviter %s balance conflict", inviterId)
		return errors.New("邀请人余额更新冲突，请重试")
	}

	// 7. 创建邀请人流水记录
	nanoTime := time.Now().UnixNano()

	// 7.1 如果有冻结部分，创建冻结流水
	if frozenAmount > 0 {
		bizType := "invite_commission_frozen_fiat"
		if isCrypto {
			bizType = "invite_commission_frozen_crypto"
		}

		ledger := models.HsUserLedger{
			UserId:         inviterId,
			CurrencyCode:   targetCurrency,
			Direction:      "1", // 1=入账
			Amount:         fmt.Sprintf("%.*f", decimalPlaces, frozenAmount),
			BalanceBefore:  fmt.Sprintf("%.*f", decimalPlaces, frozenBalanceBefore),
			BalanceAfter:   fmt.Sprintf("%.*f", decimalPlaces, frozenBalanceAfter),
			BizType:        bizType,
			BizId:          orderIdStr,
			IdempotencyKey: fmt.Sprintf("INVITE_COMMISSION_FROZEN_L%s:%d:%d", level, orderId, nanoTime),
			RefTable:       "hs_invite_commissions",
			RefId:          orderIdStr,
			Remark:         fmt.Sprintf("邀请分成冻结（%s级），订单号: %d", level, orderId),
			Status:         "1", // 1=已入账
		}
		ledger.CreateBy = createBy

		err = tx.Create(&ledger).Error
		if err != nil {
			e.Log.Errorf("Create inviter %s frozen ledger error:%s \r\n", inviterId, err)
			return errors.New("创建邀请人冻结流水记录失败")
		}
	}

	// 7.2 如果有可用部分，创建可用流水
	if availableAmount > 0 {
		bizType := "invite_commission_available_fiat"
		if isCrypto {
			bizType = "invite_commission_available_crypto"
		}

		ledger := models.HsUserLedger{
			UserId:         inviterId,
			CurrencyCode:   targetCurrency,
			Direction:      "1", // 1=入账
			Amount:         fmt.Sprintf("%.*f", decimalPlaces, availableAmount),
			BalanceBefore:  fmt.Sprintf("%.*f", decimalPlaces, availableBalanceBefore),
			BalanceAfter:   fmt.Sprintf("%.*f", decimalPlaces, availableBalanceAfter),
			BizType:        bizType,
			BizId:          orderIdStr,
			IdempotencyKey: fmt.Sprintf("INVITE_COMMISSION_AVAILABLE_L%s:%d:%d", level, orderId, nanoTime+1),
			RefTable:       "hs_invite_commissions",
			RefId:          orderIdStr,
			Remark:         fmt.Sprintf("邀请分成可用（%s级，超冻结限额），订单号: %d", level, orderId),
			Status:         "1", // 1=已入账
		}
		ledger.CreateBy = createBy

		err = tx.Create(&ledger).Error
		if err != nil {
			e.Log.Errorf("Create inviter %s available ledger error:%s \r\n", inviterId, err)
			return errors.New("创建邀请人可用流水记录失败")
		}
	}

	return nil
}

// updateInviterAvailableBalance 更新邀请人可用余额和流水（使用订单用户的货币）
func (e *OrdGiftcardWriteoffs) updateInviterAvailableBalance(tx *gorm.DB, inviterId string, sourceAmount float64, sourceCurrency string, isCrypto bool, orderId int, level string, createBy int) error {
	orderIdStr := strconv.Itoa(orderId)

	// 1. 获取邀请人信息
	var inviter models.HsUsers
	err := tx.Where("id = ?", inviterId).First(&inviter).Error
	if err != nil {
		e.Log.Errorf("Get inviter %s error:%s \r\n", inviterId, err)
		return fmt.Errorf("获取邀请人信息失败")
	}

	// 2. 使用订单用户的货币，不进行汇率转换
	convertedAmount := sourceAmount
	targetCurrency := sourceCurrency

	// 3. 更新邀请人可用余额（使用乐观锁）
	var balanceBefore float64
	var balanceAfter float64
	var balanceField string
	var decimalPlaces int

	if isCrypto {
		// 虚拟币可用余额（8位小数）
		balanceField = "crypto_balance"
		decimalPlaces = 8
		balanceBefore, _ = strconv.ParseFloat(inviter.CryptoBalance, 64)
	} else {
		// 法币可用余额（2位小数）
		balanceField = "balance"
		decimalPlaces = 2
		balanceBefore, _ = strconv.ParseFloat(inviter.Balance, 64)
	}

	balanceAfter = balanceBefore + convertedAmount

	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", inviterId, inviter.Version).
		Updates(map[string]interface{}{
			balanceField: fmt.Sprintf("%.*f", decimalPlaces, balanceAfter),
			"version":    gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		e.Log.Errorf("Update inviter %s available balance error:%s \r\n", inviterId, result.Error)
		return errors.New("更新邀请人可用余额失败")
	}

	if result.RowsAffected == 0 {
		e.Log.Errorf("Update inviter %s available balance conflict", inviterId)
		return errors.New("邀请人可用余额更新冲突，请重试")
	}

	// 5. 创建邀请人流水记录
	bizType := "invite_commission_available_fiat"
	if isCrypto {
		bizType = "invite_commission_available_crypto"
	}

	ledger := models.HsUserLedger{
		UserId:         inviterId,
		CurrencyCode:   targetCurrency,
		Direction:      "1", // 1=入账
		Amount:         fmt.Sprintf("%.*f", decimalPlaces, convertedAmount),
		BalanceBefore:  fmt.Sprintf("%.*f", decimalPlaces, balanceBefore),
		BalanceAfter:   fmt.Sprintf("%.*f", decimalPlaces, balanceAfter),
		BizType:        bizType,
		BizId:          orderIdStr,
		IdempotencyKey: fmt.Sprintf("INVITE_COMMISSION_AVAILABLE_L%s:%d:%d", level, orderId, time.Now().UnixNano()),
		RefTable:       "hs_invite_commissions",
		RefId:          orderIdStr,
		Remark:         fmt.Sprintf("邀请分成可用（%s级），订单号: %d", level, orderId),
		Status:         "1", // 1=已入账
	}
	ledger.CreateBy = createBy

	err = tx.Create(&ledger).Error
	if err != nil {
		e.Log.Errorf("Create inviter %s ledger error:%s \r\n", inviterId, err)
		return errors.New("创建邀请人流水记录失败")
	}

	return nil
}

// ValuesConfigStruct 面额配置结构
type ValuesConfigStruct struct {
	Fixed []float64 `json:"fixed"` // 固定面额列表
	Range *struct {
		Min float64 `json:"min"` // 最小值
		Max float64 `json:"max"` // 最大值
	} `json:"range"` // 面额区间
}

// validateDenomination 校验金额是否符合入账面额规则
func (e *OrdGiftcardWriteoffs) validateDenomination(amount float64, valuesConfig string) error {
	if valuesConfig == "" {
		// 没有配置面额规则，跳过校验
		return nil
	}

	// 解析面额配置
	var config ValuesConfigStruct
	if err := json.Unmarshal([]byte(valuesConfig), &config); err != nil {
		e.Log.Warnf("Failed to parse valuesConfig: %s, error: %s", valuesConfig, err.Error())
		// 解析失败时跳过校验，避免阻塞业务
		return nil
	}

	// 校验固定面额
	if len(config.Fixed) > 0 {
		for _, fixedValue := range config.Fixed {
			// 使用小数点后两位比较（避免浮点数精度问题）
			if fmt.Sprintf("%.2f", amount) == fmt.Sprintf("%.2f", fixedValue) {
				return nil // 匹配固定面额
			}
		}
	}

	// 校验面额区间
	if config.Range != nil {
		if amount >= config.Range.Min && amount <= config.Range.Max {
			return nil // 在区间范围内
		}
	}

	// 如果既没有固定面额也没有区间配置，跳过校验
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

// CalculateUserLocalCurrency 计算用户入账金额（辅助接口）
func (e *OrdGiftcardWriteoffs) CalculateUserLocalCurrency(c *dto.OrdGiftcardWriteoffsCalculateReq) (*dto.OrdGiftcardWriteoffsCalculateResp, error) {
	// 1. 查询订单信息
	var order models.OrdUserOrders
	err := e.Orm.Where("id = ?", c.OrderId).First(&order).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardWriteoffsService CalculateUserLocalCurrency get order error:%s", err)
		return nil, errors.New("获取订单信息失败")
	}

	// 获取用户ID
	userId := order.UserId

	// 2. 查询用户货币代码（直接从用户表读取）
	var user models.HsUsers
	err = e.Orm.Select("currency_code").Where("id = ?", userId).First(&user).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardWriteoffsService CalculateUserLocalCurrency get user error:%s", err)
		return nil, errors.New("获取用户信息失败")
	}

	userCurrencyCode := user.CurrencyCode
	if userCurrencyCode == "" {
		e.Log.Errorf("OrdGiftcardWriteoffsService CalculateUserLocalCurrency user currency code is empty for user %d", userId)
		return nil, errors.New("用户未设置货币代码")
	}

	// 3. 获取礼品卡区域货币（必须提供礼品卡ID）
	var sourceCurrencyCode string
	var giftcard models.OrdGiftcard
	var hasGiftcardConfig bool
	var usedDiscountRate string // 记录使用的折扣率

	if c.GiftCardId <= 0 {
		return nil, errors.New("必须提供礼品卡ID")
	}

	// 查询礼品卡配置
	err = e.Orm.Where("id = ?", c.GiftCardId).First(&giftcard).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardWriteoffsService CalculateUserLocalCurrency get giftcard error:%s", err)
		return nil, errors.New("查询礼品卡配置失败")
	}
	hasGiftcardConfig = true

	// 确定使用哪个折扣率：优先使用传入的参数，否则使用配置的折扣率
	if c.DiscountRate != "" {
		usedDiscountRate = c.DiscountRate
		e.Log.Infof("Using discount rate from request parameter: %s", usedDiscountRate)
	} else {
		usedDiscountRate = giftcard.DiscountRate
		e.Log.Infof("Using discount rate from config: %s", usedDiscountRate)
	}

	// 查询礼品卡区域获取货币代码
	if giftcard.RegionId == "" || giftcard.RegionId == "0" {
		return nil, errors.New("礼品卡未配置区域")
	}

	var region models.OrdGiftcardRegion
	regionId, _ := strconv.Atoi(giftcard.RegionId)
	err = e.Orm.Where("id = ?", regionId).First(&region).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardWriteoffsService CalculateUserLocalCurrency get region error:%s", err)
		return nil, errors.New("查询礼品卡区域失败")
	}

	if region.CurrencyCode == "" {
		return nil, errors.New("礼品卡区域未配置货币代码")
	}

	sourceCurrencyCode = region.CurrencyCode
	e.Log.Infof("Using gift card region currency: %s (region_id: %s)", sourceCurrencyCode, giftcard.RegionId)

	// 4. 计算配置汇率和目标货币（使用礼品卡区域货币）

	isCrypto := order.BalanceType == 2
	var configRate string
	var targetCurrency string

	if isCrypto {
		targetCurrency = "USD"
		if sourceCurrencyCode == "USD" {
			configRate = "1.00000000"
		} else {
			rate, err := e.getCurrencyRate(e.Orm, sourceCurrencyCode, "USD")
			if err != nil {
				return nil, err
			}
			configRate = rate
		}
	} else {
		targetCurrency = userCurrencyCode
		if sourceCurrencyCode == userCurrencyCode {
			configRate = "1.00000000"
		} else {
			rate, err := e.getCurrencyRateViaUSD(e.Orm, sourceCurrencyCode, userCurrencyCode)
			if err != nil {
				return nil, err
			}
			configRate = rate
		}
	}

	// 5. 计算用户本地货币金额（使用折扣率）
	cardValueFloat, err := strconv.ParseFloat(c.RecognizedCardValue, 64)
	if err != nil {
		e.Log.Errorf("OrdGiftcardWriteoffsService CalculateUserLocalCurrency parse card value error:%s", err)
		return nil, errors.New("卡片面值格式错误")
	}

	// 解析折扣率
	var discountRateFloat float64 = 1.0 // 默认折扣率为1.0（无折扣）
	if usedDiscountRate != "" && usedDiscountRate != "0" {
		parsedRate, err := strconv.ParseFloat(usedDiscountRate, 64)
		if err == nil && parsedRate > 0 {
			discountRateFloat = parsedRate
		}
	}

	// 用户入账金额 = 卡片面值 × 折扣率 × 汇率
	configRateFloat, _ := strconv.ParseFloat(configRate, 64)
	userLocalAmount := cardValueFloat * discountRateFloat * configRateFloat
	userLocalCurrencyAmount := fmt.Sprintf("%.8f", userLocalAmount)

	// 6. 如果提供了礼品卡配置，进行面额校验
	var denominationValidation *dto.OrdGiftcardWriteoffsDenominationValidation
	if hasGiftcardConfig && giftcard.ValuesConfig != "" {
		// 执行面额校验
		validationErr := e.validateDenomination(userLocalAmount, giftcard.ValuesConfig)

		denominationValidation = &dto.OrdGiftcardWriteoffsDenominationValidation{
			IsValid: validationErr == nil,
		}

		if validationErr != nil {
			denominationValidation.ErrorMessage = validationErr.Error()
		}

		// 解析配置以返回允许的面额信息
		var config ValuesConfigStruct
		if err := json.Unmarshal([]byte(giftcard.ValuesConfig), &config); err == nil {
			// 固定面额
			if len(config.Fixed) > 0 {
				denominationValidation.AllowedFixed = make([]string, len(config.Fixed))
				for i, v := range config.Fixed {
					denominationValidation.AllowedFixed[i] = fmt.Sprintf("%.2f", v)
				}
			}
			// 面额区间
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

	// 7. 构建响应
	resp := &dto.OrdGiftcardWriteoffsCalculateResp{
		UserLocalCurrencyAmount: userLocalCurrencyAmount,
		UserCurrencyCode:        targetCurrency,
		ConfigRate:              configRate,
		DiscountRate:            usedDiscountRate,        // 返回使用的折扣率
		IsCrypto:                isCrypto,
		OrderCurrencyCode:       sourceCurrencyCode,      // 使用实际的源货币代码（可能是礼品卡区域货币）
		DenominationValidation:  denominationValidation,
	}

	return resp, nil
}
