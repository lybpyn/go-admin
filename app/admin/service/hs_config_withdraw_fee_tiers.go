package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"
)

type HsConfigWithdrawFeeTiers struct {
	service.Service
}

// GetPage 获取阶梯手续费配置列表
func (e *HsConfigWithdrawFeeTiers) GetPage(c *dto.HsConfigWithdrawFeeTiersGetPageReq, list *[]models.HsConfigWithdrawFeeTiers, count *int64) error {
	var err error
	var data models.HsConfigWithdrawFeeTiers

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Order("sort_order ASC").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsConfigWithdrawFeeTiersService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取单条阶梯手续费配置
func (e *HsConfigWithdrawFeeTiers) Get(d *dto.HsConfigWithdrawFeeTiersGetReq, model *models.HsConfigWithdrawFeeTiers) error {
	var data models.HsConfigWithdrawFeeTiers

	err := e.Orm.Model(&data).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsConfigWithdrawFeeTiers error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetByRuleId 根据规则ID获取阶梯手续费配置列表
func (e *HsConfigWithdrawFeeTiers) GetByRuleId(ruleId int, list *[]models.HsConfigWithdrawFeeTiers) error {
	var data models.HsConfigWithdrawFeeTiers

	err := e.Orm.Model(&data).
		Where("rule_id = ?", ruleId).
		Order("sort_order ASC").
		Find(list).Error
	if err != nil {
		e.Log.Errorf("HsConfigWithdrawFeeTiersService GetByRuleId error:%s \r\n", err)
		return err
	}
	return nil
}

// Insert 创建阶梯手续费配置
func (e *HsConfigWithdrawFeeTiers) Insert(c *dto.HsConfigWithdrawFeeTiersInsertReq) error {
	var err error
	var data models.HsConfigWithdrawFeeTiers
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsConfigWithdrawFeeTiersService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改阶梯手续费配置
func (e *HsConfigWithdrawFeeTiers) Update(c *dto.HsConfigWithdrawFeeTiersUpdateReq) error {
	var err error
	var data = models.HsConfigWithdrawFeeTiers{}
	e.Orm.First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("HsConfigWithdrawFeeTiersService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除阶梯手续费配置
func (e *HsConfigWithdrawFeeTiers) Remove(d *dto.HsConfigWithdrawFeeTiersDeleteReq) error {
	var data models.HsConfigWithdrawFeeTiers

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveHsConfigWithdrawFeeTiers error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// BatchSave 批量保存阶梯手续费配置（先删除再新增）
func (e *HsConfigWithdrawFeeTiers) BatchSave(c *dto.HsConfigWithdrawFeeTiersBatchReq) error {
	tx := e.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 删除该规则ID下的所有旧配置
	var data models.HsConfigWithdrawFeeTiers
	if err := tx.Where("rule_id = ?", c.RuleId).Delete(&data).Error; err != nil {
		tx.Rollback()
		e.Log.Errorf("HsConfigWithdrawFeeTiersService BatchSave delete error:%s \r\n", err)
		return err
	}

	// 批量插入新配置
	for _, tier := range c.Tiers {
		newTier := models.HsConfigWithdrawFeeTiers{
			RuleId:    c.RuleId,
			MinAmount: tier.MinAmount,
			MaxAmount: tier.MaxAmount,
			FeeAmount: tier.FeeAmount,
			SortOrder: tier.SortOrder,
		}
		if err := tx.Create(&newTier).Error; err != nil {
			tx.Rollback()
			e.Log.Errorf("HsConfigWithdrawFeeTiersService BatchSave insert error:%s \r\n", err)
			return err
		}
	}

	return tx.Commit().Error
}

// RemoveByRuleId 根据规则ID删除所有阶梯配置
func (e *HsConfigWithdrawFeeTiers) RemoveByRuleId(ruleId int) error {
	var data models.HsConfigWithdrawFeeTiers

	db := e.Orm.Where("rule_id = ?", ruleId).Delete(&data)
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveHsConfigWithdrawFeeTiersByRuleId error:%s \r\n", err)
		return err
	}
	return nil
}
