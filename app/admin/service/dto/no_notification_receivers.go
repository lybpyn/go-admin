package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type NoNotificationReceiversGetPageReq struct {
	dto.Pagination     `search:"-"`
    NotificationId string `form:"notificationId"  search:"type:exact;column:notification_id;table:no_notification_receivers" comment:"消息ID"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:no_notification_receivers" comment:"接收用户ID"`
    ReadAt time.Time `form:"readAt"  search:"type:gte;column:read_at;table:no_notification_receivers" comment:"读取时间"`
    NoNotificationReceiversOrder
}

type NoNotificationReceiversOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:no_notification_receivers"`
    NotificationId string `form:"notificationIdOrder"  search:"type:order;column:notification_id;table:no_notification_receivers"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:no_notification_receivers"`
    IsRead string `form:"isReadOrder"  search:"type:order;column:is_read;table:no_notification_receivers"`
    ReadAt string `form:"readAtOrder"  search:"type:order;column:read_at;table:no_notification_receivers"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:no_notification_receivers"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:no_notification_receivers"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:no_notification_receivers"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:no_notification_receivers"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:no_notification_receivers"`
    
}

func (m *NoNotificationReceiversGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type NoNotificationReceiversInsertReq struct {
    Id int `json:"-" comment:""` // 
    NotificationId string `json:"notificationId" comment:"消息ID"`
    UserId string `json:"userId" comment:"接收用户ID"`
    IsRead string `json:"isRead" comment:"是否已读：0=未读,1=已读"`
    ReadAt time.Time `json:"readAt" comment:"读取时间"`
    common.ControlBy
}

func (s *NoNotificationReceiversInsertReq) Generate(model *models.NoNotificationReceivers)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.NotificationId = s.NotificationId
    model.UserId = s.UserId
    model.IsRead = s.IsRead
    model.ReadAt = s.ReadAt
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *NoNotificationReceiversInsertReq) GetId() interface{} {
	return s.Id
}

type NoNotificationReceiversUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    NotificationId string `json:"notificationId" comment:"消息ID"`
    UserId string `json:"userId" comment:"接收用户ID"`
    IsRead string `json:"isRead" comment:"是否已读：0=未读,1=已读"`
    ReadAt time.Time `json:"readAt" comment:"读取时间"`
    common.ControlBy
}

func (s *NoNotificationReceiversUpdateReq) Generate(model *models.NoNotificationReceivers)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.NotificationId = s.NotificationId
    model.UserId = s.UserId
    model.IsRead = s.IsRead
    model.ReadAt = s.ReadAt
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *NoNotificationReceiversUpdateReq) GetId() interface{} {
	return s.Id
}

// NoNotificationReceiversGetReq 功能获取请求参数
type NoNotificationReceiversGetReq struct {
     Id int `uri:"id"`
}
func (s *NoNotificationReceiversGetReq) GetId() interface{} {
	return s.Id
}

// NoNotificationReceiversDeleteReq 功能删除请求参数
type NoNotificationReceiversDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *NoNotificationReceiversDeleteReq) GetId() interface{} {
	return s.Ids
}
