package models

import (

	"go-admin/common/models"

)

type HsMerchants struct {
    models.Model
    
    MerchantCode string `json:"merchantCode" gorm:"type:varchar(64);comment:外部/内部唯一编码 (可用于对接第三方)"` 
    Name string `json:"name" gorm:"type:varchar(128);comment:卡商名称/公司名"` 
    ContactName string `json:"contactName" gorm:"type:varchar(64);comment:联系人姓名"` 
    ContactPhone string `json:"contactPhone" gorm:"type:varchar(32);comment:联系人电话"` 
    ContactEmail string `json:"contactEmail" gorm:"type:varchar(128);comment:联系人邮箱"` 
    Country string `json:"country" gorm:"type:char(2);comment:国家/地区 ISO2 (如 CN, US)"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:状态: 0=禁用,1=启用,2=冻结"` 
    DailyLimit string `json:"dailyLimit" gorm:"type:decimal(20,4);comment:日限额 (可选)"` 
    Note string `json:"note" gorm:"type:text;comment:备注/其他说明"` 
    Extra string `json:"extra" gorm:"type:json;comment:扩展信息: 如资质文件url、合同信息等"` 
    models.ModelTime
    models.ControlBy
}

func (HsMerchants) TableName() string {
    return "hs_merchants"
}

func (e *HsMerchants) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsMerchants) GetId() interface{} {
	return e.Id
}