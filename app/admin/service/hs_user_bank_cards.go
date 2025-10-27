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

type HsUserBankCards struct {
	service.Service
}

// GetPage 获取HsUserBankCards列表
func (e *HsUserBankCards) GetPage(c *dto.HsUserBankCardsGetPageReq, p *actions.DataPermission, list *[]models.HsUserBankCards, count *int64) error {
	var err error
	var data models.HsUserBankCards

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserBankCardsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUserBankCards对象
func (e *HsUserBankCards) Get(d *dto.HsUserBankCardsGetReq, p *actions.DataPermission, model *models.HsUserBankCards) error {
	var data models.HsUserBankCards

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserBankCards error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserBankCards对象
func (e *HsUserBankCards) Insert(c *dto.HsUserBankCardsInsertReq) error {
    var err error
    var data models.HsUserBankCards
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserBankCardsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserBankCards对象
func (e *HsUserBankCards) Update(c *dto.HsUserBankCardsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserBankCards{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserBankCardsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserBankCards
func (e *HsUserBankCards) Remove(d *dto.HsUserBankCardsDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserBankCards

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUserBankCards error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
