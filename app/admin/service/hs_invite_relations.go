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

type HsInviteRelations struct {
	service.Service
}

// GetPage 获取HsInviteRelations列表
func (e *HsInviteRelations) GetPage(c *dto.HsInviteRelationsGetPageReq, p *actions.DataPermission, list *[]models.HsInviteRelations, count *int64) error {
	var err error
	var data models.HsInviteRelations

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("HsInviteRelationsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetStatsPage 获取邀请统计列表
func (e *HsInviteRelations) GetStatsPage(c *dto.HsInviteRelationsStatsGetPageReq, list *[]dto.HsInviteRelationsStatsResp, count *int64) error {
	var err error

	// 构建基础查询
	query := e.Orm.Table("hs_invite_relations AS r").
		Select(`
			r.level1_inviter_id AS inviter_user_id,
			u.invite_code AS invite_code,
			COUNT(DISTINCT r.user_id) AS invite_count,
			COALESCE(SUM(c.commission_amount), 0) AS total_commission
		`).
		Joins("LEFT JOIN hs_users AS u ON r.level1_inviter_id = u.id").
		Joins("LEFT JOIN hs_invite_commissions AS c ON r.level1_inviter_id = c.user_id AND c.status = 1").
		Where("r.level1_inviter_id IS NOT NULL AND r.level1_inviter_id != ''").
		Group("r.level1_inviter_id, u.invite_code")

	// 添加筛选条件
	if c.InviterUserId != "" {
		query = query.Where("r.level1_inviter_id = ?", c.InviterUserId)
	}
	if c.InviteCode != "" {
		query = query.Where("u.invite_code = ?", c.InviteCode)
	}

	// 先获取总数
	err = query.Count(count).Error
	if err != nil {
		e.Log.Errorf("HsInviteRelationsService GetStatsPage count error:%s \r\n", err)
		return err
	}

	// 分页查询
	offset := (c.GetPageIndex() - 1) * c.GetPageSize()
	err = query.
		Limit(c.GetPageSize()).
		Offset(offset).
		Scan(list).Error

	if err != nil {
		e.Log.Errorf("HsInviteRelationsService GetStatsPage error:%s \r\n", err)
		return err
	}

	return nil
}

// Get 获取HsInviteRelations对象
func (e *HsInviteRelations) Get(d *dto.HsInviteRelationsGetReq, p *actions.DataPermission, model *models.HsInviteRelations) error {
	var data models.HsInviteRelations

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetHsInviteRelations error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建HsInviteRelations对象
func (e *HsInviteRelations) Insert(c *dto.HsInviteRelationsInsertReq) error {
    var err error
    var data models.HsInviteRelations
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("HsInviteRelationsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改HsInviteRelations对象
func (e *HsInviteRelations) Update(c *dto.HsInviteRelationsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.HsInviteRelations{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("HsInviteRelationsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除HsInviteRelations
func (e *HsInviteRelations) Remove(d *dto.HsInviteRelationsDeleteReq, p *actions.DataPermission) error {
	var data models.HsInviteRelations

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveHsInviteRelations error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
