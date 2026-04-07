package models

import "time"

type MdAppInstall struct {
	Id                      int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId                  *int      `gorm:"column:user_id" json:"userId"`
	AnalyticsInstallationId string    `gorm:"column:analytics_installation_id;size:64;not null" json:"analyticsInstallationId"`
	AdvertisingId           string    `gorm:"column:advertising_id;size:64" json:"advertisingId"`
	DeveloperDeviceId       string    `gorm:"column:developer_device_id;size:64" json:"developerDeviceId"`
	Platform                string    `gorm:"column:platform;size:16;not null" json:"platform"`
	DeviceModel             string    `gorm:"column:device_model;size:64" json:"deviceModel"`
	OsVersion               string    `gorm:"column:os_version;size:32" json:"osVersion"`
	OsVersionRelease        string    `gorm:"column:os_version_release;size:32" json:"osVersionRelease"`
	BuildId                 string    `gorm:"column:build_id;size:64" json:"buildId"`
	Locale                  string    `gorm:"column:locale;size:32" json:"locale"`
	AppVersion              string    `gorm:"column:app_version;size:32" json:"appVersion"`
	SdkVersion              string    `gorm:"column:sdk_version;size:32" json:"sdkVersion"`
	LimitAdTracking         *int      `gorm:"column:limit_ad_tracking" json:"limitAdTracking"`
	AdUserData              *int      `gorm:"column:ad_user_data" json:"adUserData"`
	AdPersonalization       *int      `gorm:"column:ad_personalization" json:"adPersonalization"`
	IpAddress               string    `gorm:"column:ip_address;size:45" json:"ipAddress"`
	Country                 string    `gorm:"column:country;size:16" json:"country"`
	SourceAppStore          string    `gorm:"column:source_app_store;size:32" json:"sourceAppStore"`
	CreateTime              time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime              time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
}

func (MdAppInstall) TableName() string {
	return "md_app_install"
}
