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

type HsUserSocialTokens struct {
	service.Service
}

// GetPage 获取HsUserSocialTokens列表
func (e *HsUserSocialTokens) GetPage(c *dto.HsUserSocialTokensGetPageReq, p *actions.DataPermission, list *[]models.HsUserSocialTokens, count *int64) error {
	var err error
	var data models.HsUserSocialTokens

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserSocialTokensService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUserSocialTokens对象
func (e *HsUserSocialTokens) Get(d *dto.HsUserSocialTokensGetReq, p *actions.DataPermission, model *models.HsUserSocialTokens) error {
	var data models.HsUserSocialTokens

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserSocialTokens error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserSocialTokens对象
func (e *HsUserSocialTokens) Insert(c *dto.HsUserSocialTokensInsertReq) error {
    var err error
    var data models.HsUserSocialTokens
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserSocialTokensService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserSocialTokens对象
func (e *HsUserSocialTokens) Update(c *dto.HsUserSocialTokensUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserSocialTokens{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserSocialTokensService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserSocialTokens
func (e *HsUserSocialTokens) Remove(d *dto.HsUserSocialTokensDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserSocialTokens

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUserSocialTokens error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
