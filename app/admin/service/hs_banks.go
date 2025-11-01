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

type HsBanks struct {
	service.Service
}

// GetPage 获取HsBanks列表
func (e *HsBanks) GetPage(c *dto.HsBanksGetPageReq, p *actions.DataPermission, list *[]models.HsBanks, count *int64) error {
	var err error
	var data models.HsBanks

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsBanksService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsBanks对象
func (e *HsBanks) Get(d *dto.HsBanksGetReq, p *actions.DataPermission, model *models.HsBanks) error {
	var data models.HsBanks

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsBanks error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsBanks对象
func (e *HsBanks) Insert(c *dto.HsBanksInsertReq) error {
    var err error
    var data models.HsBanks
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsBanksService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsBanks对象
func (e *HsBanks) Update(c *dto.HsBanksUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsBanks{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsBanksService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsBanks
func (e *HsBanks) Remove(d *dto.HsBanksDeleteReq, p *actions.DataPermission) error {
	var data models.HsBanks

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsBanks error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
