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

type HsInviteCommissions struct {
	service.Service
}

// GetPage 获取HsInviteCommissions列表
func (e *HsInviteCommissions) GetPage(c *dto.HsInviteCommissionsGetPageReq, p *actions.DataPermission, list *[]models.HsInviteCommissions, count *int64) error {
	var err error
	var data models.HsInviteCommissions

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsInviteCommissionsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsInviteCommissions对象
func (e *HsInviteCommissions) Get(d *dto.HsInviteCommissionsGetReq, p *actions.DataPermission, model *models.HsInviteCommissions) error {
	var data models.HsInviteCommissions

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsInviteCommissions error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsInviteCommissions对象
func (e *HsInviteCommissions) Insert(c *dto.HsInviteCommissionsInsertReq) error {
    var err error
    var data models.HsInviteCommissions
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsInviteCommissionsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsInviteCommissions对象
func (e *HsInviteCommissions) Update(c *dto.HsInviteCommissionsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsInviteCommissions{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsInviteCommissionsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsInviteCommissions
func (e *HsInviteCommissions) Remove(d *dto.HsInviteCommissionsDeleteReq, p *actions.DataPermission) error {
	var data models.HsInviteCommissions

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsInviteCommissions error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
