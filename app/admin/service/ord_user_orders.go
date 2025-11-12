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

type OrdUserOrders struct {
	service.Service
}

// GetPage 获取OrdUserOrders列表
func (e *OrdUserOrders) GetPage(c *dto.OrdUserOrdersGetPageReq, p *actions.DataPermission, list *[]models.OrdUserOrders, count *int64) error {
	var err error
	var data models.OrdUserOrders

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Omit("gift_card_code").
		Order("created_at DESC").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdUserOrdersService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取OrdUserOrders对象
func (e *OrdUserOrders) Get(d *dto.OrdUserOrdersGetReq, p *actions.DataPermission, model *models.OrdUserOrders) error {
	var data models.OrdUserOrders

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetOrdUserOrders error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	// 只有已完成状态（status=3）的订单才能查看兑换码
	if model.Status != 3 {
		model.GiftCardCode = ""
	}

	return nil
}

// Insert 创建OrdUserOrders对象
func (e *OrdUserOrders) Insert(c *dto.OrdUserOrdersInsertReq) error {
    var err error
    var data models.OrdUserOrders
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrdUserOrdersService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改OrdUserOrders对象
func (e *OrdUserOrders) Update(c *dto.OrdUserOrdersUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.OrdUserOrders{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())

    // 记录原始状态
    oldStatus := data.Status

    c.Generate(&data)

    // 如果订单状态从非完成状态变更为完成状态（3），自动设置 ProcessingStatus = 3
    if oldStatus != 3 && data.Status == 3 {
        data.ProcessingStatus = 3
    }

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrdUserOrdersService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除OrdUserOrders
func (e *OrdUserOrders) Remove(d *dto.OrdUserOrdersDeleteReq, p *actions.DataPermission) error {
	var data models.OrdUserOrders

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrdUserOrders error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

// GetPageByAssign 根据接单人ID和状态查询订单列表
// 只返回当前接单人的订单，并包含兑换码信息
func (e *OrdUserOrders) GetPageByAssign(c *dto.OrdUserOrdersGetByAssignReq, p *actions.DataPermission, list *[]models.OrdUserOrders, count *int64) error {
	var err error
	var data models.OrdUserOrders

	// 确保AssignBy参数不为空（必须由API层从context自动设置）
	if c.AssignBy == 0 {
		e.Log.Errorf("GetPageByAssign error: AssignBy is required")
		return errors.New("接单人ID不能为空")
	}

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Order("created_at DESC").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrdUserOrdersService GetPageByAssign error:%s \r\n", err)
		return err
	}
	return nil
}

// AcceptOrder 管理员接单
func (e *OrdUserOrders) AcceptOrder(c *dto.OrdUserOrdersAcceptReq, adminId int, adminName string, p *actions.DataPermission) error {
	var err error
	var order models.OrdUserOrders

	// 查询订单是否存在
	err = e.Orm.Model(&order).
		Scopes(
			actions.Permission(order.TableName(), p),
		).
		First(&order, c.GetId()).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("AcceptOrder error: order not found, id:%v", c.GetId())
			return errors.New("订单不存在或无权操作")
		}
		e.Log.Errorf("AcceptOrder query error:%s", err)
		return err
	}

	// 检查订单是否已被接单
	if order.AssignBy != 0 {
		e.Log.Errorf("AcceptOrder error: order already accepted, id:%v, assign_by:%d", c.GetId(), order.AssignBy)
		return errors.New("订单已被接单，无法重复接单")
	}

	// 检查订单状态是否允许接单（例如：只有已支付状态才能接单）
	if order.Status != 1 {
		e.Log.Errorf("AcceptOrder error: invalid order status, id:%v, status:%d", c.GetId(), order.Status)
		return errors.New("订单状态不允许接单")
	}

	// 更新接单信息
	updates := map[string]interface{}{
		"assign_by":              adminId,
		"assign_name":            adminName,
		"assign_type":            1, // 1=人工接单
		"processing_started_at":  gorm.Expr("NOW()"),
		"processing_status":      1, // 1=正在处理
		"update_by":              adminId,
	}

	db := e.Orm.Model(&order).
		Where("id = ?", c.GetId()).
		Where("(assign_by IS NULL OR assign_by = 0)"). // 再次确认未被接单
		Updates(updates)

	if err = db.Error; err != nil {
		e.Log.Errorf("AcceptOrder update error:%s", err)
		return err
	}

	if db.RowsAffected == 0 {
		e.Log.Errorf("AcceptOrder error: no rows affected, id:%v", c.GetId())
		return errors.New("接单失败，订单可能已被其他管理员接单")
	}

	return nil
}

// CancelAcceptOrder 取消接单
func (e *OrdUserOrders) CancelAcceptOrder(c *dto.OrdUserOrdersCancelAcceptReq, adminId int, p *actions.DataPermission) error {
	var err error
	var order models.OrdUserOrders

	// 查询订单是否存在
	err = e.Orm.Model(&order).
		Scopes(
			actions.Permission(order.TableName(), p),
		).
		First(&order, c.GetId()).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("CancelAcceptOrder error: order not found, id:%v", c.GetId())
			return errors.New("订单不存在或无权操作")
		}
		e.Log.Errorf("CancelAcceptOrder query error:%s", err)
		return err
	}

	// 检查订单是否已被接单
	if order.AssignBy == 0 {
		e.Log.Errorf("CancelAcceptOrder error: order not accepted yet, id:%v", c.GetId())
		return errors.New("订单未被接单，无法取消接单")
	}

	// 检查是否是当前管理员接的单
	if order.AssignBy != adminId {
		e.Log.Errorf("CancelAcceptOrder error: order accepted by another admin, id:%v, assign_by:%d, current_admin:%d", c.GetId(), order.AssignBy, adminId)
		return errors.New("只能取消自己接的单")
	}

	// 检查订单状态是否允许取消接单（已完成或已取消的订单不能取消接单）
	if order.Status == 3 || order.Status == 4 {
		e.Log.Errorf("CancelAcceptOrder error: invalid order status, id:%v, status:%d", c.GetId(), order.Status)
		return errors.New("订单已完成或已取消，无法取消接单")
	}

	// 更新订单信息，清空接单相关字段
	updates := map[string]interface{}{
		"assign_by":              nil,
		"assign_name":            nil,
		"assign_type":            nil,
		"processing_started_at":  nil,
		"processing_started_end": nil,
		"processing_status":      2, // 2=取消
		"update_by":              adminId,
	}

	// 如果提供了管理员备注，则更新备注
	if c.AdminRemark != "" {
		updates["admin_remark"] = c.AdminRemark
	}

	db := e.Orm.Model(&order).
		Where("id = ?", c.GetId()).
		Where("assign_by = ?", adminId). // 再次确认是当前管理员接的单
		Updates(updates)

	if err = db.Error; err != nil {
		e.Log.Errorf("CancelAcceptOrder update error:%s", err)
		return err
	}

	if db.RowsAffected == 0 {
		e.Log.Errorf("CancelAcceptOrder error: no rows affected, id:%v", c.GetId())
		return errors.New("取消接单失败，订单可能已被其他操作修改")
	}

	return nil
}
