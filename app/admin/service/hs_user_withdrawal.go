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

type HsUserWithdrawal struct {
	service.Service
}

// GetPage 获取HsUserWithdrawal列表
func (e *HsUserWithdrawal) GetPage(c *dto.HsUserWithdrawalGetPageReq, p *actions.DataPermission, list *[]models.HsUserWithdrawal, count *int64) error {
	var err error
	var data models.HsUserWithdrawal

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawalService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取HsUserWithdrawal对象
func (e *HsUserWithdrawal) Get(d *dto.HsUserWithdrawalGetReq, p *actions.DataPermission, model *models.HsUserWithdrawal) error {
	var data models.HsUserWithdrawal

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserWithdrawal error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserWithdrawal对象
func (e *HsUserWithdrawal) Insert(c *dto.HsUserWithdrawalInsertReq) error {
    var err error
    var data models.HsUserWithdrawal
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawalService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserWithdrawal对象
func (e *HsUserWithdrawal) Update(c *dto.HsUserWithdrawalUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserWithdrawal{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserWithdrawalService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserWithdrawal
func (e *HsUserWithdrawal) Remove(d *dto.HsUserWithdrawalDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserWithdrawal

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUserWithdrawal error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

// GetAvailable 获取可接单的提现订单列表
func (e *HsUserWithdrawal) GetAvailable(c *dto.HsUserWithdrawalAvailableReq, list *[]models.HsUserWithdrawal, count *int64) error {
	var err error
	var data models.HsUserWithdrawal

	err = e.Orm.Model(&data).
		Where("is_claimed = ?", 0).
		Where("status IN ?", []string{"pending", "review"}).
		Order("requested_at ASC").
		Scopes(
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawalService GetAvailable error:%s \r\n", err)
		return err
	}
	return nil
}

// GetMyOrders 获取我的处理订单列表
func (e *HsUserWithdrawal) GetMyOrders(c *dto.HsUserWithdrawalMyOrdersReq, handlerId int, list *[]models.HsUserWithdrawal, count *int64) error {
	var err error
	var data models.HsUserWithdrawal

	query := e.Orm.Model(&data).Where("handler_id = ?", handlerId)

	if c.Status != "" {
		query = query.Where("status = ?", c.Status)
	}

	err = query.
		Order("claimed_at DESC").
		Scopes(
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawalService GetMyOrders error:%s \r\n", err)
		return err
	}
	return nil
}

// Claim 接单
func (e *HsUserWithdrawal) Claim(orderId int, handlerId int, handlerName string) error {
	// 1. 先检查管理员的上班状态
	var sysUser models.SysUser
	err := e.Orm.Where("user_id = ?", handlerId).First(&sysUser).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawalService Claim - get user error:%s \r\n", err)
		return errors.New("获取管理员信息失败")
	}

	if sysUser.WorkStatus != "on_duty" {
		return errors.New("您当前不在上班状态，无法接单")
	}

	// 2. 使用乐观锁更新订单（防止并发接单）
	result := e.Orm.Model(&models.HsUserWithdrawal{}).
		Where("id = ?", orderId).
		Where("is_claimed = ?", 0).
		Where("status IN ?", []string{"pending", "review"}).
		Updates(map[string]interface{}{
			"handler_id":   handlerId,
			"handler_name": handlerName,
			"claimed_at":   gorm.Expr("NOW(3)"),
			"is_claimed":   1,
			"status":       "processing",
		})

	if result.Error != nil {
		e.Log.Errorf("HsUserWithdrawalService Claim error:%s \r\n", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("接单失败：该订单已被其他管理员接单或状态已变更")
	}

	return nil
}

// Release 释放订单（取消接单）
func (e *HsUserWithdrawal) Release(orderId int, handlerId int) error {
	result := e.Orm.Model(&models.HsUserWithdrawal{}).
		Where("id = ?", orderId).
		Where("handler_id = ?", handlerId).
		Where("status = ?", "processing").
		Updates(map[string]interface{}{
			"handler_id":   nil,
			"handler_name": nil,
			"claimed_at":   nil,
			"is_claimed":   0,
			"status":       "pending",
		})

	if result.Error != nil {
		e.Log.Errorf("HsUserWithdrawalService Release error:%s \r\n", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("释放订单失败：您没有接单该订单或订单状态已变更")
	}

	return nil
}

// Complete 完成订单处理
func (e *HsUserWithdrawal) Complete(c *dto.HsUserWithdrawalCompleteReq, handlerId int) error {
	// 验证失败原因
	if c.Status == "failed" && c.Reason == "" {
		return errors.New("处理失败时必须填写失败原因")
	}

	// 更新订单状态
	updates := map[string]interface{}{
		"status":       c.Status,
		"reason":       c.Reason,
		"processed_at": gorm.Expr("NOW(3)"),
	}

	if c.ChannelTxnId != "" {
		updates["channel_txn_id"] = c.ChannelTxnId
	}

	result := e.Orm.Model(&models.HsUserWithdrawal{}).
		Where("id = ?", c.Id).
		Where("handler_id = ?", handlerId).
		Where("status = ?", "processing").
		Updates(updates)

	if result.Error != nil {
		e.Log.Errorf("HsUserWithdrawalService Complete error:%s \r\n", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("完成订单失败：您没有接单该订单或订单状态已变更")
	}

	return nil
}

