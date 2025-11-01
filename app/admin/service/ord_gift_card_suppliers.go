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

type OrdGiftCardSuppliers struct {
	service.Service
}

// GetPage 获取OrdGiftCardSuppliers列表
func (e *OrdGiftCardSuppliers) GetPage(c *dto.OrdGiftCardSuppliersGetPageReq, p *actions.DataPermission, list *[]models.OrdGiftCardSuppliers, count *int64) error {
	var err error
	var data models.OrdGiftCardSuppliers

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdGiftCardSuppliersService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdGiftCardSuppliers对象
func (e *OrdGiftCardSuppliers) Get(d *dto.OrdGiftCardSuppliersGetReq, p *actions.DataPermission, model *models.OrdGiftCardSuppliers) error {
	var data models.OrdGiftCardSuppliers

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdGiftCardSuppliers error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrdGiftCardSuppliers对象
func (e *OrdGiftCardSuppliers) Insert(c *dto.OrdGiftCardSuppliersInsertReq) error {
    var err error
    var data models.OrdGiftCardSuppliers
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdGiftCardSuppliersService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改OrdGiftCardSuppliers对象
func (e *OrdGiftCardSuppliers) Update(c *dto.OrdGiftCardSuppliersUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.OrdGiftCardSuppliers{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrdGiftCardSuppliersService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除OrdGiftCardSuppliers
func (e *OrdGiftCardSuppliers) Remove(d *dto.OrdGiftCardSuppliersDeleteReq, p *actions.DataPermission) error {
	var data models.OrdGiftCardSuppliers

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrdGiftCardSuppliers error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
