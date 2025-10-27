package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsConfigCheckinRewardsGetPageReq struct {
	dto.Pagination     `search:"-"`
    HsConfigCheckinRewardsOrder
}

type HsConfigCheckinRewardsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_config_checkin_rewards"`
    ConsecutiveDays string `form:"consecutiveDaysOrder"  search:"type:order;column:consecutive_days;table:hs_config_checkin_rewards"`
    ExperienceReward string `form:"experienceRewardOrder"  search:"type:order;column:experience_reward;table:hs_config_checkin_rewards"`
    ExtraRewards string `form:"extraRewardsOrder"  search:"type:order;column:extra_rewards;table:hs_config_checkin_rewards"`
    IsActive string `form:"isActiveOrder"  search:"type:order;column:is_active;table:hs_config_checkin_rewards"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_config_checkin_rewards"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_config_checkin_rewards"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_config_checkin_rewards"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_config_checkin_rewards"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_config_checkin_rewards"`
    
}

func (m *HsConfigCheckinRewardsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsConfigCheckinRewardsInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    ConsecutiveDays string `json:"consecutiveDays" comment:"连续签到天数"`
    ExperienceReward string `json:"experienceReward" comment:"经验值奖励"`
    ExtraRewards string `json:"extraRewards" comment:"额外奖励配置(JSON格式)"`
    IsActive string `json:"isActive" comment:"是否启用"`
    common.ControlBy
}

func (s *HsConfigCheckinRewardsInsertReq) Generate(model *models.HsConfigCheckinRewards)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.ConsecutiveDays = s.ConsecutiveDays
    model.ExperienceReward = s.ExperienceReward
    model.ExtraRewards = s.ExtraRewards
    model.IsActive = s.IsActive
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsConfigCheckinRewardsInsertReq) GetId() interface{} {
	return s.Id
}

type HsConfigCheckinRewardsUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    ConsecutiveDays string `json:"consecutiveDays" comment:"连续签到天数"`
    ExperienceReward string `json:"experienceReward" comment:"经验值奖励"`
    ExtraRewards string `json:"extraRewards" comment:"额外奖励配置(JSON格式)"`
    IsActive string `json:"isActive" comment:"是否启用"`
    common.ControlBy
}

func (s *HsConfigCheckinRewardsUpdateReq) Generate(model *models.HsConfigCheckinRewards)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.ConsecutiveDays = s.ConsecutiveDays
    model.ExperienceReward = s.ExperienceReward
    model.ExtraRewards = s.ExtraRewards
    model.IsActive = s.IsActive
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsConfigCheckinRewardsUpdateReq) GetId() interface{} {
	return s.Id
}

// HsConfigCheckinRewardsGetReq 功能获取请求参数
type HsConfigCheckinRewardsGetReq struct {
     Id int `uri:"id"`
}
func (s *HsConfigCheckinRewardsGetReq) GetId() interface{} {
	return s.Id
}

// HsConfigCheckinRewardsDeleteReq 功能删除请求参数
type HsConfigCheckinRewardsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsConfigCheckinRewardsDeleteReq) GetId() interface{} {
	return s.Ids
}
