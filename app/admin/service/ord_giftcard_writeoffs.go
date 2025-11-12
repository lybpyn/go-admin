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
			"ord_giftcard.name as gift_card_name").
		Joins("LEFT JOIN ord_user_orders ON ord_giftcard_writeoffs.order_id = ord_user_orders.id").
		Joins("LEFT JOIN hs_users ON ord_giftcard_writeoffs.user_id = hs_users.id").
		Joins("LEFT JOIN ord_giftcard ON ord_giftcard_writeoffs.gift_card_id = ord_giftcard.id").
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
			"ord_giftcard.name as gift_card_name").
		Joins("LEFT JOIN ord_user_orders ON ord_giftcard_writeoffs.order_id = ord_user_orders.id").
		Joins("LEFT JOIN hs_users ON ord_giftcard_writeoffs.user_id = hs_users.id").
		Joins("LEFT JOIN ord_giftcard ON ord_giftcard_writeoffs.gift_card_id = ord_giftcard.id").
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
		).Delete(&data, d.GetId())
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

	// 使用事务处理
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		// 1. 获取用户信息以获取用户的货币代码
		var user models.HsUsers
		err = tx.Select("hs_users.*, hs_config_regions.currency_code").
			Joins("LEFT JOIN hs_config_regions ON hs_users.region_id = hs_config_regions.id").
			Where("hs_users.id = ?", c.UserId).
			First(&user).Error
		if err != nil {
			e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert get user error:%s \r\n", err)
			return errors.New("获取用户信息失败")
		}

		// 2. 获取用户货币代码
		var userCurrencyCode string
		err = tx.Table("hs_config_regions").
			Select("currency_code").
			Where("id = ?", user.RegionId).
			Scan(&userCurrencyCode).Error
		if err != nil {
			e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert get user currency error:%s \r\n", err)
			return errors.New("获取用户信息失败")
		}

		// 3. 获取订单信息以获取订单的货币代码和入账类型
		var order models.OrdUserOrders
		err = tx.Where("id = ?", c.OrderId).First(&order).Error
		if err != nil {
			e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert get order error:%s \r\n", err)
			return errors.New("获取订单信息失败")
		}

		// 4. 判断入账类型：1=法币余额，2=虚拟币余额
		isCrypto := order.BalanceType == 2

		// 5. 计算汇率
		// 如果是虚拟币，统一按USD计算
		// 如果是法币，根据订单货币和用户货币计算
		var configRate string
		var targetCurrency string

		if isCrypto {
			// 虚拟币统一按USD计算
			targetCurrency = "USD"
			if order.CurrencyCode == "USD" {
				configRate = "1.00000000"
			} else {
				// 从订单货币转换到USD
				rate, err := e.getCurrencyRate(tx, order.CurrencyCode, "USD")
				if err != nil {
					return err
				}
				configRate = rate
			}
		} else {
			// 法币按用户本地货币计算
			targetCurrency = userCurrencyCode
			if order.CurrencyCode == userCurrencyCode {
				configRate = "1.00000000"
			} else {
				// 通过USD中转：订单货币 -> USD -> 用户货币
				rate, err := e.getCurrencyRateViaUSD(tx, order.CurrencyCode, userCurrencyCode)
				if err != nil {
					return err
				}
				configRate = rate
			}
		}

		// 6. 构建批量插入数据并计算总金额
		writeoffRecords := make([]models.OrdGiftcardWriteoffs, 0, len(c.WriteoffList))
		var totalAmount float64 = 0

		for _, item := range c.WriteoffList {
			// 计算转换后的金额
			convertedAmount := "0.00000000"
			if item.RecognizedCardValue != "" {
				cardValue, err1 := strconv.ParseFloat(item.RecognizedCardValue, 64)
				rate, err2 := strconv.ParseFloat(configRate, 64)
				if err1 == nil && err2 == nil {
					amount := cardValue * rate
					convertedAmount = fmt.Sprintf("%.8f", amount)

					// 只累加状态为"已核销"（status=1）的金额
					if item.Status == 1 {
						totalAmount += amount
					}
				}
			}

			// 计算平台入账货币对应的美元汇率
			var platformToUsdRate string
			if item.PlatformSettlementCurrency != "" {
				if item.PlatformSettlementCurrency == "USD" {
					platformToUsdRate = "1.00000000"
				} else {
					// 获取平台入账货币到USD的汇率
					rate, err := e.getCurrencyRate(tx, item.PlatformSettlementCurrency, "USD")
					if err != nil {
						e.Log.Warnf("Get platform currency %s to USD rate error:%s, use 0", item.PlatformSettlementCurrency, err)
						platformToUsdRate = "0.00000000"
					} else {
						platformToUsdRate = rate
					}
				}
			}

			record := models.OrdGiftcardWriteoffs{
				UserId:                     c.UserId,
				OrderId:                    c.OrderId,
				GiftCardId:                 c.GiftCardId,
				Status:                     item.Status,
				Remark:                     item.Remark,
				AdminRecognizedCode:        item.AdminRecognizedCode,
				PlatformSaleRate:           item.PlatformSaleRate,
				RecognizedCardValue:        item.RecognizedCardValue,
				FailureImageUrl:            item.FailureImageUrl,
				SupplierId:                 item.SupplierId,
				ConfigRate:                 configRate,
				UserLocalCurrencyAmount:    convertedAmount,
				UserCurrencyCode:           targetCurrency,
				PlatformSettlementAmount:   item.PlatformSettlementAmount,
				PlatformSettlementCurrency: item.PlatformSettlementCurrency,
				PlatformToUsdRate:          platformToUsdRate,
			}
			record.CreateBy = c.CreateBy
			writeoffRecords = append(writeoffRecords, record)
		}

		// 7. 批量插入核销记录
		err = tx.Create(&writeoffRecords).Error
		if err != nil {
			e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert error:%s \r\n", err)
			return err
		}

		// 如果总金额大于0，需要进行余额增加、流水记录和账户分成
		if totalAmount > 0 {
			// 8. 更新用户余额（根据入账类型选择法币或虚拟币余额）
			var balanceBefore float64
			var balanceAfter float64
			var balanceField string
			var decimalPlaces int

			if isCrypto {
				// 虚拟币余额（8位小数）
				balanceField = "crypto_balance"
				decimalPlaces = 8
				balanceBefore, _ = strconv.ParseFloat(user.CryptoBalance, 64)
			} else {
				// 法币余额（2位小数）
				balanceField = "balance"
				decimalPlaces = 2
				balanceBefore, _ = strconv.ParseFloat(user.Balance, 64)
			}

			balanceAfter = balanceBefore + totalAmount

			result := tx.Model(&models.HsUsers{}).
				Where("id = ? AND version = ?", c.UserId, user.Version).
				Updates(map[string]interface{}{
					balanceField: fmt.Sprintf("%.*f", decimalPlaces, balanceAfter),
					"version":    gorm.Expr("version + 1"),
				})

			if result.Error != nil {
				e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert update balance error:%s \r\n", result.Error)
				return errors.New("更新用户余额失败")
			}

			if result.RowsAffected == 0 {
				e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert update balance conflict")
				return errors.New("余额更新冲突，请重试")
			}

			// 9. 创建流水记录
			bizType := "giftcard_writeoff_fiat"
			if isCrypto {
				bizType = "giftcard_writeoff_crypto"
			}

			ledger := models.HsUserLedger{
				UserId:         c.UserId,
				CurrencyCode:   targetCurrency,
				Direction:      "1", // 1=入账
				Amount:         fmt.Sprintf("%.*f", decimalPlaces, totalAmount),
				BalanceBefore:  fmt.Sprintf("%.*f", decimalPlaces, balanceBefore),
				BalanceAfter:   fmt.Sprintf("%.*f", decimalPlaces, balanceAfter),
				BizType:        bizType,
				BizId:          c.OrderId,
				IdempotencyKey: fmt.Sprintf("GIFTCARD_WRITEOFF:%s:%d", c.OrderId, time.Now().UnixNano()),
				RefTable:       "ord_giftcard_writeoffs",
				RefId:          c.OrderId,
				Remark:         fmt.Sprintf("礼品卡核销到账，订单号: %s", c.OrderId),
				Status:         "1", // 1=已入账
			}
			ledger.CreateBy = c.CreateBy

			err = tx.Create(&ledger).Error
			if err != nil {
				e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert create ledger error:%s \r\n", err)
				return errors.New("创建流水记录失败")
			}

			// 10. 处理邀请分成（分成进入冻结余额）
			err = e.processInviteCommissions(tx, c.UserId, c.OrderId, totalAmount, targetCurrency, isCrypto, c.CreateBy)
			if err != nil {
				e.Log.Errorf("OrdGiftcardWriteoffsService BatchInsert process commissions error:%s \r\n", err)
				return err
			}
		}

		return nil
	})
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
func (e *OrdGiftcardWriteoffs) processInviteCommissions(tx *gorm.DB, userId, orderId string, amount float64, sourceCurrency string, isCrypto bool, createBy int) error {
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
				OrderId:          orderId,
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
				OrderId:          orderId,
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

	// 7. 获取邀请人的等级配置，解析level_privileges获取gift_card_discount
	var frozenRate float64 = 1.0 // 默认100%进入冻结余额
	if inviter.LevelId != "" && inviter.LevelId != "0" {
		var userLevel models.HsConfgiUserLevels
		err = tx.Where("id = ? AND is_active = 1", inviter.LevelId).First(&userLevel).Error
		if err == nil && userLevel.LevelPrivileges != "" {
			var privileges map[string]interface{}
			err = json.Unmarshal([]byte(userLevel.LevelPrivileges), &privileges)
			if err == nil {
				// 获取gift_card_discount值
				if discount, ok := privileges["gift_card_discount"]; ok {
					switch v := discount.(type) {
					case float64:
						frozenRate = v
					case string:
						frozenRate, _ = strconv.ParseFloat(v, 64)
					}
				}
			}
		}
	}

	return &InviterCommissionConfig{
		CommissionRate: commissionRate,
		FrozenRate:     frozenRate,
	}, nil
}

