package models

import "time"

type HsTenjinReportedAt struct {
	Id         int64     `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	BizType    string    `json:"bizType" gorm:"type:varchar(32);not null;comment:业务类型"`
	BizId      string    `json:"bizId" gorm:"type:varchar(64);not null;comment:业务主键ID"`
	UserId     int64     `json:"userId" gorm:"type:bigint(20);not null;default:0;comment:用户ID"`
	EventName  string    `json:"eventName" gorm:"type:varchar(64);not null;comment:事件名称"`
	ReportedAt time.Time `json:"reportedAt" gorm:"type:datetime(3);not null;comment:上报时间"`
	CreatedAt  time.Time `json:"createdAt" gorm:"type:datetime(3);autoCreateTime;comment:创建时间"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"type:datetime(3);autoUpdateTime;comment:最后更新时间"`
}

func (HsTenjinReportedAt) TableName() string {
	return "hs_tenjin_reported_at"
}
