package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type MdEventLogTabGetPageReq struct {
	dto.Pagination     `search:"-"`
    EventId string `form:"eventId"  search:"type:exact;column:event_id;table:md_event_log_tab" comment:"事件ID（唯一标识）"`
    EventCode string `form:"eventCode"  search:"type:exact;column:event_code;table:md_event_log_tab" comment:"事件代码，如：xx_Installation, xx_register_success"`
    EventName string `form:"eventName"  search:"type:exact;column:event_name;table:md_event_log_tab" comment:"事件名称，如：安装、注册成功"`
    ModuleName string `form:"moduleName"  search:"type:exact;column:module_name;table:md_event_log_tab" comment:"模块名称"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:md_event_log_tab" comment:"用户ID（未登录为0）"`
    DeviceId string `form:"deviceId"  search:"type:exact;column:device_id;table:md_event_log_tab" comment:"设备唯一标识（IDFA/GAID等）"`
    MdEventLogTabOrder
}

type MdEventLogTabOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:md_event_log_tab"`
    EventId string `form:"eventIdOrder"  search:"type:order;column:event_id;table:md_event_log_tab"`
    EventCode string `form:"eventCodeOrder"  search:"type:order;column:event_code;table:md_event_log_tab"`
    EventName string `form:"eventNameOrder"  search:"type:order;column:event_name;table:md_event_log_tab"`
    ModuleName string `form:"moduleNameOrder"  search:"type:order;column:module_name;table:md_event_log_tab"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:md_event_log_tab"`
    DeviceId string `form:"deviceIdOrder"  search:"type:order;column:device_id;table:md_event_log_tab"`
    Platform string `form:"platformOrder"  search:"type:order;column:platform;table:md_event_log_tab"`
    PageName string `form:"pageNameOrder"  search:"type:order;column:page_name;table:md_event_log_tab"`
    AppVersion string `form:"appVersionOrder"  search:"type:order;column:app_version;table:md_event_log_tab"`
    OsType string `form:"osTypeOrder"  search:"type:order;column:os_type;table:md_event_log_tab"`
    OsVersion string `form:"osVersionOrder"  search:"type:order;column:os_version;table:md_event_log_tab"`
    DeviceModel string `form:"deviceModelOrder"  search:"type:order;column:device_model;table:md_event_log_tab"`
    DeviceBrand string `form:"deviceBrandOrder"  search:"type:order;column:device_brand;table:md_event_log_tab"`
    NetworkType string `form:"networkTypeOrder"  search:"type:order;column:network_type;table:md_event_log_tab"`
    CountryCode string `form:"countryCodeOrder"  search:"type:order;column:country_code;table:md_event_log_tab"`
    Language string `form:"languageOrder"  search:"type:order;column:language;table:md_event_log_tab"`
    IpAddress string `form:"ipAddressOrder"  search:"type:order;column:ip_address;table:md_event_log_tab"`
    EventParams string `form:"eventParamsOrder"  search:"type:order;column:event_params;table:md_event_log_tab"`
    SessionId string `form:"sessionIdOrder"  search:"type:order;column:session_id;table:md_event_log_tab"`
    Referrer string `form:"referrerOrder"  search:"type:order;column:referrer;table:md_event_log_tab"`
    EventTime string `form:"eventTimeOrder"  search:"type:order;column:event_time;table:md_event_log_tab"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:md_event_log_tab"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:md_event_log_tab"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:md_event_log_tab"`
    
}

func (m *MdEventLogTabGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type MdEventLogTabInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    EventId string `json:"eventId" comment:"事件ID（唯一标识）"`
    EventCode string `json:"eventCode" comment:"事件代码，如：xx_Installation, xx_register_success"`
    EventName string `json:"eventName" comment:"事件名称，如：安装、注册成功"`
    ModuleName string `json:"moduleName" comment:"模块名称"`
    UserId string `json:"userId" comment:"用户ID（未登录为0）"`
    DeviceId string `json:"deviceId" comment:"设备唯一标识（IDFA/GAID等）"`
    Platform string `json:"platform" comment:"平台类型：1-TikTok, 2-Google, 3-Apple"`
    PageName string `json:"pageName" comment:"页面名称（Activity/Fragment/Ctr）"`
    AppVersion string `json:"appVersion" comment:"APP版本号"`
    OsType string `json:"osType" comment:"操作系统类型（iOS/Android）"`
    OsVersion string `json:"osVersion" comment:"操作系统版本"`
    DeviceModel string `json:"deviceModel" comment:"设备型号"`
    DeviceBrand string `json:"deviceBrand" comment:"设备品牌"`
    NetworkType string `json:"networkType" comment:"网络类型（WiFi/4G/5G）"`
    CountryCode string `json:"countryCode" comment:"国家代码"`
    Language string `json:"language" comment:"语言"`
    IpAddress string `json:"ipAddress" comment:"IP地址"`
    EventParams string `json:"eventParams" comment:"事件参数（JSON格式，存储自定义参数）"`
    SessionId string `json:"sessionId" comment:"会话ID"`
    Referrer string `json:"referrer" comment:"来源页面"`
    EventTime string `json:"eventTime" comment:"事件发生时间戳（毫秒）"`
    common.ControlBy
}

