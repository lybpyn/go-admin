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

type OrdGiftcardDiscounts struct {
	service.Service
}

// GetPage 获取OrdGiftcardDiscounts列表
func (e *OrdGiftcardDiscounts) GetPage(c *dto.OrdGiftcardDiscountsGetPageReq, p *actions.DataPermission, list *[]models.OrdGiftcardDiscounts, count *int64) error {
	var err error
	var data models.OrdGiftcardDiscounts

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardDiscountsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdGiftcardDiscounts对象
func (e *OrdGiftcardDiscounts) Get(d *dto.OrdGiftcardDiscountsGetReq, p *actions.DataPermission, model *models.OrdGiftcardDiscounts) error {
	var data models.OrdGiftcardDiscounts

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdGiftcardDiscounts error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrdGiftcardDiscounts对象
func (e *OrdGiftcardDiscounts) Insert(c *dto.OrdGiftcardDiscountsInsertReq) error {
    var err error
    var data models.OrdGiftcardDiscounts
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdGiftcardDiscountsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改OrdGiftcardDiscounts对象
func (e *OrdGiftcardDiscounts) Update(c *dto.OrdGiftcardDiscountsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.OrdGiftcardDiscounts{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrdGiftcardDiscountsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除OrdGiftcardDiscounts
func (e *OrdGiftcardDiscounts) Remove(d *dto.OrdGiftcardDiscountsDeleteReq, p *actions.DataPermission) error {
	var data models.OrdGiftcardDiscounts

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrdGiftcardDiscounts error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

// BatchUpdateDiscountRate 批量修改折扣率
func (e *OrdGiftcardDiscounts) BatchUpdateDiscountRate(c *dto.OrdGiftcardDiscountsBatchUpdateReq, p *actions.DataPermission) error {
	var data models.OrdGiftcardDiscounts

	// 使用事务处理批量更新
	tx := e.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		e.Log.Errorf("OrdGiftcardDiscountsService BatchUpdateDiscountRate transaction error:%s \r\n", err)
		return err
	}

	// 遍历每个更新项
	for _, item := range c.Items {
		// 先查询记录是否存在且有权限
		var record models.OrdGiftcardDiscounts
		err := tx.Model(&data).
			Scopes(
				actions.Permission(data.TableName(), p),
			).
			First(&record, item.Id).Error

		if err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				e.Log.Errorf("OrdGiftcardDiscountsService BatchUpdateDiscountRate record not found or no permission, id:%d \r\n", item.Id)
				return errors.New("记录不存在或无权更新")
			}
			e.Log.Errorf("OrdGiftcardDiscountsService BatchUpdateDiscountRate query error:%s \r\n", err)
			return err
		}

		// 更新折扣率
		record.DiscountRate = item.DiscountRate
		record.UpdateBy = c.UpdateBy

		err = tx.Save(&record).Error
		if err != nil {
			tx.Rollback()
			e.Log.Errorf("OrdGiftcardDiscountsService BatchUpdateDiscountRate save error:%s \r\n", err)
			return err
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		e.Log.Errorf("OrdGiftcardDiscountsService BatchUpdateDiscountRate commit error:%s \r\n", err)
		return err
	}

	return nil
}

// BatchInsert 批量新增折扣率
func (e *OrdGiftcardDiscounts) BatchInsert(c *dto.OrdGiftcardDiscountsBatchInsertReq) error {
	// 使用事务处理批量插入
	tx := e.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		e.Log.Errorf("OrdGiftcardDiscountsService BatchInsert transaction error:%s \r\n", err)
		return err
	}

	// 遍历每个插入项
	for _, item := range c.Items {
		var record models.OrdGiftcardDiscounts
		record.GiftcardId = item.GiftcardId
		record.CardType = item.CardType
		record.DiscountRate = item.DiscountRate
		record.CreateBy = c.CreateBy

		err := tx.Create(&record).Error
		if err != nil {
			tx.Rollback()
			e.Log.Errorf("OrdGiftcardDiscountsService BatchInsert create error:%s \r\n", err)
			return err
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		e.Log.Errorf("OrdGiftcardDiscountsService BatchInsert commit error:%s \r\n", err)
		return err
	}

	return nil
}
