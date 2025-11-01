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

type HsUserWithdrawAddresses struct {
	service.Service
}

// GetPage 获取HsUserWithdrawAddresses列表
func (e *HsUserWithdrawAddresses) GetPage(c *dto.HsUserWithdrawAddressesGetPageReq, p *actions.DataPermission, list *[]models.HsUserWithdrawAddresses, count *int64) error {
	var err error
	var data models.HsUserWithdrawAddresses

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawAddressesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUserWithdrawAddresses对象
func (e *HsUserWithdrawAddresses) Get(d *dto.HsUserWithdrawAddressesGetReq, p *actions.DataPermission, model *models.HsUserWithdrawAddresses) error {
	var data models.HsUserWithdrawAddresses

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserWithdrawAddresses error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserWithdrawAddresses对象
func (e *HsUserWithdrawAddresses) Insert(c *dto.HsUserWithdrawAddressesInsertReq) error {
    var err error
    var data models.HsUserWithdrawAddresses
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawAddressesService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserWithdrawAddresses对象
func (e *HsUserWithdrawAddresses) Update(c *dto.HsUserWithdrawAddressesUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserWithdrawAddresses{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserWithdrawAddressesService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserWithdrawAddresses
func (e *HsUserWithdrawAddresses) Remove(d *dto.HsUserWithdrawAddressesDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserWithdrawAddresses

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUserWithdrawAddresses error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
