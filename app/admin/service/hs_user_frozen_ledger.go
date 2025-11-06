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

type HsUserFrozenLedger struct {
	service.Service
}

// GetPage 获取HsUserFrozenLedger列表
func (e *HsUserFrozenLedger) GetPage(c *dto.HsUserFrozenLedgerGetPageReq, p *actions.DataPermission, list *[]models.HsUserFrozenLedger, count *int64) error {
	var err error
	var data models.HsUserFrozenLedger

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserFrozenLedgerService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUserFrozenLedger对象
func (e *HsUserFrozenLedger) Get(d *dto.HsUserFrozenLedgerGetReq, p *actions.DataPermission, model *models.HsUserFrozenLedger) error {
	var data models.HsUserFrozenLedger

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserFrozenLedger error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserFrozenLedger对象
func (e *HsUserFrozenLedger) Insert(c *dto.HsUserFrozenLedgerInsertReq) error {
    var err error
    var data models.HsUserFrozenLedger
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserFrozenLedgerService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserFrozenLedger对象
func (e *HsUserFrozenLedger) Update(c *dto.HsUserFrozenLedgerUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserFrozenLedger{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserFrozenLedgerService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserFrozenLedger
func (e *HsUserFrozenLedger) Remove(d *dto.HsUserFrozenLedgerDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserFrozenLedger

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUserFrozenLedger error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
