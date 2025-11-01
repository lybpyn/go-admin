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

type OrdGiftcardCategory struct {
	service.Service
}

// GetPage 获取OrdGiftcardCategory列表
func (e *OrdGiftcardCategory) GetPage(c *dto.OrdGiftcardCategoryGetPageReq, p *actions.DataPermission, list *[]models.OrdGiftcardCategory, count *int64) error {
	var err error
	var data models.OrdGiftcardCategory

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardCategoryService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdGiftcardCategory对象
func (e *OrdGiftcardCategory) Get(d *dto.OrdGiftcardCategoryGetReq, p *actions.DataPermission, model *models.OrdGiftcardCategory) error {
	var data models.OrdGiftcardCategory

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdGiftcardCategory error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrdGiftcardCategory对象
func (e *OrdGiftcardCategory) Insert(c *dto.OrdGiftcardCategoryInsertReq) error {
    var err error
    var data models.OrdGiftcardCategory
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardCategoryService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改OrdGiftcardCategory对象
func (e *OrdGiftcardCategory) Update(c *dto.OrdGiftcardCategoryUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.OrdGiftcardCategory{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrdGiftcardCategoryService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除OrdGiftcardCategory
func (e *OrdGiftcardCategory) Remove(d *dto.OrdGiftcardCategoryDeleteReq, p *actions.DataPermission) error {
	var data models.OrdGiftcardCategory

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrdGiftcardCategory error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
