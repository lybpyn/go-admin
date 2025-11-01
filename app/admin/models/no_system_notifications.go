package models

import (
    "time"

	"go-admin/common/models"

)

type NoSystemNotifications struct {
    models.Model
    
    Title string `json:"title" gorm:"type:varchar(255);comment:消息标题"` 
    Content string `json:"content" gorm:"type:text;comment:消息内容"` 
    Type string `json:"type" gorm:"type:enum('system','order','security','promotion','other');comment:消息类型"` 
    TargetType string `json:"targetType" gorm:"type:enum('all','user');comment:目标类型：all=全体用户, user=单用户"` 
    TargetUserId string `json:"targetUserId" gorm:"type:bigint(20) unsigned;comment:目标用户ID（仅 target_type=user 时有效）"` 
    IsRead string `json:"isRead" gorm:"type:tinyint(1);comment:是否已读：0=未读,1=已读"` 
    ReadAt time.Time `json:"readAt" gorm:"type:timestamp;comment:读取时间"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:状态：1=正常，0=禁用/撤回"` 
    models.ModelTime
    models.ControlBy
}

func (NoSystemNotifications) TableName() string {
    return "no_system_notifications"
}

func (e *NoSystemNotifications) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *NoSystemNotifications) GetId() interface{} {
	return e.Id
}