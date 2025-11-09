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

type HsConfigInviteCommission struct {
	service.Service
}

// GetPage 获取HsConfigInviteCommission列表
func (e *HsConfigInviteCommission) GetPage(c *dto.HsConfigInviteCommissionGetPageReq, p *actions.DataPermission, list *[]models.HsConfigInviteCommission, count *int64) error {
	var err error
	var data models.HsConfigInviteCommission

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsConfigInviteCommissionService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsConfigInviteCommission对象
func (e *HsConfigInviteCommission) Get(d *dto.HsConfigInviteCommissionGetReq, p *actions.DataPermission, model *models.HsConfigInviteCommission) error {
	var data models.HsConfigInviteCommission

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsConfigInviteCommission error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsConfigInviteCommission对象
func (e *HsConfigInviteCommission) Insert(c *dto.HsConfigInviteCommissionInsertReq) error {
    var err error
    var data models.HsConfigInviteCommission
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsConfigInviteCommissionService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsConfigInviteCommission对象
func (e *HsConfigInviteCommission) Update(c *dto.HsConfigInviteCommissionUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsConfigInviteCommission{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsConfigInviteCommissionService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsConfigInviteCommission
func (e *HsConfigInviteCommission) Remove(d *dto.HsConfigInviteCommissionDeleteReq, p *actions.DataPermission) error {
	var data models.HsConfigInviteCommission

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsConfigInviteCommission error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
