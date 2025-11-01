package models

import (
    "time"

	"go-admin/common/models"

)

type MdUserDeviceTab struct {
    models.Model
    
    DeviceId string `json:"deviceId" gorm:"type:varchar(128);comment:设备唯一标识"` 
    UserId string `json:"userId" gorm:"type:bigint(20) unsigned;comment:用户ID（未绑定为0）"` 
    Platform string `json:"platform" gorm:"type:tinyint(4);comment:平台类型：1-TikTok, 2-Google, 3-Apple"` 
    OsType string `json:"osType" gorm:"type:varchar(16);comment:操作系统类型（iOS/Android）"` 
    OsVersion string `json:"osVersion" gorm:"type:varchar(32);comment:操作系统版本"` 
    DeviceModel string `json:"deviceModel" gorm:"type:varchar(64);comment:设备型号"` 
    DeviceBrand string `json:"deviceBrand" gorm:"type:varchar(64);comment:设备品牌"` 
    AppVersion string `json:"appVersion" gorm:"type:varchar(32);comment:APP版本号"` 
    CountryCode string `json:"countryCode" gorm:"type:varchar(8);comment:国家代码"` 
    Language string `json:"language" gorm:"type:varchar(16);comment:语言"` 
    InstallTime time.Time `json:"installTime" gorm:"type:datetime;comment:首次安装时间"` 
    FirstOpenTime time.Time `json:"firstOpenTime" gorm:"type:datetime;comment:首次打开时间"` 
    LastActiveTime time.Time `json:"lastActiveTime" gorm:"type:datetime;comment:最后活跃时间"` 
    IsRegistered string `json:"isRegistered" gorm:"type:tinyint(4);comment:是否已注册：0-否, 1-是"` 
    RegisterTime time.Time `json:"registerTime" gorm:"type:datetime;comment:注册时间"` 
    models.ModelTime
    models.ControlBy
}

func (MdUserDeviceTab) TableName() string {
    return "md_user_device_tab"
}

func (e *MdUserDeviceTab) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *MdUserDeviceTab) GetId() interface{} {
	return e.Id
}