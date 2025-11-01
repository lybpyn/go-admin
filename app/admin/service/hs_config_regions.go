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

type HsConfigRegions struct {
	service.Service
}

// GetPage 获取HsConfigRegions列表
func (e *HsConfigRegions) GetPage(c *dto.HsConfigRegionsGetPageReq, p *actions.DataPermission, list *[]models.HsConfigRegions, count *int64) error {
	var err error
	var data models.HsConfigRegions

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsConfigRegionsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsConfigRegions对象
func (e *HsConfigRegions) Get(d *dto.HsConfigRegionsGetReq, p *actions.DataPermission, model *models.HsConfigRegions) error {
	var data models.HsConfigRegions

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsConfigRegions error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsConfigRegions对象
func (e *HsConfigRegions) Insert(c *dto.HsConfigRegionsInsertReq) error {
    var err error
    var data models.HsConfigRegions
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsConfigRegionsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsConfigRegions对象
func (e *HsConfigRegions) Update(c *dto.HsConfigRegionsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsConfigRegions{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsConfigRegionsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsConfigRegions
func (e *HsConfigRegions) Remove(d *dto.HsConfigRegionsDeleteReq, p *actions.DataPermission) error {
	var data models.HsConfigRegions

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsConfigRegions error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
