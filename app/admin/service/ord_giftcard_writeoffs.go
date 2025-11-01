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

type OrdGiftcardWriteoffs struct {
	service.Service
}

// GetPage 获取OrdGiftcardWriteoffs列表
func (e *OrdGiftcardWriteoffs) GetPage(c *dto.OrdGiftcardWriteoffsGetPageReq, p *actions.DataPermission, list *[]models.OrdGiftcardWriteoffs, count *int64) error {
	var err error
	var data models.OrdGiftcardWriteoffs

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardWriteoffsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdGiftcardWriteoffs对象
func (e *OrdGiftcardWriteoffs) Get(d *dto.OrdGiftcardWriteoffsGetReq, p *actions.DataPermission, model *models.OrdGiftcardWriteoffs) error {
	var data models.OrdGiftcardWriteoffs

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdGiftcardWriteoffs error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrdGiftcardWriteoffs对象
func (e *OrdGiftcardWriteoffs) Insert(c *dto.OrdGiftcardWriteoffsInsertReq) error {
    var err error
    var data models.OrdGiftcardWriteoffs
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardWriteoffsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改OrdGiftcardWriteoffs对象
func (e *OrdGiftcardWriteoffs) Update(c *dto.OrdGiftcardWriteoffsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.OrdGiftcardWriteoffs{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrdGiftcardWriteoffsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除OrdGiftcardWriteoffs
func (e *OrdGiftcardWriteoffs) Remove(d *dto.OrdGiftcardWriteoffsDeleteReq, p *actions.DataPermission) error {
	var data models.OrdGiftcardWriteoffs

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrdGiftcardWriteoffs error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
