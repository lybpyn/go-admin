package models

import (

	"go-admin/common/models"

)

type MdEventLogTab struct {
    models.Model
    
    EventId string `json:"eventId" gorm:"type:varchar(64);comment:事件ID（唯一标识）"` 
    EventCode string `json:"eventCode" gorm:"type:varchar(64);comment:事件代码，如：xx_Installation, xx_register_success"` 
    EventName string `json:"eventName" gorm:"type:varchar(128);comment:事件名称，如：安装、注册成功"` 
    ModuleName string `json:"moduleName" gorm:"type:varchar(64);comment:模块名称"` 
    UserId string `json:"userId" gorm:"type:bigint(20) unsigned;comment:用户ID（未登录为0）"` 
    DeviceId string `json:"deviceId" gorm:"type:varchar(128);comment:设备唯一标识（IDFA/GAID等）"` 
    Platform string `json:"platform" gorm:"type:tinyint(4);comment:平台类型：1-TikTok, 2-Google, 3-Apple"` 
    PageName string `json:"pageName" gorm:"type:varchar(128);comment:页面名称（Activity/Fragment/Ctr）"` 
    AppVersion string `json:"appVersion" gorm:"type:varchar(32);comment:APP版本号"` 
    OsType string `json:"osType" gorm:"type:varchar(16);comment:操作系统类型（iOS/Android）"` 
    OsVersion string `json:"osVersion" gorm:"type:varchar(32);comment:操作系统版本"` 
    DeviceModel string `json:"deviceModel" gorm:"type:varchar(64);comment:设备型号"` 
    DeviceBrand string `json:"deviceBrand" gorm:"type:varchar(64);comment:设备品牌"` 
    NetworkType string `json:"networkType" gorm:"type:varchar(16);comment:网络类型（WiFi/4G/5G）"` 
    CountryCode string `json:"countryCode" gorm:"type:varchar(8);comment:国家代码"` 
    Language string `json:"language" gorm:"type:varchar(16);comment:语言"` 
    IpAddress string `json:"ipAddress" gorm:"type:varchar(64);comment:IP地址"` 
    EventParams string `json:"eventParams" gorm:"type:text;comment:事件参数（JSON格式，存储自定义参数）"` 
    SessionId string `json:"sessionId" gorm:"type:varchar(128);comment:会话ID"` 
    Referrer string `json:"referrer" gorm:"type:varchar(256);comment:来源页面"` 
    EventTime string `json:"eventTime" gorm:"type:bigint(20);comment:事件发生时间戳（毫秒）"` 
    models.ModelTime
    models.ControlBy
}

func (MdEventLogTab) TableName() string {
    return "md_event_log_tab"
}

func (e *MdEventLogTab) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *MdEventLogTab) GetId() interface{} {
	return e.Id
}