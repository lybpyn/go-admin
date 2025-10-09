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

type HsReceivingAccounts struct {
	service.Service
}

// GetPage 获取HsReceivingAccounts列表
func (e *HsReceivingAccounts) GetPage(c *dto.HsReceivingAccountsGetPageReq, p *actions.DataPermission, list *[]models.HsReceivingAccounts, count *int64) error {
	var err error
	var data models.HsReceivingAccounts

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsReceivingAccountsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsReceivingAccounts对象
func (e *HsReceivingAccounts) Get(d *dto.HsReceivingAccountsGetReq, p *actions.DataPermission, model *models.HsReceivingAccounts) error {
	var data models.HsReceivingAccounts

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsReceivingAccounts error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsReceivingAccounts对象
func (e *HsReceivingAccounts) Insert(c *dto.HsReceivingAccountsInsertReq) error {
    var err error
    var data models.HsReceivingAccounts
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsReceivingAccountsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsReceivingAccounts对象
func (e *HsReceivingAccounts) Update(c *dto.HsReceivingAccountsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsReceivingAccounts{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsReceivingAccountsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsReceivingAccounts
func (e *HsReceivingAccounts) Remove(d *dto.HsReceivingAccountsDeleteReq, p *actions.DataPermission) error {
	var data models.HsReceivingAccounts

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsReceivingAccounts error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
