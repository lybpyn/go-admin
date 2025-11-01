package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdGiftcardGetPageReq struct {
	dto.Pagination     `search:"-"`
    OrdGiftcardOrder
}

type OrdGiftcardOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ord_giftcard"`
    RegionId string `form:"regionIdOrder"  search:"type:order;column:region_id;table:ord_giftcard"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:ord_giftcard"`
    ValuesConfig string `form:"valuesConfigOrder"  search:"type:order;column:values_config;table:ord_giftcard"`
    DiscountRate string `form:"discountRateOrder"  search:"type:order;column:discount_rate;table:ord_giftcard"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:ord_giftcard"`
    Extra string `form:"extraOrder"  search:"type:order;column:extra;table:ord_giftcard"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ord_giftcard"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ord_giftcard"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_giftcard"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ord_giftcard"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ord_giftcard"`
    
}

func (m *OrdGiftcardGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdGiftcardInsertReq struct {
    Id int `json:"-" comment:"礼品卡ID"` // 礼品卡ID
    RegionId string `json:"regionId" comment:"地区ID"`
    Name string `json:"name" comment:"卡名称，例如 Steam 50 USD"`
    ValuesConfig string `json:"valuesConfig" comment:"面额配置，支持区间和固定值"`
    DiscountRate string `json:"discountRate" comment:"折扣汇率，例如0.95 表示95折"`
    Status string `json:"status" comment:"状态: 1=启用, 0=禁用"`
    Extra string `json:"extra" comment:"扩展信息，如购买说明、限制条件"`
    common.ControlBy
}

func (s *OrdGiftcardInsertReq) Generate(model *models.OrdGiftcard)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.RegionId = s.RegionId
    model.Name = s.Name
    model.ValuesConfig = s.ValuesConfig
    model.DiscountRate = s.DiscountRate
    model.Status = s.Status
    model.Extra = s.Extra
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *OrdGiftcardInsertReq) GetId() interface{} {
	return s.Id
}

type OrdGiftcardUpdateReq struct {
    Id int `uri:"id" comment:"礼品卡ID"` // 礼品卡ID
    RegionId string `json:"regionId" comment:"地区ID"`
    Name string `json:"name" comment:"卡名称，例如 Steam 50 USD"`
    ValuesConfig string `json:"valuesConfig" comment:"面额配置，支持区间和固定值"`
    DiscountRate string `json:"discountRate" comment:"折扣汇率，例如0.95 表示95折"`
    Status string `json:"status" comment:"状态: 1=启用, 0=禁用"`
    Extra string `json:"extra" comment:"扩展信息，如购买说明、限制条件"`
    common.ControlBy
}

func (s *OrdGiftcardUpdateReq) Generate(model *models.OrdGiftcard)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.RegionId = s.RegionId
    model.Name = s.Name
    model.ValuesConfig = s.ValuesConfig
    model.DiscountRate = s.DiscountRate
    model.Status = s.Status
    model.Extra = s.Extra
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *OrdGiftcardUpdateReq) GetId() interface{} {
	return s.Id
}

// OrdGiftcardGetReq 功能获取请求参数
type OrdGiftcardGetReq struct {
     Id int `uri:"id"`
}
func (s *OrdGiftcardGetReq) GetId() interface{} {
	return s.Id
}

// OrdGiftcardDeleteReq 功能删除请求参数
type OrdGiftcardDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrdGiftcardDeleteReq) GetId() interface{} {
	return s.Ids
}
