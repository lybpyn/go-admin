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

type NoNotificationReceivers struct {
	service.Service
}

// GetPage 获取NoNotificationReceivers列表
func (e *NoNotificationReceivers) GetPage(c *dto.NoNotificationReceiversGetPageReq, p *actions.DataPermission, list *[]models.NoNotificationReceivers, count *int64) error {
	var err error
	var data models.NoNotificationReceivers

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("NoNotificationReceiversService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取NoNotificationReceivers对象
func (e *NoNotificationReceivers) Get(d *dto.NoNotificationReceiversGetReq, p *actions.DataPermission, model *models.NoNotificationReceivers) error {
	var data models.NoNotificationReceivers

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetNoNotificationReceivers error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建NoNotificationReceivers对象
func (e *NoNotificationReceivers) Insert(c *dto.NoNotificationReceiversInsertReq) error {
    var err error
    var data models.NoNotificationReceivers
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("NoNotificationReceiversService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改NoNotificationReceivers对象
func (e *NoNotificationReceivers) Update(c *dto.NoNotificationReceiversUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.NoNotificationReceivers{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("NoNotificationReceiversService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除NoNotificationReceivers
func (e *NoNotificationReceivers) Remove(d *dto.NoNotificationReceiversDeleteReq, p *actions.DataPermission) error {
	var data models.NoNotificationReceivers

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveNoNotificationReceivers error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