// updateInviterFrozenBalance 更新邀请人冻结余额和流水（支持多货币汇率转换）
func (e *OrdGiftcardWriteoffs) updateInviterFrozenBalance(tx *gorm.DB, inviterId string, sourceAmount float64, sourceCurrency string, isCrypto bool, orderId, level string, createBy int) error {
	// 1. 获取邀请人信息和所在区域的货币代码
	var inviter models.HsUsers
	err := tx.Select("hs_users.*, hs_config_regions.currency_code").
		Joins("LEFT JOIN hs_config_regions ON hs_users.region_id = hs_config_regions.id").
		Where("hs_users.id = ?", inviterId).
		First(&inviter).Error
	if err != nil {
		e.Log.Errorf("Get inviter %s error:%s \r\n", inviterId, err)
		return fmt.Errorf("获取邀请人信息失败")
	}

	// 2. 使用订单用户的货币，不进行汇率转换
	convertedAmount := sourceAmount
	targetCurrency := sourceCurrency

	// 3. 更新邀请人冻结余额（使用乐观锁）
	var balanceBefore float64
	var balanceAfter float64
	var balanceField string
	var decimalPlaces int

	if isCrypto {
		// 虚拟币冻结余额（8位小数）
		balanceField = "crypto_frozen_balance"
		decimalPlaces = 8
		balanceBefore, _ = strconv.ParseFloat(inviter.CryptoFrozenBalance, 64)
	} else {
		// 法币冻结余额（2位小数）
		balanceField = "frozen_balance"
		decimalPlaces = 2
		balanceBefore, _ = strconv.ParseFloat(inviter.FrozenBalance, 64)
	}

	balanceAfter = balanceBefore + convertedAmount

	result := tx.Model(&models.HsUsers{}).
		Where("id = ? AND version = ?", inviterId, inviter.Version).
		Updates(map[string]interface{}{
			balanceField: fmt.Sprintf("%.*f", decimalPlaces, balanceAfter),
			"version":    gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		e.Log.Errorf("Update inviter %s frozen balance error:%s \r\n", inviterId, result.Error)
		return errors.New("更新邀请人冻结余额失败")
	}

	if result.RowsAffected == 0 {
		e.Log.Errorf("Update inviter %s frozen balance conflict", inviterId)
		return errors.New("邀请人冻结余额更新冲突，请重试")
	}

	// 5. 创建邀请人流水记录
	bizType := "invite_commission_frozen_fiat"
	if isCrypto {
		bizType = "invite_commission_frozen_crypto"
	}

	ledger := models.HsUserLedger{
		UserId:         inviterId,
		CurrencyCode:   targetCurrency,
		Direction:      "1", // 1=入账
		Amount:         fmt.Sprintf("%.*f", decimalPlaces, convertedAmount),
		BalanceBefore:  fmt.Sprintf("%.*f", decimalPlaces, balanceBefore),
		BalanceAfter:   fmt.Sprintf("%.*f", decimalPlaces, balanceAfter),
		BizType:        bizType,
		BizId:          orderId,
		IdempotencyKey: fmt.Sprintf("INVITE_COMMISSION_FROZEN_L%s:%s:%d", level, orderId, time.Now().UnixNano()),
		RefTable:       "hs_invite_commissions",
		RefId:          orderId,
		Remark:         fmt.Sprintf("邀请分成冻结（%s级），订单号: %s", level, orderId),
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

// updateInviterAvailableBalance 更新邀请人可用余额和流水（使用订单用户的货币）
func (e *OrdGiftcardWriteoffs) updateInviterAvailableBalance(tx *gorm.DB, inviterId string, sourceAmount float64, sourceCurrency string, isCrypto bool, orderId, level string, createBy int) error {
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
		BizId:          orderId,
		IdempotencyKey: fmt.Sprintf("INVITE_COMMISSION_AVAILABLE_L%s:%s:%d", level, orderId, time.Now().UnixNano()),
		RefTable:       "hs_invite_commissions",
		RefId:          orderId,
		Remark:         fmt.Sprintf("邀请分成可用（%s级），订单号: %s", level, orderId),
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
