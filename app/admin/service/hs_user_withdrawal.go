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

type HsUserWithdrawal struct {
	service.Service
}

// GetPage 获取HsUserWithdrawal列表
func (e *HsUserWithdrawal) GetPage(c *dto.HsUserWithdrawalGetPageReq, p *actions.DataPermission, list *[]models.HsUserWithdrawal, count *int64) error {
	var err error
	var data models.HsUserWithdrawal

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawalService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUserWithdrawal对象
func (e *HsUserWithdrawal) Get(d *dto.HsUserWithdrawalGetReq, p *actions.DataPermission, model *models.HsUserWithdrawal) error {
	var data models.HsUserWithdrawal

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserWithdrawal error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserWithdrawal对象
func (e *HsUserWithdrawal) Insert(c *dto.HsUserWithdrawalInsertReq) error {
    var err error
    var data models.HsUserWithdrawal
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawalService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserWithdrawal对象
func (e *HsUserWithdrawal) Update(c *dto.HsUserWithdrawalUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserWithdrawal{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserWithdrawalService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserWithdrawal
func (e *HsUserWithdrawal) Remove(d *dto.HsUserWithdrawalDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserWithdrawal

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUserWithdrawal error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
