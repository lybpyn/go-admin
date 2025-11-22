package models

import (
	"go-admin/common/models"
)

type HsUsers struct {
	models.Model

	Username            string `json:"username" gorm:"type:varchar(50);comment:用户名（可选展示用）"`
	PasswordHash        string `json:"passwordHash" gorm:"type:varchar(255);comment:密码"`
	Firstname           string `json:"firstname" gorm:"type:varchar(20);comment:姓"`
	Lastname            string `json:"lastname" gorm:"type:varchar(20);comment:名"`
	Avatar              string `json:"avatar" gorm:"type:varchar(255);comment:头像URL"`
	Balance             string `json:"balance" gorm:"type:decimal(20,2);comment:法币可用余额"`
	CryptoBalance       string `json:"cryptoBalance" gorm:"type:decimal(20,8);comment:虚拟币可用余额"`
	FrozenBalance       string `json:"frozenBalance" gorm:"type:decimal(20,2);comment:法币冻结余额"`
	CryptoFrozenBalance string `json:"cryptoFrozenBalance" gorm:"type:decimal(20,8);comment:冻结虚拟币余额(USDT计价)"`
	LevelId             string `json:"levelId" gorm:"type:int(11);comment:用户等级ID"`
	Experience          string `json:"experience" gorm:"type:int(11);comment:当前经验"`
	RegionId            string `json:"regionId" gorm:"type:bigint(20);comment:区域id"`
	TotalExperience     string `json:"totalExperience" gorm:"type:int(11);comment:累计经验"`
	InviteCode   string `json:"inviteCode" gorm:"type:varchar(8);comment:邀请码"`
	Status       string `json:"status" gorm:"type:tinyint(4);comment:状态：1正常，0封禁"`
	Version      string `json:"version" gorm:"type:bigint(20);comment:Version"`
	CurrencyCode string `json:"currencyCode" gorm:"type:char(3);comment:用户所属地区货币，如 NGN/KES/USD"`
	models.ModelTime
	models.ControlBy
}

func (HsUsers) TableName() string {
	return "hs_users"
}

func (e *HsUsers) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsUsers) GetId() interface{} {
	return e.Id
}
