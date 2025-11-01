package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsUserCheckinsGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_user_checkins" comment:"用户ID"`
    HsUserCheckinsOrder
}

type HsUserCheckinsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_user_checkins"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_user_checkins"`
    CheckinDate string `form:"checkinDateOrder"  search:"type:order;column:checkin_date;table:hs_user_checkins"`
    CheckinTime string `form:"checkinTimeOrder"  search:"type:order;column:checkin_time;table:hs_user_checkins"`
    ConsecutiveDays string `form:"consecutiveDaysOrder"  search:"type:order;column:consecutive_days;table:hs_user_checkins"`
    RewardPoints string `form:"rewardPointsOrder"  search:"type:order;column:reward_points;table:hs_user_checkins"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_user_checkins"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_user_checkins"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_user_checkins"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_user_checkins"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_user_checkins"`
    
}

func (m *HsUserCheckinsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsUserCheckinsInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    UserId string `json:"userId" comment:"用户ID"`
    CheckinDate string `json:"checkinDate" comment:"签到日期"`
    CheckinTime time.Time `json:"checkinTime" comment:"签到时间"`
    ConsecutiveDays string `json:"consecutiveDays" comment:"连续签到天数"`
    RewardPoints string `json:"rewardPoints" comment:"签到奖励积分"`
    common.ControlBy
}

func (s *HsUserCheckinsInsertReq) Generate(model *models.HsUserCheckins)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.CheckinDate = s.CheckinDate
    model.CheckinTime = s.CheckinTime
    model.ConsecutiveDays = s.ConsecutiveDays
    model.RewardPoints = s.RewardPoints
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsUserCheckinsInsertReq) GetId() interface{} {
	return s.Id
}

type HsUserCheckinsUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    UserId string `json:"userId" comment:"用户ID"`
    CheckinDate string `json:"checkinDate" comment:"签到日期"`
    CheckinTime time.Time `json:"checkinTime" comment:"签到时间"`
    ConsecutiveDays string `json:"consecutiveDays" comment:"连续签到天数"`
    RewardPoints string `json:"rewardPoints" comment:"签到奖励积分"`
    common.ControlBy
}

func (s *HsUserCheckinsUpdateReq) Generate(model *models.HsUserCheckins)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.CheckinDate = s.CheckinDate
    model.CheckinTime = s.CheckinTime
    model.ConsecutiveDays = s.ConsecutiveDays
    model.RewardPoints = s.RewardPoints
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsUserCheckinsUpdateReq) GetId() interface{} {
	return s.Id
}

// HsUserCheckinsGetReq 功能获取请求参数
type HsUserCheckinsGetReq struct {
     Id int `uri:"id"`
}
func (s *HsUserCheckinsGetReq) GetId() interface{} {
	return s.Id
}

// HsUserCheckinsDeleteReq 功能删除请求参数
type HsUserCheckinsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsUserCheckinsDeleteReq) GetId() interface{} {
	return s.Ids
}
