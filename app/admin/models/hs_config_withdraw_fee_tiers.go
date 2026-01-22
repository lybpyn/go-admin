package models

import (
	"go-admin/common/models"
)

// HsConfigWithdrawFeeTiers 提现阶梯收费配置表
type HsConfigWithdrawFeeTiers struct {
	models.Model

	RuleId    int    `json:"ruleId" gorm:"type:bigint(20) unsigned;not null;comment:关联的提现规则ID"`
	MinAmount string `json:"minAmount" gorm:"type:decimal(36,2);not null;default:0.00;comment:区间最小金额（包含）"`
	MaxAmount string `json:"maxAmount" gorm:"type:decimal(36,2);not null;default:0.00;comment:区间最大金额（包含），0表示无上限"`
	FeeAmount string `json:"feeAmount" gorm:"type:decimal(36,2);not null;default:0.00;comment:该区间的手续费金额"`
	SortOrder int    `json:"sortOrder" gorm:"type:int(11);not null;default:0;comment:排序顺序，值越小越优先"`
	models.ModelTime
}

func (HsConfigWithdrawFeeTiers) TableName() string {
	return "hs_config_withdraw_fee_tiers"
}

func (e *HsConfigWithdrawFeeTiers) GetId() interface{} {
	return e.Id
}
