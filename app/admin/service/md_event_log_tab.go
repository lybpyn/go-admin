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

type MdEventLogTab struct {
	service.Service
}

// GetPage 获取MdEventLogTab列表
func (e *MdEventLogTab) GetPage(c *dto.MdEventLogTabGetPageReq, p *actions.DataPermission, list *[]models.MdEventLogTab, count *int64) error {
	var err error
	var data models.MdEventLogTab

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("MdEventLogTabService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取MdEventLogTab对象
func (e *MdEventLogTab) Get(d *dto.MdEventLogTabGetReq, p *actions.DataPermission, model *models.MdEventLogTab) error {
	var data models.MdEventLogTab

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetMdEventLogTab error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建MdEventLogTab对象
func (e *MdEventLogTab) Insert(c *dto.MdEventLogTabInsertReq) error {
    var err error
    var data models.MdEventLogTab
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("MdEventLogTabService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改MdEventLogTab对象
func (e *MdEventLogTab) Update(c *dto.MdEventLogTabUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.MdEventLogTab{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("MdEventLogTabService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除MdEventLogTab
func (e *MdEventLogTab) Remove(d *dto.MdEventLogTabDeleteReq, p *actions.DataPermission) error {
	var data models.MdEventLogTab

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveMdEventLogTab error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