func (s *MdEventLogTabInsertReq) Generate(model *models.MdEventLogTab)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.EventId = s.EventId
    model.EventCode = s.EventCode
    model.EventName = s.EventName
    model.ModuleName = s.ModuleName
    model.UserId = s.UserId
    model.DeviceId = s.DeviceId
    model.Platform = s.Platform
    model.PageName = s.PageName
    model.AppVersion = s.AppVersion
    model.OsType = s.OsType
    model.OsVersion = s.OsVersion
    model.DeviceModel = s.DeviceModel
    model.DeviceBrand = s.DeviceBrand
    model.NetworkType = s.NetworkType
    model.CountryCode = s.CountryCode
    model.Language = s.Language
    model.IpAddress = s.IpAddress
    model.EventParams = s.EventParams
    model.SessionId = s.SessionId
    model.Referrer = s.Referrer
    model.EventTime = s.EventTime
}

func (s *MdEventLogTabInsertReq) GetId() interface{} {
	return s.Id
}

type MdEventLogTabUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    EventId string `json:"eventId" comment:"事件ID（唯一标识）"`
    EventCode string `json:"eventCode" comment:"事件代码，如：xx_Installation, xx_register_success"`
    EventName string `json:"eventName" comment:"事件名称，如：安装、注册成功"`
    ModuleName string `json:"moduleName" comment:"模块名称"`
    UserId string `json:"userId" comment:"用户ID（未登录为0）"`
    DeviceId string `json:"deviceId" comment:"设备唯一标识（IDFA/GAID等）"`
    Platform string `json:"platform" comment:"平台类型：1-TikTok, 2-Google, 3-Apple"`
    PageName string `json:"pageName" comment:"页面名称（Activity/Fragment/Ctr）"`
    AppVersion string `json:"appVersion" comment:"APP版本号"`
    OsType string `json:"osType" comment:"操作系统类型（iOS/Android）"`
    OsVersion string `json:"osVersion" comment:"操作系统版本"`
    DeviceModel string `json:"deviceModel" comment:"设备型号"`
    DeviceBrand string `json:"deviceBrand" comment:"设备品牌"`
    NetworkType string `json:"networkType" comment:"网络类型（WiFi/4G/5G）"`
    CountryCode string `json:"countryCode" comment:"国家代码"`
    Language string `json:"language" comment:"语言"`
    IpAddress string `json:"ipAddress" comment:"IP地址"`
    EventParams string `json:"eventParams" comment:"事件参数（JSON格式，存储自定义参数）"`
    SessionId string `json:"sessionId" comment:"会话ID"`
    Referrer string `json:"referrer" comment:"来源页面"`
    EventTime string `json:"eventTime" comment:"事件发生时间戳（毫秒）"`
    common.ControlBy
}

func (s *MdEventLogTabUpdateReq) Generate(model *models.MdEventLogTab)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.EventId = s.EventId
    model.EventCode = s.EventCode
    model.EventName = s.EventName
    model.ModuleName = s.ModuleName
    model.UserId = s.UserId
    model.DeviceId = s.DeviceId
    model.Platform = s.Platform
    model.PageName = s.PageName
    model.AppVersion = s.AppVersion
    model.OsType = s.OsType
    model.OsVersion = s.OsVersion
    model.DeviceModel = s.DeviceModel
    model.DeviceBrand = s.DeviceBrand
    model.NetworkType = s.NetworkType
    model.CountryCode = s.CountryCode
    model.Language = s.Language
    model.IpAddress = s.IpAddress
    model.EventParams = s.EventParams
    model.SessionId = s.SessionId
    model.Referrer = s.Referrer
    model.EventTime = s.EventTime
}

func (s *MdEventLogTabUpdateReq) GetId() interface{} {
	return s.Id
}

// MdEventLogTabGetReq 功能获取请求参数
type MdEventLogTabGetReq struct {
     Id int `uri:"id"`
}
func (s *MdEventLogTabGetReq) GetId() interface{} {
	return s.Id
}

// MdEventLogTabDeleteReq 功能删除请求参数
type MdEventLogTabDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *MdEventLogTabDeleteReq) GetId() interface{} {
	return s.Ids
}
