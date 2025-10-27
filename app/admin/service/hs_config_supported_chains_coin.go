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

type HsConfigSupportedChainsCoin struct {
	service.Service
}

// GetPage 获取HsConfigSupportedChainsCoin列表
func (e *HsConfigSupportedChainsCoin) GetPage(c *dto.HsConfigSupportedChainsCoinGetPageReq, p *actions.DataPermission, list *[]models.HsConfigSupportedChainsCoin, count *int64) error {
	var err error
	var data models.HsConfigSupportedChainsCoin

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsConfigSupportedChainsCoinService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsConfigSupportedChainsCoin对象
func (e *HsConfigSupportedChainsCoin) Get(d *dto.HsConfigSupportedChainsCoinGetReq, p *actions.DataPermission, model *models.HsConfigSupportedChainsCoin) error {
	var data models.HsConfigSupportedChainsCoin

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsConfigSupportedChainsCoin error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsConfigSupportedChainsCoin对象
func (e *HsConfigSupportedChainsCoin) Insert(c *dto.HsConfigSupportedChainsCoinInsertReq) error {
    var err error
    var data models.HsConfigSupportedChainsCoin
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsConfigSupportedChainsCoinService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsConfigSupportedChainsCoin对象
func (e *HsConfigSupportedChainsCoin) Update(c *dto.HsConfigSupportedChainsCoinUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsConfigSupportedChainsCoin{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsConfigSupportedChainsCoinService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsConfigSupportedChainsCoin
func (e *HsConfigSupportedChainsCoin) Remove(d *dto.HsConfigSupportedChainsCoinDeleteReq, p *actions.DataPermission) error {
	var data models.HsConfigSupportedChainsCoin

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsConfigSupportedChainsCoin error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
