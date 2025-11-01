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

type OrdCurrencyRates struct {
	service.Service
}

// GetPage 获取OrdCurrencyRates列表
func (e *OrdCurrencyRates) GetPage(c *dto.OrdCurrencyRatesGetPageReq, p *actions.DataPermission, list *[]models.OrdCurrencyRates, count *int64) error {
	var err error
	var data models.OrdCurrencyRates

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdCurrencyRatesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdCurrencyRates对象
func (e *OrdCurrencyRates) Get(d *dto.OrdCurrencyRatesGetReq, p *actions.DataPermission, model *models.OrdCurrencyRates) error {
	var data models.OrdCurrencyRates

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdCurrencyRates error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrdCurrencyRates对象
func (e *OrdCurrencyRates) Insert(c *dto.OrdCurrencyRatesInsertReq) error {
    var err error
    var data models.OrdCurrencyRates
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdCurrencyRatesService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改OrdCurrencyRates对象
func (e *OrdCurrencyRates) Update(c *dto.OrdCurrencyRatesUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.OrdCurrencyRates{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrdCurrencyRatesService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除OrdCurrencyRates
func (e *OrdCurrencyRates) Remove(d *dto.OrdCurrencyRatesDeleteReq, p *actions.DataPermission) error {
	var data models.OrdCurrencyRates

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrdCurrencyRates error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
