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

type HsConfigWithdrawRules struct {
	service.Service
}

// GetPage 获取HsConfigWithdrawRules列表
func (e *HsConfigWithdrawRules) GetPage(c *dto.HsConfigWithdrawRulesGetPageReq, p *actions.DataPermission, list *[]models.HsConfigWithdrawRules, count *int64) error {
	var err error
	var data models.HsConfigWithdrawRules

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsConfigWithdrawRulesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsConfigWithdrawRules对象
func (e *HsConfigWithdrawRules) Get(d *dto.HsConfigWithdrawRulesGetReq, p *actions.DataPermission, model *models.HsConfigWithdrawRules) error {
	var data models.HsConfigWithdrawRules

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsConfigWithdrawRules error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsConfigWithdrawRules对象
func (e *HsConfigWithdrawRules) Insert(c *dto.HsConfigWithdrawRulesInsertReq) error {
    var err error
    var data models.HsConfigWithdrawRules
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsConfigWithdrawRulesService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsConfigWithdrawRules对象
func (e *HsConfigWithdrawRules) Update(c *dto.HsConfigWithdrawRulesUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsConfigWithdrawRules{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsConfigWithdrawRulesService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsConfigWithdrawRules
func (e *HsConfigWithdrawRules) Remove(d *dto.HsConfigWithdrawRulesDeleteReq, p *actions.DataPermission) error {
	var data models.HsConfigWithdrawRules

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsConfigWithdrawRules error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
