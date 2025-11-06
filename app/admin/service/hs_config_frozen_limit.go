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

type HsConfigFrozenLimit struct {
	service.Service
}

// GetPage 获取HsConfigFrozenLimit列表
func (e *HsConfigFrozenLimit) GetPage(c *dto.HsConfigFrozenLimitGetPageReq, p *actions.DataPermission, list *[]models.HsConfigFrozenLimit, count *int64) error {
	var err error
	var data models.HsConfigFrozenLimit

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsConfigFrozenLimitService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsConfigFrozenLimit对象
func (e *HsConfigFrozenLimit) Get(d *dto.HsConfigFrozenLimitGetReq, p *actions.DataPermission, model *models.HsConfigFrozenLimit) error {
	var data models.HsConfigFrozenLimit

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsConfigFrozenLimit error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsConfigFrozenLimit对象
func (e *HsConfigFrozenLimit) Insert(c *dto.HsConfigFrozenLimitInsertReq) error {
    var err error
    var data models.HsConfigFrozenLimit
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsConfigFrozenLimitService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsConfigFrozenLimit对象
func (e *HsConfigFrozenLimit) Update(c *dto.HsConfigFrozenLimitUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsConfigFrozenLimit{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsConfigFrozenLimitService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsConfigFrozenLimit
func (e *HsConfigFrozenLimit) Remove(d *dto.HsConfigFrozenLimitDeleteReq, p *actions.DataPermission) error {
	var data models.HsConfigFrozenLimit

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsConfigFrozenLimit error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
