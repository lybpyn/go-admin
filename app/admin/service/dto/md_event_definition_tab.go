package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type MdEventDefinitionTabGetPageReq struct {
	dto.Pagination     `search:"-"`
    EventCode string `form:"eventCode"  search:"type:exact;column:event_code;table:md_event_definition_tab" comment:"事件代码（唯一）"`
    MdEventDefinitionTabOrder
}

type MdEventDefinitionTabOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:md_event_definition_tab"`
    EventCode string `form:"eventCodeOrder"  search:"type:order;column:event_code;table:md_event_definition_tab"`
    EventName string `form:"eventNameOrder"  search:"type:order;column:event_name;table:md_event_definition_tab"`
    EventDesc string `form:"eventDescOrder"  search:"type:order;column:event_desc;table:md_event_definition_tab"`
    ModuleName string `form:"moduleNameOrder"  search:"type:order;column:module_name;table:md_event_definition_tab"`
    EventType string `form:"eventTypeOrder"  search:"type:order;column:event_type;table:md_event_definition_tab"`
    ParamSchema string `form:"paramSchemaOrder"  search:"type:order;column:param_schema;table:md_event_definition_tab"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:md_event_definition_tab"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:md_event_definition_tab"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:md_event_definition_tab"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:md_event_definition_tab"`
    
}

func (m *MdEventDefinitionTabGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type MdEventDefinitionTabInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    EventCode string `json:"eventCode" comment:"事件代码（唯一）"`
    EventName string `json:"eventName" comment:"事件名称"`
    EventDesc string `json:"eventDesc" comment:"事件描述"`
    ModuleName string `json:"moduleName" comment:"所属模块"`
    EventType string `json:"eventType" comment:"事件类型：1-系统事件, 2-业务事件"`
    ParamSchema string `json:"paramSchema" comment:"参数定义（JSON Schema格式）"`
    Status string `json:"status" comment:"状态：0-禁用, 1-启用"`
    common.ControlBy
}

func (s *MdEventDefinitionTabInsertReq) Generate(model *models.MdEventDefinitionTab)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.EventCode = s.EventCode
    model.EventName = s.EventName
    model.EventDesc = s.EventDesc
    model.ModuleName = s.ModuleName
    model.EventType = s.EventType
    model.ParamSchema = s.ParamSchema
    model.Status = s.Status
}

func (s *MdEventDefinitionTabInsertReq) GetId() interface{} {
	return s.Id
}

type MdEventDefinitionTabUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    EventCode string `json:"eventCode" comment:"事件代码（唯一）"`
    EventName string `json:"eventName" comment:"事件名称"`
    EventDesc string `json:"eventDesc" comment:"事件描述"`
    ModuleName string `json:"moduleName" comment:"所属模块"`
    EventType string `json:"eventType" comment:"事件类型：1-系统事件, 2-业务事件"`
    ParamSchema string `json:"paramSchema" comment:"参数定义（JSON Schema格式）"`
    Status string `json:"status" comment:"状态：0-禁用, 1-启用"`
    common.ControlBy
}

func (s *MdEventDefinitionTabUpdateReq) Generate(model *models.MdEventDefinitionTab)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.EventCode = s.EventCode
    model.EventName = s.EventName
    model.EventDesc = s.EventDesc
    model.ModuleName = s.ModuleName
    model.EventType = s.EventType
    model.ParamSchema = s.ParamSchema
    model.Status = s.Status
}

func (s *MdEventDefinitionTabUpdateReq) GetId() interface{} {
	return s.Id
}

// MdEventDefinitionTabGetReq 功能获取请求参数
type MdEventDefinitionTabGetReq struct {
     Id int `uri:"id"`
}
func (s *MdEventDefinitionTabGetReq) GetId() interface{} {
	return s.Id
}

// MdEventDefinitionTabDeleteReq 功能删除请求参数
type MdEventDefinitionTabDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *MdEventDefinitionTabDeleteReq) GetId() interface{} {
	return s.Ids
}
