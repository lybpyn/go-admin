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

type OrdUserOrders struct {
	service.Service
}

// GetPage 获取OrdUserOrders列表
func (e *OrdUserOrders) GetPage(c *dto.OrdUserOrdersGetPageReq, p *actions.DataPermission, list *[]models.OrdUserOrders, count *int64) error {
	var err error
	var data models.OrdUserOrders

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdUserOrdersService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdUserOrders对象
func (e *OrdUserOrders) Get(d *dto.OrdUserOrdersGetReq, p *actions.DataPermission, model *models.OrdUserOrders) error {
	var data models.OrdUserOrders

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdUserOrders error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrdUserOrders对象
func (e *OrdUserOrders) Insert(c *dto.OrdUserOrdersInsertReq) error {
    var err error
    var data models.OrdUserOrders
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdUserOrdersService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改OrdUserOrders对象
func (e *OrdUserOrders) Update(c *dto.OrdUserOrdersUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.OrdUserOrders{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrdUserOrdersService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除OrdUserOrders
func (e *OrdUserOrders) Remove(d *dto.OrdUserOrdersDeleteReq, p *actions.DataPermission) error {
	var data models.OrdUserOrders

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrdUserOrders error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
