package service

import (
	"errors"
	"fmt"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type HsUserLedger struct {
	service.Service
}

// GetPage 获取HsUserLedger列表
func (e *HsUserLedger) GetPage(c *dto.HsUserLedgerGetPageReq, p *actions.DataPermission, list *[]models.HsUserLedger, count *int64) error {
	var err error
	var data models.HsUserLedger

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserLedgerService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetFinanceStats 获取财务统计数据
func (e *HsUserLedger) GetFinanceStats(c *dto.HsUserLedgerFinanceStatsReq, list *[]dto.HsUserLedgerFinanceStatsResp) error {
	var err error
	var dateExpr string

	// 根据维度确定分组和格式化方式 - SELECT 和 GROUP BY 使用相同的表达式
	switch c.Dimension {
	case "day":
		dateExpr = "DATE(created_at)"
	case "week":
		// 按周统计，返回该周的第一天（周一）
		// 使用 DATE 函数减去工作日偏移，确保返回周一的日期
		dateExpr = "DATE(DATE_SUB(created_at, INTERVAL WEEKDAY(created_at) DAY))"
	case "month":
		dateExpr = "DATE_FORMAT(created_at, '%Y-%m')"
	default:
		return errors.New("invalid dimension, must be day/week/month")
	}

	query := e.Orm.Table("hs_user_ledger").
		Select(fmt.Sprintf(`
			%s AS date_period,
			COALESCE(SUM(CASE WHEN biz_type = 'withdraw' AND direction = -1 THEN ABS(amount) ELSE 0 END), 0) AS total_withdraw,
			COALESCE(SUM(CASE WHEN biz_type = 'withdraw_fee' AND direction = -1 THEN ABS(amount) ELSE 0 END), 0) AS total_withdraw_fee,
			COALESCE(SUM(CASE WHEN biz_type = 'invite_commissions' AND direction = 1 THEN amount ELSE 0 END), 0) AS total_commission
		`, dateExpr)).
		Where("created_at >= ? AND created_at <= ?", c.StartDate+" 00:00:00", c.EndDate+" 23:59:59").
		Where("status = 1"). // 只统计已入账的记录
		Group(dateExpr).
		Order(dateExpr)

	err = query.Scan(list).Error
	if err != nil {
		e.Log.Errorf("HsUserLedgerService GetFinanceStats error:%s \r\n", err)
		return err
	}

	return nil
}

// Get 获取HsUserLedger对象
func (e *HsUserLedger) Get(d *dto.HsUserLedgerGetReq, p *actions.DataPermission, model *models.HsUserLedger) error {
	var data models.HsUserLedger

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsUserLedger error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsUserLedger对象
func (e *HsUserLedger) Insert(c *dto.HsUserLedgerInsertReq) error {
    var err error
    var data models.HsUserLedger
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsUserLedgerService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsUserLedger对象
func (e *HsUserLedger) Update(c *dto.HsUserLedgerUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsUserLedger{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsUserLedgerService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsUserLedger
func (e *HsUserLedger) Remove(d *dto.HsUserLedgerDeleteReq, p *actions.DataPermission) error {
	var data models.HsUserLedger

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsUserLedger error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
