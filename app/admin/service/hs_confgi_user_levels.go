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

type HsConfgiUserLevels struct {
	service.Service
}

// GetPage 获取HsConfgiUserLevels列表
func (e *HsConfgiUserLevels) GetPage(c *dto.HsConfgiUserLevelsGetPageReq, p *actions.DataPermission, list *[]models.HsConfgiUserLevels, count *int64) error {
	var err error
	var data models.HsConfgiUserLevels

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsConfgiUserLevelsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsConfgiUserLevels对象
func (e *HsConfgiUserLevels) Get(d *dto.HsConfgiUserLevelsGetReq, p *actions.DataPermission, model *models.HsConfgiUserLevels) error {
	var data models.HsConfgiUserLevels

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsConfgiUserLevels error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsConfgiUserLevels对象
func (e *HsConfgiUserLevels) Insert(c *dto.HsConfgiUserLevelsInsertReq) error {
    var err error
    var data models.HsConfgiUserLevels
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsConfgiUserLevelsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsConfgiUserLevels对象
func (e *HsConfgiUserLevels) Update(c *dto.HsConfgiUserLevelsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsConfgiUserLevels{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsConfgiUserLevelsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsConfgiUserLevels
func (e *HsConfgiUserLevels) Remove(d *dto.HsConfgiUserLevelsDeleteReq, p *actions.DataPermission) error {
	var data models.HsConfgiUserLevels

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsConfgiUserLevels error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
