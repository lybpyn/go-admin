package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type HsConfigWithdrawRules struct {
	service.Service
}

// GetPage 获取HsConfigWithdrawRules列表
func (e *HsConfigWithdrawRules) GetPage(c *dto.HsConfigWithdrawRulesGetPageReq, p *actions.DataPermission, list *[]models.HsConfigWithdrawRules, count *int64) error {
	var err error
	var data models.HsConfigWithdrawRules

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsConfigWithdrawRulesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsConfigWithdrawRules对象
func (e *HsConfigWithdrawRules) Get(d *dto.HsConfigWithdrawRulesGetReq, p *actions.DataPermission, model *models.HsConfigWithdrawRules) error {
	var data models.HsConfigWithdrawRules

	err := e.Orm.Model(&data).
		Preload("FeeTiers", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC")
		}).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsConfigWithdrawRules error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsConfigWithdrawRules对象
func (e *HsConfigWithdrawRules) Insert(c *dto.HsConfigWithdrawRulesInsertReq) error {
	var err error
	var data models.HsConfigWithdrawRules
	c.Generate(&data)

	// 开启事务
	tx := e.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 创建提现规则
	err = tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		e.Log.Errorf("HsConfigWithdrawRulesService Insert error:%s \r\n", err)
		return err
	}

	// 如果是阶梯收费类型且有阶梯配置，则保存阶梯配置
	if c.FeeType == "tiered" && len(c.FeeTiers) > 0 {
		for _, tier := range c.FeeTiers {
			feeTier := models.HsConfigWithdrawFeeTiers{
				RuleId:    data.Id,
				MinAmount: tier.MinAmount,
				MaxAmount: tier.MaxAmount,
				FeeAmount: tier.FeeAmount,
				SortOrder: tier.SortOrder,
			}
			if err = tx.Create(&feeTier).Error; err != nil {
				tx.Rollback()
				e.Log.Errorf("HsConfigWithdrawRulesService Insert fee tiers error:%s \r\n", err)
				return err
			}
		}
	}

	return tx.Commit().Error
}

// Update 修改HsConfigWithdrawRules对象
func (e *HsConfigWithdrawRules) Update(c *dto.HsConfigWithdrawRulesUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.HsConfigWithdrawRules{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	// 开启事务
	tx := e.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新提现规则
	db := tx.Save(&data)
	if err = db.Error; err != nil {
		tx.Rollback()
		e.Log.Errorf("HsConfigWithdrawRulesService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("无权更新该数据")
	}

	// 如果是阶梯收费类型，更新阶梯配置
	if c.FeeType == "tiered" {
		// 先删除旧的阶梯配置
		var feeTier models.HsConfigWithdrawFeeTiers
		if err = tx.Where("rule_id = ?", data.Id).Delete(&feeTier).Error; err != nil {
			tx.Rollback()
			e.Log.Errorf("HsConfigWithdrawRulesService delete old fee tiers error:%s \r\n", err)
			return err
		}

		// 保存新的阶梯配置
		for _, tier := range c.FeeTiers {
			newTier := models.HsConfigWithdrawFeeTiers{
				RuleId:    data.Id,
				MinAmount: tier.MinAmount,
				MaxAmount: tier.MaxAmount,
				FeeAmount: tier.FeeAmount,
				SortOrder: tier.SortOrder,
			}
			if err = tx.Create(&newTier).Error; err != nil {
				tx.Rollback()
				e.Log.Errorf("HsConfigWithdrawRulesService insert fee tiers error:%s \r\n", err)
				return err
			}
		}
	} else {
		// 如果不是阶梯收费类型，删除可能存在的阶梯配置
		var feeTier models.HsConfigWithdrawFeeTiers
		if err = tx.Where("rule_id = ?", data.Id).Delete(&feeTier).Error; err != nil {
			tx.Rollback()
			e.Log.Errorf("HsConfigWithdrawRulesService delete fee tiers error:%s \r\n", err)
			return err
		}
	}

	return tx.Commit().Error
}

// Remove 删除HsConfigWithdrawRules
func (e *HsConfigWithdrawRules) Remove(d *dto.HsConfigWithdrawRulesDeleteReq, p *actions.DataPermission) error {
	var data models.HsConfigWithdrawRules

	// 开启事务
	tx := e.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 先删除关联的阶梯配置
	ids := d.GetId().([]int)
	for _, id := range ids {
		var feeTier models.HsConfigWithdrawFeeTiers
		if err := tx.Where("rule_id = ?", id).Delete(&feeTier).Error; err != nil {
			tx.Rollback()
			e.Log.Errorf("Service RemoveHsConfigWithdrawFeeTiers error:%s \r\n", err)
			return err
		}
	}

	// 删除提现规则
	db := tx.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		tx.Rollback()
		e.Log.Errorf("Service RemoveHsConfigWithdrawRules error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("无权删除该数据")
	}

	return tx.Commit().Error
}
