package models

import (

	"go-admin/common/models"

)

type HsUserBankCards struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20) unsigned;comment:用户ID"` 
    BankId string `json:"bankId" gorm:"type:bigint(20);comment:银行ID（关联hs_banks表）"` 
    BankName string `json:"bankName" gorm:"type:varchar(100);comment:银行名称"` 
    CardNumber string `json:"cardNumber" gorm:"type:varchar(32);comment:银行卡号（加密存储）"` 
    CardHolderName string `json:"cardHolderName" gorm:"type:varchar(100);comment:持卡人姓名"` 
    CardType string `json:"cardType" gorm:"type:tinyint(1);comment:卡片类型：1=储蓄卡，2=信用卡"` 
    BranchName string `json:"branchName" gorm:"type:varchar(200);comment:开户行/支行名称"` 
    Phone string `json:"phone" gorm:"type:varchar(20);comment:预留手机号"` 
    IdCard string `json:"idCard" gorm:"type:varchar(32);comment:身份证号（加密存储）"` 
    IsDefault string `json:"isDefault" gorm:"type:tinyint(1);comment:是否默认银行卡：0=否，1=是"` 
    Status string `json:"status" gorm:"type:tinyint(1);comment:状态：0=禁用，1=启用，2=审核中"` 
    Remark string `json:"remark" gorm:"type:varchar(500);comment:备注"` 
    models.ModelTime
    models.ControlBy
}

func (HsUserBankCards) TableName() string {
    return "hs_user_bank_cards"
}

func (e *HsUserBankCards) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsUserBankCards) GetId() interface{} {
	return e.Id
}