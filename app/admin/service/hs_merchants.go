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

type HsMerchants struct {
	service.Service
}

// GetPage 获取HsMerchants列表
func (e *HsMerchants) GetPage(c *dto.HsMerchantsGetPageReq, p *actions.DataPermission, list *[]models.HsMerchants, count *int64) error {
	var err error
	var data models.HsMerchants

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsMerchantsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsMerchants对象
func (e *HsMerchants) Get(d *dto.HsMerchantsGetReq, p *actions.DataPermission, model *models.HsMerchants) error {
	var data models.HsMerchants

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsMerchants error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsMerchants对象
func (e *HsMerchants) Insert(c *dto.HsMerchantsInsertReq) error {
    var err error
    var data models.HsMerchants
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsMerchantsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsMerchants对象
func (e *HsMerchants) Update(c *dto.HsMerchantsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsMerchants{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsMerchantsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsMerchants
func (e *HsMerchants) Remove(d *dto.HsMerchantsDeleteReq, p *actions.DataPermission) error {
	var data models.HsMerchants

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsMerchants error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
