package models

import (

	"go-admin/common/models"

)

type HsUserIdentity struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20) unsigned;comment:UserId"` 
    IdentityType string `json:"identityType" gorm:"type:varchar(30);comment:凭证类型: email, phone, google, apple, twitter, tiktok"` 
    Identifier string `json:"identifier" gorm:"type:varchar(255);comment:凭证值: 邮箱, 手机号, 第三方平台的唯一ID"` 
    Credential string `json:"credential" gorm:"type:varchar(512);comment:存储密码哈希(email/phone)或refresh_token(social)"` 
    Verified string `json:"verified" gorm:"type:tinyint(1);comment:是否已验证（0:未验证, 1:已验证）"` 
    models.ModelTime
    models.ControlBy
}

func (HsUserIdentity) TableName() string {
    return "hs_user_identity"
}

func (e *HsUserIdentity) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsUserIdentity) GetId() interface{} {
	return e.Id
}