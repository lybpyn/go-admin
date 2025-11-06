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

// GetPageWithStats 获取带统计信息的HsUserWithdrawal列表
func (e *HsUserWithdrawal) GetPageWithStats(c *dto.HsUserWithdrawalGetPageReq, p *actions.DataPermission, list *[]dto.HsUserWithdrawalWithStats, count *int64) error {
	var err error
	var data models.HsUserWithdrawal
	var withdrawalList []models.HsUserWithdrawal

	// 先获取基础的提现记录列表
	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(&withdrawalList).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsUserWithdrawalService GetPageWithStats error:%s \r\n", err)
		return err
	}

	// 为每个提现记录计算统计信息
	for _, withdrawal := range withdrawalList {
		statsItem := dto.HsUserWithdrawalWithStats{
			HsUserWithdrawal: withdrawal,
		}

		// 计算累计提现金额和笔数（只统计成功的提现）
		var withdrawStats struct {
			TotalAmount string
			Count       int64
		}
		err = e.Orm.Model(&models.HsUserWithdrawal{}).
			Select("COALESCE(SUM(CAST(amount AS DECIMAL(18,2))), 0) as total_amount, COUNT(*) as count").
			Where("user_id = ? AND status = 'success'", withdrawal.UserId).
			Scan(&withdrawStats).Error
		if err != nil {
			e.Log.Errorf("Failed to calculate withdrawal stats for user %s: %s", withdrawal.UserId, err)
			withdrawStats.TotalAmount = "0.00"
			withdrawStats.Count = 0
		}

		// 计算累计售卖金额和笔数（只统计已完成的订单）
		var orderStats struct {
			TotalAmount string
			Count       int64
		}
		err = e.Orm.Model(&models.OrdUserOrders{}).
			Select("COALESCE(SUM(CAST(balance AS DECIMAL(20,8))), 0) as total_amount, COUNT(*) as count").
			Where("user_id = ? AND status = 3", withdrawal.UserId). // status=3表示已完成
			Scan(&orderStats).Error
		if err != nil {
			e.Log.Errorf("Failed to calculate order stats for user %s: %s", withdrawal.UserId, err)
			orderStats.TotalAmount = "0.00"
			orderStats.Count = 0
		}

		statsItem.TotalWithdrawn = withdrawStats.TotalAmount
		statsItem.WithdrawCount = withdrawStats.Count
		statsItem.TotalSales = orderStats.TotalAmount
		statsItem.OrderCount = orderStats.Count

		*list = append(*list, statsItem)
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
