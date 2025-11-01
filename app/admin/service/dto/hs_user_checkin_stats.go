package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsUserCheckinStatsGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_user_checkin_stats" comment:"用户ID"`
    HsUserCheckinStatsOrder
}

type HsUserCheckinStatsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_user_checkin_stats"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_user_checkin_stats"`
    TotalCheckinDays string `form:"totalCheckinDaysOrder"  search:"type:order;column:total_checkin_days;table:hs_user_checkin_stats"`
    CurrentConsecutiveDays string `form:"currentConsecutiveDaysOrder"  search:"type:order;column:current_consecutive_days;table:hs_user_checkin_stats"`
    MaxConsecutiveDays string `form:"maxConsecutiveDaysOrder"  search:"type:order;column:max_consecutive_days;table:hs_user_checkin_stats"`
    LastCheckinDate string `form:"lastCheckinDateOrder"  search:"type:order;column:last_checkin_date;table:hs_user_checkin_stats"`
    ThisMonthCheckinDays string `form:"thisMonthCheckinDaysOrder"  search:"type:order;column:this_month_checkin_days;table:hs_user_checkin_stats"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_user_checkin_stats"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_user_checkin_stats"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_user_checkin_stats"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_user_checkin_stats"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_user_checkin_stats"`
    
}

func (m *HsUserCheckinStatsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsUserCheckinStatsInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    UserId string `json:"userId" comment:"用户ID"`
    TotalCheckinDays string `json:"totalCheckinDays" comment:"总签到天数"`
    CurrentConsecutiveDays string `json:"currentConsecutiveDays" comment:"当前连续签到天数"`
    MaxConsecutiveDays string `json:"maxConsecutiveDays" comment:"最大连续签到天数"`
    LastCheckinDate string `json:"lastCheckinDate" comment:"最后签到日期"`
    ThisMonthCheckinDays string `json:"thisMonthCheckinDays" comment:"本月签到天数"`
    common.ControlBy
}

func (s *HsUserCheckinStatsInsertReq) Generate(model *models.HsUserCheckinStats)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.TotalCheckinDays = s.TotalCheckinDays
    model.CurrentConsecutiveDays = s.CurrentConsecutiveDays
    model.MaxConsecutiveDays = s.MaxConsecutiveDays
    model.LastCheckinDate = s.LastCheckinDate
    model.ThisMonthCheckinDays = s.ThisMonthCheckinDays
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsUserCheckinStatsInsertReq) GetId() interface{} {
	return s.Id
}

type HsUserCheckinStatsUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    UserId string `json:"userId" comment:"用户ID"`
    TotalCheckinDays string `json:"totalCheckinDays" comment:"总签到天数"`
    CurrentConsecutiveDays string `json:"currentConsecutiveDays" comment:"当前连续签到天数"`
    MaxConsecutiveDays string `json:"maxConsecutiveDays" comment:"最大连续签到天数"`
    LastCheckinDate string `json:"lastCheckinDate" comment:"最后签到日期"`
    ThisMonthCheckinDays string `json:"thisMonthCheckinDays" comment:"本月签到天数"`
    common.ControlBy
}

func (s *HsUserCheckinStatsUpdateReq) Generate(model *models.HsUserCheckinStats)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.TotalCheckinDays = s.TotalCheckinDays
    model.CurrentConsecutiveDays = s.CurrentConsecutiveDays
    model.MaxConsecutiveDays = s.MaxConsecutiveDays
    model.LastCheckinDate = s.LastCheckinDate
    model.ThisMonthCheckinDays = s.ThisMonthCheckinDays
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsUserCheckinStatsUpdateReq) GetId() interface{} {
	return s.Id
}

// HsUserCheckinStatsGetReq 功能获取请求参数
type HsUserCheckinStatsGetReq struct {
     Id int `uri:"id"`
}
func (s *HsUserCheckinStatsGetReq) GetId() interface{} {
	return s.Id
}

// HsUserCheckinStatsDeleteReq 功能删除请求参数
type HsUserCheckinStatsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsUserCheckinStatsDeleteReq) GetId() interface{} {
	return s.Ids
}
