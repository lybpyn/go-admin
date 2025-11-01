package models

import (

	"go-admin/common/models"

)

type MdEventDefinitionTab struct {
    models.Model
    
    EventCode string `json:"eventCode" gorm:"type:varchar(64);comment:事件代码（唯一）"` 
    EventName string `json:"eventName" gorm:"type:varchar(128);comment:事件名称"` 
    EventDesc string `json:"eventDesc" gorm:"type:varchar(512);comment:事件描述"` 
    ModuleName string `json:"moduleName" gorm:"type:varchar(64);comment:所属模块"` 
    EventType string `json:"eventType" gorm:"type:tinyint(4);comment:事件类型：1-系统事件, 2-业务事件"` 
    ParamSchema string `json:"paramSchema" gorm:"type:text;comment:参数定义（JSON Schema格式）"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:状态：0-禁用, 1-启用"` 
    models.ModelTime
    models.ControlBy
}

func (MdEventDefinitionTab) TableName() string {
    return "md_event_definition_tab"
}

func (e *MdEventDefinitionTab) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *MdEventDefinitionTab) GetId() interface{} {
	return e.Id
}