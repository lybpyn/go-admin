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

type OrdOrderGiftcardImages struct {
	service.Service
}

// GetPage 获取OrdOrderGiftcardImages列表
func (e *OrdOrderGiftcardImages) GetPage(c *dto.OrdOrderGiftcardImagesGetPageReq, p *actions.DataPermission, list *[]models.OrdOrderGiftcardImages, count *int64) error {
	var err error
	var data models.OrdOrderGiftcardImages

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdOrderGiftcardImagesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdOrderGiftcardImages对象
func (e *OrdOrderGiftcardImages) Get(d *dto.OrdOrderGiftcardImagesGetReq, p *actions.DataPermission, model *models.OrdOrderGiftcardImages) error {
	var data models.OrdOrderGiftcardImages

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdOrderGiftcardImages error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrdOrderGiftcardImages对象
func (e *OrdOrderGiftcardImages) Insert(c *dto.OrdOrderGiftcardImagesInsertReq) error {
    var err error
    var data models.OrdOrderGiftcardImages
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdOrderGiftcardImagesService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改OrdOrderGiftcardImages对象
func (e *OrdOrderGiftcardImages) Update(c *dto.OrdOrderGiftcardImagesUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.OrdOrderGiftcardImages{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrdOrderGiftcardImagesService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除OrdOrderGiftcardImages
func (e *OrdOrderGiftcardImages) Remove(d *dto.OrdOrderGiftcardImagesDeleteReq, p *actions.DataPermission) error {
	var data models.OrdOrderGiftcardImages

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrdOrderGiftcardImages error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
