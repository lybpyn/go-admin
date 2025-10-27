package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsUserExperienceLogsGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_user_experience_logs" comment:"用户ID"`
    SourceType string `form:"sourceType"  search:"type:exact;column:source_type;table:hs_user_experience_logs" comment:"经验来源类型(checkin,gift_card_exchange,order_complete等)"`
    HsUserExperienceLogsOrder
}

type HsUserExperienceLogsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_user_experience_logs"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_user_experience_logs"`
    ExperienceChange string `form:"experienceChangeOrder"  search:"type:order;column:experience_change;table:hs_user_experience_logs"`
    ExperienceBefore string `form:"experienceBeforeOrder"  search:"type:order;column:experience_before;table:hs_user_experience_logs"`
    ExperienceAfter string `form:"experienceAfterOrder"  search:"type:order;column:experience_after;table:hs_user_experience_logs"`
    SourceType string `form:"sourceTypeOrder"  search:"type:order;column:source_type;table:hs_user_experience_logs"`
    SourceId string `form:"sourceIdOrder"  search:"type:order;column:source_id;table:hs_user_experience_logs"`
    Description string `form:"descriptionOrder"  search:"type:order;column:description;table:hs_user_experience_logs"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_user_experience_logs"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_user_experience_logs"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_user_experience_logs"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_user_experience_logs"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_user_experience_logs"`
    
}

func (m *HsUserExperienceLogsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsUserExperienceLogsInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    UserId string `json:"userId" comment:"用户ID"`
    ExperienceChange string `json:"experienceChange" comment:"经验值变化(正数为增加，负数为减少)"`
    ExperienceBefore string `json:"experienceBefore" comment:"变化前经验值"`
    ExperienceAfter string `json:"experienceAfter" comment:"变化后经验值"`
    SourceType string `json:"sourceType" comment:"经验来源类型(checkin,gift_card_exchange,order_complete等)"`
    SourceId string `json:"sourceId" comment:"来源记录ID"`
    Description string `json:"description" comment:"经验变化描述"`
    common.ControlBy
}

func (s *HsUserExperienceLogsInsertReq) Generate(model *models.HsUserExperienceLogs)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.ExperienceChange = s.ExperienceChange
    model.ExperienceBefore = s.ExperienceBefore
    model.ExperienceAfter = s.ExperienceAfter
    model.SourceType = s.SourceType
    model.SourceId = s.SourceId
    model.Description = s.Description
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsUserExperienceLogsInsertReq) GetId() interface{} {
	return s.Id
}

type HsUserExperienceLogsUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    UserId string `json:"userId" comment:"用户ID"`
    ExperienceChange string `json:"experienceChange" comment:"经验值变化(正数为增加，负数为减少)"`
    ExperienceBefore string `json:"experienceBefore" comment:"变化前经验值"`
    ExperienceAfter string `json:"experienceAfter" comment:"变化后经验值"`
    SourceType string `json:"sourceType" comment:"经验来源类型(checkin,gift_card_exchange,order_complete等)"`
    SourceId string `json:"sourceId" comment:"来源记录ID"`
    Description string `json:"description" comment:"经验变化描述"`
    common.ControlBy
}

func (s *HsUserExperienceLogsUpdateReq) Generate(model *models.HsUserExperienceLogs)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.ExperienceChange = s.ExperienceChange
    model.ExperienceBefore = s.ExperienceBefore
    model.ExperienceAfter = s.ExperienceAfter
    model.SourceType = s.SourceType
    model.SourceId = s.SourceId
    model.Description = s.Description
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsUserExperienceLogsUpdateReq) GetId() interface{} {
	return s.Id
}

// HsUserExperienceLogsGetReq 功能获取请求参数
type HsUserExperienceLogsGetReq struct {
     Id int `uri:"id"`
}
func (s *HsUserExperienceLogsGetReq) GetId() interface{} {
	return s.Id
}

// HsUserExperienceLogsDeleteReq 功能删除请求参数
type HsUserExperienceLogsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsUserExperienceLogsDeleteReq) GetId() interface{} {
	return s.Ids
}
