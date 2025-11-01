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

type HsUserCheckins struct {
	service.Service
}

// GetPage 获取HsUserCheckins列表
func (e *HsUserCheckins) GetPage(c *dto.HsUserCheckinsGetPageReq, p *actions.DataPermission, list *[]models.HsUserCheckins, count *int64) error {
	var err error
	var data models.HsUserCheckins

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserCheckinsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUserCheckins对象
func (e *HsUserCheckins) Get(d *dto.HsUserCheckinsGetReq, p *actions.DataPermission, model *models.HsUserCheckins) error {
	var data models.HsUserCheckins

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserCheckins error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserCheckins对象
func (e *HsUserCheckins) Insert(c *dto.HsUserCheckinsInsertReq) error {
    var err error
    var data models.HsUserCheckins
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserCheckinsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserCheckins对象
func (e *HsUserCheckins) Update(c *dto.HsUserCheckinsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserCheckins{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserCheckinsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserCheckins
func (e *HsUserCheckins) Remove(d *dto.HsUserCheckinsDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserCheckins

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUserCheckins error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
