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

type HsUserExperienceLogs struct {
	service.Service
}

// GetPage 获取HsUserExperienceLogs列表
func (e *HsUserExperienceLogs) GetPage(c *dto.HsUserExperienceLogsGetPageReq, p *actions.DataPermission, list *[]models.HsUserExperienceLogs, count *int64) error {
	var err error
	var data models.HsUserExperienceLogs

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserExperienceLogsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUserExperienceLogs对象
func (e *HsUserExperienceLogs) Get(d *dto.HsUserExperienceLogsGetReq, p *actions.DataPermission, model *models.HsUserExperienceLogs) error {
	var data models.HsUserExperienceLogs

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserExperienceLogs error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserExperienceLogs对象
func (e *HsUserExperienceLogs) Insert(c *dto.HsUserExperienceLogsInsertReq) error {
    var err error
    var data models.HsUserExperienceLogs
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserExperienceLogsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserExperienceLogs对象
func (e *HsUserExperienceLogs) Update(c *dto.HsUserExperienceLogsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserExperienceLogs{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserExperienceLogsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserExperienceLogs
func (e *HsUserExperienceLogs) Remove(d *dto.HsUserExperienceLogsDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserExperienceLogs

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUserExperienceLogs error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
