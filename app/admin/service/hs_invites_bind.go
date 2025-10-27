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

type HsInvitesBind struct {
	service.Service
}

// GetPage 获取HsInvitesBind列表
func (e *HsInvitesBind) GetPage(c *dto.HsInvitesBindGetPageReq, p *actions.DataPermission, list *[]models.HsInvitesBind, count *int64) error {
	var err error
	var data models.HsInvitesBind

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsInvitesBindService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsInvitesBind对象
func (e *HsInvitesBind) Get(d *dto.HsInvitesBindGetReq, p *actions.DataPermission, model *models.HsInvitesBind) error {
	var data models.HsInvitesBind

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsInvitesBind error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsInvitesBind对象
func (e *HsInvitesBind) Insert(c *dto.HsInvitesBindInsertReq) error {
    var err error
    var data models.HsInvitesBind
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsInvitesBindService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsInvitesBind对象
func (e *HsInvitesBind) Update(c *dto.HsInvitesBindUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsInvitesBind{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsInvitesBindService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsInvitesBind
func (e *HsInvitesBind) Remove(d *dto.HsInvitesBindDeleteReq, p *actions.DataPermission) error {
	var data models.HsInvitesBind

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsInvitesBind error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
