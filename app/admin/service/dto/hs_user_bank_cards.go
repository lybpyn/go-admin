package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsUserBankCardsGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_user_bank_cards" comment:"用户ID"`
    BankId string `form:"bankId"  search:"type:exact;column:bank_id;table:hs_user_bank_cards" comment:"银行ID（关联hs_banks表）"`
    CardNumber string `form:"cardNumber"  search:"type:exact;column:card_number;table:hs_user_bank_cards" comment:"银行卡号（加密存储）"`
    CardHolderName string `form:"cardHolderName"  search:"type:exact;column:card_holder_name;table:hs_user_bank_cards" comment:"持卡人姓名"`
    HsUserBankCardsOrder
}

type HsUserBankCardsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_user_bank_cards"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_user_bank_cards"`
    BankId string `form:"bankIdOrder"  search:"type:order;column:bank_id;table:hs_user_bank_cards"`
    BankName string `form:"bankNameOrder"  search:"type:order;column:bank_name;table:hs_user_bank_cards"`
    CardNumber string `form:"cardNumberOrder"  search:"type:order;column:card_number;table:hs_user_bank_cards"`
    CardHolderName string `form:"cardHolderNameOrder"  search:"type:order;column:card_holder_name;table:hs_user_bank_cards"`
    CardType string `form:"cardTypeOrder"  search:"type:order;column:card_type;table:hs_user_bank_cards"`
    BranchName string `form:"branchNameOrder"  search:"type:order;column:branch_name;table:hs_user_bank_cards"`
    Phone string `form:"phoneOrder"  search:"type:order;column:phone;table:hs_user_bank_cards"`
    IdCard string `form:"idCardOrder"  search:"type:order;column:id_card;table:hs_user_bank_cards"`
    IsDefault string `form:"isDefaultOrder"  search:"type:order;column:is_default;table:hs_user_bank_cards"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_user_bank_cards"`
    Remark string `form:"remarkOrder"  search:"type:order;column:remark;table:hs_user_bank_cards"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_user_bank_cards"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_user_bank_cards"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_user_bank_cards"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_user_bank_cards"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_user_bank_cards"`
    
}

func (m *HsUserBankCardsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsUserBankCardsInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    UserId string `json:"userId" comment:"用户ID"`
    BankId string `json:"bankId" comment:"银行ID（关联hs_banks表）"`
    BankName string `json:"bankName" comment:"银行名称"`
    CardNumber string `json:"cardNumber" comment:"银行卡号（加密存储）"`
    CardHolderName string `json:"cardHolderName" comment:"持卡人姓名"`
    CardType string `json:"cardType" comment:"卡片类型：1=储蓄卡，2=信用卡"`
    BranchName string `json:"branchName" comment:"开户行/支行名称"`
    Phone string `json:"phone" comment:"预留手机号"`
    IdCard string `json:"idCard" comment:"身份证号（加密存储）"`
    IsDefault string `json:"isDefault" comment:"是否默认银行卡：0=否，1=是"`
    Status string `json:"status" comment:"状态：0=禁用，1=启用，2=审核中"`
    Remark string `json:"remark" comment:"备注"`
    common.ControlBy
}

func (s *HsUserBankCardsInsertReq) Generate(model *models.HsUserBankCards)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.BankId = s.BankId
    model.BankName = s.BankName
    model.CardNumber = s.CardNumber
    model.CardHolderName = s.CardHolderName
    model.CardType = s.CardType
    model.BranchName = s.BranchName
    model.Phone = s.Phone
    model.IdCard = s.IdCard
    model.IsDefault = s.IsDefault
    model.Status = s.Status
    model.Remark = s.Remark
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsUserBankCardsInsertReq) GetId() interface{} {
	return s.Id
}

type HsUserBankCardsUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    UserId string `json:"userId" comment:"用户ID"`
    BankId string `json:"bankId" comment:"银行ID（关联hs_banks表）"`
    BankName string `json:"bankName" comment:"银行名称"`
    CardNumber string `json:"cardNumber" comment:"银行卡号（加密存储）"`
    CardHolderName string `json:"cardHolderName" comment:"持卡人姓名"`
    CardType string `json:"cardType" comment:"卡片类型：1=储蓄卡，2=信用卡"`
    BranchName string `json:"branchName" comment:"开户行/支行名称"`
    Phone string `json:"phone" comment:"预留手机号"`
    IdCard string `json:"idCard" comment:"身份证号（加密存储）"`
    IsDefault string `json:"isDefault" comment:"是否默认银行卡：0=否，1=是"`
    Status string `json:"status" comment:"状态：0=禁用，1=启用，2=审核中"`
    Remark string `json:"remark" comment:"备注"`
    common.ControlBy
}

func (s *HsUserBankCardsUpdateReq) Generate(model *models.HsUserBankCards)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.BankId = s.BankId
    model.BankName = s.BankName
    model.CardNumber = s.CardNumber
    model.CardHolderName = s.CardHolderName
    model.CardType = s.CardType
    model.BranchName = s.BranchName
    model.Phone = s.Phone
    model.IdCard = s.IdCard
    model.IsDefault = s.IsDefault
    model.Status = s.Status
    model.Remark = s.Remark
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsUserBankCardsUpdateReq) GetId() interface{} {
	return s.Id
}

// HsUserBankCardsGetReq 功能获取请求参数
type HsUserBankCardsGetReq struct {
     Id int `uri:"id"`
}
func (s *HsUserBankCardsGetReq) GetId() interface{} {
	return s.Id
}

// HsUserBankCardsDeleteReq 功能删除请求参数
type HsUserBankCardsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsUserBankCardsDeleteReq) GetId() interface{} {
	return s.Ids
}
