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

type HsUserLedger struct {
	service.Service
}

// GetPage 获取HsUserLedger列表
func (e *HsUserLedger) GetPage(c *dto.HsUserLedgerGetPageReq, p *actions.DataPermission, list *[]models.HsUserLedger, count *int64) error {
	var err error
	var data models.HsUserLedger

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserLedgerService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUserLedger对象
func (e *HsUserLedger) Get(d *dto.HsUserLedgerGetReq, p *actions.DataPermission, model *models.HsUserLedger) error {
	var data models.HsUserLedger

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserLedger error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserLedger对象
func (e *HsUserLedger) Insert(c *dto.HsUserLedgerInsertReq) error {
    var err error
    var data models.HsUserLedger
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserLedgerService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserLedger对象
func (e *HsUserLedger) Update(c *dto.HsUserLedgerUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserLedger{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserLedgerService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserLedger
func (e *HsUserLedger) Remove(d *dto.HsUserLedgerDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserLedger

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUserLedger error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
