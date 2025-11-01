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

type HsUserIdentity struct {
	service.Service
}

// GetPage 获取HsUserIdentity列表
func (e *HsUserIdentity) GetPage(c *dto.HsUserIdentityGetPageReq, p *actions.DataPermission, list *[]models.HsUserIdentity, count *int64) error {
	var err error
	var data models.HsUserIdentity

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserIdentityService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUserIdentity对象
func (e *HsUserIdentity) Get(d *dto.HsUserIdentityGetReq, p *actions.DataPermission, model *models.HsUserIdentity) error {
	var data models.HsUserIdentity

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserIdentity error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserIdentity对象
func (e *HsUserIdentity) Insert(c *dto.HsUserIdentityInsertReq) error {
    var err error
    var data models.HsUserIdentity
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserIdentityService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserIdentity对象
func (e *HsUserIdentity) Update(c *dto.HsUserIdentityUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserIdentity{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserIdentityService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserIdentity
func (e *HsUserIdentity) Remove(d *dto.HsUserIdentityDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserIdentity

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUserIdentity error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
