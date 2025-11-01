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

type MdUserDeviceTab struct {
	service.Service
}

// GetPage 获取MdUserDeviceTab列表
func (e *MdUserDeviceTab) GetPage(c *dto.MdUserDeviceTabGetPageReq, p *actions.DataPermission, list *[]models.MdUserDeviceTab, count *int64) error {
	var err error
	var data models.MdUserDeviceTab

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("MdUserDeviceTabService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取MdUserDeviceTab对象
func (e *MdUserDeviceTab) Get(d *dto.MdUserDeviceTabGetReq, p *actions.DataPermission, model *models.MdUserDeviceTab) error {
	var data models.MdUserDeviceTab

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetMdUserDeviceTab error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建MdUserDeviceTab对象
func (e *MdUserDeviceTab) Insert(c *dto.MdUserDeviceTabInsertReq) error {
    var err error
    var data models.MdUserDeviceTab
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("MdUserDeviceTabService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改MdUserDeviceTab对象
func (e *MdUserDeviceTab) Update(c *dto.MdUserDeviceTabUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.MdUserDeviceTab{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("MdUserDeviceTabService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除MdUserDeviceTab
func (e *MdUserDeviceTab) Remove(d *dto.MdUserDeviceTabDeleteReq, p *actions.DataPermission) error {
	var data models.MdUserDeviceTab

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveMdUserDeviceTab error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
