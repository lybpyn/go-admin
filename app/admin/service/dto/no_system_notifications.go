package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type NoSystemNotificationsGetPageReq struct {
	dto.Pagination     `search:"-"`
    TargetUserId string `form:"targetUserId"  search:"type:exact;column:target_user_id;table:no_system_notifications" comment:"目标用户ID（仅 target_type=user 时有效）"`
    NoSystemNotificationsOrder
}

type NoSystemNotificationsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:no_system_notifications"`
    Title string `form:"titleOrder"  search:"type:order;column:title;table:no_system_notifications"`
    Content string `form:"contentOrder"  search:"type:order;column:content;table:no_system_notifications"`
    Type string `form:"typeOrder"  search:"type:order;column:type;table:no_system_notifications"`
    TargetType string `form:"targetTypeOrder"  search:"type:order;column:target_type;table:no_system_notifications"`
    TargetUserId string `form:"targetUserIdOrder"  search:"type:order;column:target_user_id;table:no_system_notifications"`
    IsRead string `form:"isReadOrder"  search:"type:order;column:is_read;table:no_system_notifications"`
    ReadAt string `form:"readAtOrder"  search:"type:order;column:read_at;table:no_system_notifications"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:no_system_notifications"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:no_system_notifications"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:no_system_notifications"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:no_system_notifications"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:no_system_notifications"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:no_system_notifications"`
    
}

func (m *NoSystemNotificationsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type NoSystemNotificationsInsertReq struct {
    Id int `json:"-" comment:"消息ID"` // 消息ID
    Title string `json:"title" comment:"消息标题"`
    Content string `json:"content" comment:"消息内容"`
    Type string `json:"type" comment:"消息类型"`
    TargetType string `json:"targetType" comment:"目标类型：all=全体用户, user=单用户"`
    TargetUserId string `json:"targetUserId" comment:"目标用户ID（仅 target_type=user 时有效）"`
    IsRead string `json:"isRead" comment:"是否已读：0=未读,1=已读"`
    ReadAt time.Time `json:"readAt" comment:"读取时间"`
    Status string `json:"status" comment:"状态：1=正常，0=禁用/撤回"`
    common.ControlBy
}

func (s *NoSystemNotificationsInsertReq) Generate(model *models.NoSystemNotifications)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Title = s.Title
    model.Content = s.Content
    model.Type = s.Type
    model.TargetType = s.TargetType
    model.TargetUserId = s.TargetUserId
    model.IsRead = s.IsRead
    model.ReadAt = s.ReadAt
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *NoSystemNotificationsInsertReq) GetId() interface{} {
	return s.Id
}

type NoSystemNotificationsUpdateReq struct {
    Id int `uri:"id" comment:"消息ID"` // 消息ID
    Title string `json:"title" comment:"消息标题"`
    Content string `json:"content" comment:"消息内容"`
    Type string `json:"type" comment:"消息类型"`
    TargetType string `json:"targetType" comment:"目标类型：all=全体用户, user=单用户"`
    TargetUserId string `json:"targetUserId" comment:"目标用户ID（仅 target_type=user 时有效）"`
    IsRead string `json:"isRead" comment:"是否已读：0=未读,1=已读"`
    ReadAt time.Time `json:"readAt" comment:"读取时间"`
    Status string `json:"status" comment:"状态：1=正常，0=禁用/撤回"`
    common.ControlBy
}

func (s *NoSystemNotificationsUpdateReq) Generate(model *models.NoSystemNotifications)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Title = s.Title
    model.Content = s.Content
    model.Type = s.Type
    model.TargetType = s.TargetType
    model.TargetUserId = s.TargetUserId
    model.IsRead = s.IsRead
    model.ReadAt = s.ReadAt
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *NoSystemNotificationsUpdateReq) GetId() interface{} {
	return s.Id
}

// NoSystemNotificationsGetReq 功能获取请求参数
type NoSystemNotificationsGetReq struct {
     Id int `uri:"id"`
}
func (s *NoSystemNotificationsGetReq) GetId() interface{} {
	return s.Id
}

// NoSystemNotificationsDeleteReq 功能删除请求参数
type NoSystemNotificationsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *NoSystemNotificationsDeleteReq) GetId() interface{} {
	return s.Ids
}
