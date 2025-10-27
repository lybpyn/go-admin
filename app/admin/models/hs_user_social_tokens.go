package models

import (
    "time"

	"go-admin/common/models"

)

type HsUserSocialTokens struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20) unsigned;comment:UserId"` 
    Provider string `json:"provider" gorm:"type:varchar(30);comment:google, apple, twitter, etc."` 
    AccessToken string `json:"accessToken" gorm:"type:varchar(1024);comment:AccessToken"` 
    RefreshToken string `json:"refreshToken" gorm:"type:varchar(1024);comment:RefreshToken"` 
    ExpiresIn string `json:"expiresIn" gorm:"type:int(11);comment:有效期（秒）"` 
    TokenType string `json:"tokenType" gorm:"type:varchar(50);comment:TokenType"` 
    Scope string `json:"scope" gorm:"type:varchar(255);comment:Scope"` 
    IssuedAt time.Time `json:"issuedAt" gorm:"type:datetime(3);comment:IssuedAt"` 
    models.ModelTime
    models.ControlBy
}

func (HsUserSocialTokens) TableName() string {
    return "hs_user_social_tokens"
}

func (e *HsUserSocialTokens) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsUserSocialTokens) GetId() interface{} {
	return e.Id
}