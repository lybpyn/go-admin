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

type HsConfigWithdrawFee struct {
	service.Service
}

// GetPage 获取HsConfigWithdrawFee列表
func (e *HsConfigWithdrawFee) GetPage(c *dto.HsConfigWithdrawFeeGetPageReq, p *actions.DataPermission, list *[]models.HsConfigWithdrawFee, count *int64) error {
	var err error
	var data models.HsConfigWithdrawFee

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsConfigWithdrawFeeService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsConfigWithdrawFee对象
func (e *HsConfigWithdrawFee) Get(d *dto.HsConfigWithdrawFeeGetReq, p *actions.DataPermission, model *models.HsConfigWithdrawFee) error {
	var data models.HsConfigWithdrawFee

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsConfigWithdrawFee error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsConfigWithdrawFee对象
func (e *HsConfigWithdrawFee) Insert(c *dto.HsConfigWithdrawFeeInsertReq) error {
    var err error
    var data models.HsConfigWithdrawFee
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsConfigWithdrawFeeService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsConfigWithdrawFee对象
func (e *HsConfigWithdrawFee) Update(c *dto.HsConfigWithdrawFeeUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsConfigWithdrawFee{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsConfigWithdrawFeeService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsConfigWithdrawFee
func (e *HsConfigWithdrawFee) Remove(d *dto.HsConfigWithdrawFeeDeleteReq, p *actions.DataPermission) error {
	var data models.HsConfigWithdrawFee

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsConfigWithdrawFee error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
