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

type OrdGiftcard struct {
	service.Service
}

// GetPage 获取OrdGiftcard列表
func (e *OrdGiftcard) GetPage(c *dto.OrdGiftcardGetPageReq, p *actions.DataPermission, list *[]models.OrdGiftcard, count *int64) error {
	var err error
	var data models.OrdGiftcard

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdGiftcard对象
func (e *OrdGiftcard) Get(d *dto.OrdGiftcardGetReq, p *actions.DataPermission, model *models.OrdGiftcard) error {
	var data models.OrdGiftcard

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdGiftcard error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrdGiftcard对象
func (e *OrdGiftcard) Insert(c *dto.OrdGiftcardInsertReq) error {
    var err error
    var data models.OrdGiftcard
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改OrdGiftcard对象
func (e *OrdGiftcard) Update(c *dto.OrdGiftcardUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.OrdGiftcard{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrdGiftcardService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除OrdGiftcard
func (e *OrdGiftcard) Remove(d *dto.OrdGiftcardDeleteReq, p *actions.DataPermission) error {
	var data models.OrdGiftcard

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrdGiftcard error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
