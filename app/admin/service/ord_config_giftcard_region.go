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

type OrdConfigGiftcardRegion struct {
	service.Service
}

// GetPage 获取OrdConfigGiftcardRegion列表
func (e *OrdConfigGiftcardRegion) GetPage(c *dto.OrdConfigGiftcardRegionGetPageReq, p *actions.DataPermission, list *[]models.OrdConfigGiftcardRegion, count *int64) error {
	var err error
	var data models.OrdConfigGiftcardRegion

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdConfigGiftcardRegionService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdConfigGiftcardRegion对象
func (e *OrdConfigGiftcardRegion) Get(d *dto.OrdConfigGiftcardRegionGetReq, p *actions.DataPermission, model *models.OrdConfigGiftcardRegion) error {
	var data models.OrdConfigGiftcardRegion

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdConfigGiftcardRegion error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrdConfigGiftcardRegion对象
func (e *OrdConfigGiftcardRegion) Insert(c *dto.OrdConfigGiftcardRegionInsertReq) error {
    var err error
    var data models.OrdConfigGiftcardRegion
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdConfigGiftcardRegionService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改OrdConfigGiftcardRegion对象
func (e *OrdConfigGiftcardRegion) Update(c *dto.OrdConfigGiftcardRegionUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.OrdConfigGiftcardRegion{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrdConfigGiftcardRegionService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除OrdConfigGiftcardRegion
func (e *OrdConfigGiftcardRegion) Remove(d *dto.OrdConfigGiftcardRegionDeleteReq, p *actions.DataPermission) error {
	var data models.OrdConfigGiftcardRegion

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrdConfigGiftcardRegion error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
