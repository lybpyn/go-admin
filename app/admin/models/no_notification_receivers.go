package models

import (
    "time"

	"go-admin/common/models"

)

type NoNotificationReceivers struct {
    models.Model
    
    NotificationId string `json:"notificationId" gorm:"type:bigint(20) unsigned;comment:消息ID"` 
    UserId string `json:"userId" gorm:"type:bigint(20) unsigned;comment:接收用户ID"` 
    IsRead string `json:"isRead" gorm:"type:tinyint(1);comment:是否已读：0=未读,1=已读"` 
    ReadAt time.Time `json:"readAt" gorm:"type:timestamp;comment:读取时间"` 
    models.ModelTime
    models.ControlBy
}

func (NoNotificationReceivers) TableName() string {
    return "no_notification_receivers"
}

func (e *NoNotificationReceivers) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *NoNotificationReceivers) GetId() interface{} {
	return e.Id
}