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

type HsCarouselBanner struct {
	service.Service
}

// GetPage 获取HsCarouselBanner列表
func (e *HsCarouselBanner) GetPage(c *dto.HsCarouselBannerGetPageReq, p *actions.DataPermission, list *[]models.HsCarouselBanner, count *int64) error {
	var err error
	var data models.HsCarouselBanner

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsCarouselBannerService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsCarouselBanner对象
func (e *HsCarouselBanner) Get(d *dto.HsCarouselBannerGetReq, p *actions.DataPermission, model *models.HsCarouselBanner) error {
	var data models.HsCarouselBanner

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsCarouselBanner error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsCarouselBanner对象
func (e *HsCarouselBanner) Insert(c *dto.HsCarouselBannerInsertReq) error {
    var err error
    var data models.HsCarouselBanner
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsCarouselBannerService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsCarouselBanner对象
func (e *HsCarouselBanner) Update(c *dto.HsCarouselBannerUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsCarouselBanner{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsCarouselBannerService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsCarouselBanner
func (e *HsCarouselBanner) Remove(d *dto.HsCarouselBannerDeleteReq, p *actions.DataPermission) error {
	var data models.HsCarouselBanner

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsCarouselBanner error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
