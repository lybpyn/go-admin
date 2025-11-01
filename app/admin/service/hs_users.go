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

type HsUsers struct {
	service.Service
}

// GetPage 获取HsUsers列表
func (e *HsUsers) GetPage(c *dto.HsUsersGetPageReq, p *actions.DataPermission, list *[]models.HsUsers, count *int64) error {
	var err error
	var data models.HsUsers

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUsersService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUsers对象
func (e *HsUsers) Get(d *dto.HsUsersGetReq, p *actions.DataPermission, model *models.HsUsers) error {
	var data models.HsUsers

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUsers error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUsers对象
func (e *HsUsers) Insert(c *dto.HsUsersInsertReq) error {
    var err error
    var data models.HsUsers
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUsersService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUsers对象
func (e *HsUsers) Update(c *dto.HsUsersUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUsers{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUsersService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUsers
func (e *HsUsers) Remove(d *dto.HsUsersDeleteReq, p *actions.DataPermission) error {
	var data models.HsUsers

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUsers error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
