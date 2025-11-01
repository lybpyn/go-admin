package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type MdUserDeviceTabGetPageReq struct {
	dto.Pagination     `search:"-"`
    DeviceId string `form:"deviceId"  search:"type:exact;column:device_id;table:md_user_device_tab" comment:"设备唯一标识"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:md_user_device_tab" comment:"用户ID（未绑定为0）"`
    MdUserDeviceTabOrder
}

type MdUserDeviceTabOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:md_user_device_tab"`
    DeviceId string `form:"deviceIdOrder"  search:"type:order;column:device_id;table:md_user_device_tab"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:md_user_device_tab"`
    Platform string `form:"platformOrder"  search:"type:order;column:platform;table:md_user_device_tab"`
    OsType string `form:"osTypeOrder"  search:"type:order;column:os_type;table:md_user_device_tab"`
    OsVersion string `form:"osVersionOrder"  search:"type:order;column:os_version;table:md_user_device_tab"`
    DeviceModel string `form:"deviceModelOrder"  search:"type:order;column:device_model;table:md_user_device_tab"`
    DeviceBrand string `form:"deviceBrandOrder"  search:"type:order;column:device_brand;table:md_user_device_tab"`
    AppVersion string `form:"appVersionOrder"  search:"type:order;column:app_version;table:md_user_device_tab"`
    CountryCode string `form:"countryCodeOrder"  search:"type:order;column:country_code;table:md_user_device_tab"`
    Language string `form:"languageOrder"  search:"type:order;column:language;table:md_user_device_tab"`
    InstallTime string `form:"installTimeOrder"  search:"type:order;column:install_time;table:md_user_device_tab"`
    FirstOpenTime string `form:"firstOpenTimeOrder"  search:"type:order;column:first_open_time;table:md_user_device_tab"`
    LastActiveTime string `form:"lastActiveTimeOrder"  search:"type:order;column:last_active_time;table:md_user_device_tab"`
    IsRegistered string `form:"isRegisteredOrder"  search:"type:order;column:is_registered;table:md_user_device_tab"`
    RegisterTime string `form:"registerTimeOrder"  search:"type:order;column:register_time;table:md_user_device_tab"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:md_user_device_tab"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:md_user_device_tab"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:md_user_device_tab"`
    
}

func (m *MdUserDeviceTabGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type MdUserDeviceTabInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    DeviceId string `json:"deviceId" comment:"设备唯一标识"`
    UserId string `json:"userId" comment:"用户ID（未绑定为0）"`
    Platform string `json:"platform" comment:"平台类型：1-TikTok, 2-Google, 3-Apple"`
    OsType string `json:"osType" comment:"操作系统类型（iOS/Android）"`
    OsVersion string `json:"osVersion" comment:"操作系统版本"`
    DeviceModel string `json:"deviceModel" comment:"设备型号"`
    DeviceBrand string `json:"deviceBrand" comment:"设备品牌"`
    AppVersion string `json:"appVersion" comment:"APP版本号"`
    CountryCode string `json:"countryCode" comment:"国家代码"`
    Language string `json:"language" comment:"语言"`
    InstallTime time.Time `json:"installTime" comment:"首次安装时间"`
    FirstOpenTime time.Time `json:"firstOpenTime" comment:"首次打开时间"`
    LastActiveTime time.Time `json:"lastActiveTime" comment:"最后活跃时间"`
    IsRegistered string `json:"isRegistered" comment:"是否已注册：0-否, 1-是"`
    RegisterTime time.Time `json:"registerTime" comment:"注册时间"`
    common.ControlBy
}

func (s *MdUserDeviceTabInsertReq) Generate(model *models.MdUserDeviceTab)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.DeviceId = s.DeviceId
    model.UserId = s.UserId
    model.Platform = s.Platform
    model.OsType = s.OsType
    model.OsVersion = s.OsVersion
    model.DeviceModel = s.DeviceModel
    model.DeviceBrand = s.DeviceBrand
    model.AppVersion = s.AppVersion
    model.CountryCode = s.CountryCode
    model.Language = s.Language
    model.InstallTime = s.InstallTime
    model.FirstOpenTime = s.FirstOpenTime
    model.LastActiveTime = s.LastActiveTime
    model.IsRegistered = s.IsRegistered
    model.RegisterTime = s.RegisterTime
}

func (s *MdUserDeviceTabInsertReq) GetId() interface{} {
	return s.Id
}

type MdUserDeviceTabUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    DeviceId string `json:"deviceId" comment:"设备唯一标识"`
    UserId string `json:"userId" comment:"用户ID（未绑定为0）"`
    Platform string `json:"platform" comment:"平台类型：1-TikTok, 2-Google, 3-Apple"`
    OsType string `json:"osType" comment:"操作系统类型（iOS/Android）"`
    OsVersion string `json:"osVersion" comment:"操作系统版本"`
    DeviceModel string `json:"deviceModel" comment:"设备型号"`
    DeviceBrand string `json:"deviceBrand" comment:"设备品牌"`
    AppVersion string `json:"appVersion" comment:"APP版本号"`
    CountryCode string `json:"countryCode" comment:"国家代码"`
    Language string `json:"language" comment:"语言"`
    InstallTime time.Time `json:"installTime" comment:"首次安装时间"`
    FirstOpenTime time.Time `json:"firstOpenTime" comment:"首次打开时间"`
    LastActiveTime time.Time `json:"lastActiveTime" comment:"最后活跃时间"`
    IsRegistered string `json:"isRegistered" comment:"是否已注册：0-否, 1-是"`
    RegisterTime time.Time `json:"registerTime" comment:"注册时间"`
    common.ControlBy
}

func (s *MdUserDeviceTabUpdateReq) Generate(model *models.MdUserDeviceTab)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.DeviceId = s.DeviceId
    model.UserId = s.UserId
    model.Platform = s.Platform
    model.OsType = s.OsType
    model.OsVersion = s.OsVersion
    model.DeviceModel = s.DeviceModel
    model.DeviceBrand = s.DeviceBrand
    model.AppVersion = s.AppVersion
    model.CountryCode = s.CountryCode
    model.Language = s.Language
    model.InstallTime = s.InstallTime
    model.FirstOpenTime = s.FirstOpenTime
    model.LastActiveTime = s.LastActiveTime
    model.IsRegistered = s.IsRegistered
    model.RegisterTime = s.RegisterTime
}

func (s *MdUserDeviceTabUpdateReq) GetId() interface{} {
	return s.Id
}

// MdUserDeviceTabGetReq 功能获取请求参数
type MdUserDeviceTabGetReq struct {
     Id int `uri:"id"`
}
func (s *MdUserDeviceTabGetReq) GetId() interface{} {
	return s.Id
}

// MdUserDeviceTabDeleteReq 功能删除请求参数
type MdUserDeviceTabDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *MdUserDeviceTabDeleteReq) GetId() interface{} {
	return s.Ids
}
