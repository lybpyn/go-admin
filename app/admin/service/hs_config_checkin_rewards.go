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

type HsConfigCheckinRewards struct {
	service.Service
}

// GetPage 获取HsConfigCheckinRewards列表
func (e *HsConfigCheckinRewards) GetPage(c *dto.HsConfigCheckinRewardsGetPageReq, p *actions.DataPermission, list *[]models.HsConfigCheckinRewards, count *int64) error {
	var err error
	var data models.HsConfigCheckinRewards

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsConfigCheckinRewardsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsConfigCheckinRewards对象
func (e *HsConfigCheckinRewards) Get(d *dto.HsConfigCheckinRewardsGetReq, p *actions.DataPermission, model *models.HsConfigCheckinRewards) error {
	var data models.HsConfigCheckinRewards

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsConfigCheckinRewards error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsConfigCheckinRewards对象
func (e *HsConfigCheckinRewards) Insert(c *dto.HsConfigCheckinRewardsInsertReq) error {
    var err error
    var data models.HsConfigCheckinRewards
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsConfigCheckinRewardsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsConfigCheckinRewards对象
func (e *HsConfigCheckinRewards) Update(c *dto.HsConfigCheckinRewardsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsConfigCheckinRewards{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsConfigCheckinRewardsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsConfigCheckinRewards
func (e *HsConfigCheckinRewards) Remove(d *dto.HsConfigCheckinRewardsDeleteReq, p *actions.DataPermission) error {
	var data models.HsConfigCheckinRewards

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsConfigCheckinRewards error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
