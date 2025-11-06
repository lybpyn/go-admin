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

type OrdConfigCurrencyRates struct {
	service.Service
}

// GetPage 获取OrdConfigCurrencyRates列表
func (e *OrdConfigCurrencyRates) GetPage(c *dto.OrdConfigCurrencyRatesGetPageReq, p *actions.DataPermission, list *[]models.OrdConfigCurrencyRates, count *int64) error {
	var err error
	var data models.OrdConfigCurrencyRates

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdConfigCurrencyRatesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdConfigCurrencyRates对象
func (e *OrdConfigCurrencyRates) Get(d *dto.OrdConfigCurrencyRatesGetReq, p *actions.DataPermission, model *models.OrdConfigCurrencyRates) error {
	var data models.OrdConfigCurrencyRates

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdConfigCurrencyRates error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrdConfigCurrencyRates对象
func (e *OrdConfigCurrencyRates) Insert(c *dto.OrdConfigCurrencyRatesInsertReq) error {
    var err error
    var data models.OrdConfigCurrencyRates
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdConfigCurrencyRatesService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改OrdConfigCurrencyRates对象
func (e *OrdConfigCurrencyRates) Update(c *dto.OrdConfigCurrencyRatesUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.OrdConfigCurrencyRates{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrdConfigCurrencyRatesService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除OrdConfigCurrencyRates
func (e *OrdConfigCurrencyRates) Remove(d *dto.OrdConfigCurrencyRatesDeleteReq, p *actions.DataPermission) error {
	var data models.OrdConfigCurrencyRates

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrdConfigCurrencyRates error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
