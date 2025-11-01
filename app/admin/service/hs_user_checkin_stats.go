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

type HsUserCheckinStats struct {
	service.Service
}

// GetPage 获取HsUserCheckinStats列表
func (e *HsUserCheckinStats) GetPage(c *dto.HsUserCheckinStatsGetPageReq, p *actions.DataPermission, list *[]models.HsUserCheckinStats, count *int64) error {
	var err error
	var data models.HsUserCheckinStats

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserCheckinStatsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUserCheckinStats对象
func (e *HsUserCheckinStats) Get(d *dto.HsUserCheckinStatsGetReq, p *actions.DataPermission, model *models.HsUserCheckinStats) error {
	var data models.HsUserCheckinStats

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserCheckinStats error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserCheckinStats对象
func (e *HsUserCheckinStats) Insert(c *dto.HsUserCheckinStatsInsertReq) error {
    var err error
    var data models.HsUserCheckinStats
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserCheckinStatsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserCheckinStats对象
func (e *HsUserCheckinStats) Update(c *dto.HsUserCheckinStatsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserCheckinStats{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserCheckinStatsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserCheckinStats
func (e *HsUserCheckinStats) Remove(d *dto.HsUserCheckinStatsDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserCheckinStats

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUserCheckinStats error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
