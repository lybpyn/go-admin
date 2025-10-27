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

type MdEventDefinitionTab struct {
	service.Service
}

// GetPage 获取MdEventDefinitionTab列表
func (e *MdEventDefinitionTab) GetPage(c *dto.MdEventDefinitionTabGetPageReq, p *actions.DataPermission, list *[]models.MdEventDefinitionTab, count *int64) error {
	var err error
	var data models.MdEventDefinitionTab

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("MdEventDefinitionTabService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取MdEventDefinitionTab对象
func (e *MdEventDefinitionTab) Get(d *dto.MdEventDefinitionTabGetReq, p *actions.DataPermission, model *models.MdEventDefinitionTab) error {
	var data models.MdEventDefinitionTab

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetMdEventDefinitionTab error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建MdEventDefinitionTab对象
func (e *MdEventDefinitionTab) Insert(c *dto.MdEventDefinitionTabInsertReq) error {
    var err error
    var data models.MdEventDefinitionTab
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("MdEventDefinitionTabService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改MdEventDefinitionTab对象
func (e *MdEventDefinitionTab) Update(c *dto.MdEventDefinitionTabUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.MdEventDefinitionTab{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("MdEventDefinitionTabService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除MdEventDefinitionTab
func (e *MdEventDefinitionTab) Remove(d *dto.MdEventDefinitionTabDeleteReq, p *actions.DataPermission) error {
	var data models.MdEventDefinitionTab

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveMdEventDefinitionTab error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
