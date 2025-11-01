package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsUserIdentityGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_user_identity" comment:""`
    IdentityType string `form:"identityType"  search:"type:exact;column:identity_type;table:hs_user_identity" comment:"凭证类型: email, phone, google, apple, twitter, tiktok"`
    Identifier string `form:"identifier"  search:"type:exact;column:identifier;table:hs_user_identity" comment:"凭证值: 邮箱, 手机号, 第三方平台的唯一ID"`
    Credential string `form:"credential"  search:"type:exact;column:credential;table:hs_user_identity" comment:"存储密码哈希(email/phone)或refresh_token(social)"`
    HsUserIdentityOrder
}

type HsUserIdentityOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_user_identity"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_user_identity"`
    IdentityType string `form:"identityTypeOrder"  search:"type:order;column:identity_type;table:hs_user_identity"`
    Identifier string `form:"identifierOrder"  search:"type:order;column:identifier;table:hs_user_identity"`
    Credential string `form:"credentialOrder"  search:"type:order;column:credential;table:hs_user_identity"`
    Verified string `form:"verifiedOrder"  search:"type:order;column:verified;table:hs_user_identity"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_user_identity"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_user_identity"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_user_identity"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_user_identity"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_user_identity"`
    
}

func (m *HsUserIdentityGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsUserIdentityInsertReq struct {
    Id int `json:"-" comment:""` // 
    UserId string `json:"userId" comment:""`
    IdentityType string `json:"identityType" comment:"凭证类型: email, phone, google, apple, twitter, tiktok"`
    Identifier string `json:"identifier" comment:"凭证值: 邮箱, 手机号, 第三方平台的唯一ID"`
    Credential string `json:"credential" comment:"存储密码哈希(email/phone)或refresh_token(social)"`
    Verified string `json:"verified" comment:"是否已验证（0:未验证, 1:已验证）"`
    common.ControlBy
}

func (s *HsUserIdentityInsertReq) Generate(model *models.HsUserIdentity)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.IdentityType = s.IdentityType
    model.Identifier = s.Identifier
    model.Credential = s.Credential
    model.Verified = s.Verified
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsUserIdentityInsertReq) GetId() interface{} {
	return s.Id
}

type HsUserIdentityUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    UserId string `json:"userId" comment:""`
    IdentityType string `json:"identityType" comment:"凭证类型: email, phone, google, apple, twitter, tiktok"`
    Identifier string `json:"identifier" comment:"凭证值: 邮箱, 手机号, 第三方平台的唯一ID"`
    Credential string `json:"credential" comment:"存储密码哈希(email/phone)或refresh_token(social)"`
    Verified string `json:"verified" comment:"是否已验证（0:未验证, 1:已验证）"`
    common.ControlBy
}

func (s *HsUserIdentityUpdateReq) Generate(model *models.HsUserIdentity)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.IdentityType = s.IdentityType
    model.Identifier = s.Identifier
    model.Credential = s.Credential
    model.Verified = s.Verified
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsUserIdentityUpdateReq) GetId() interface{} {
	return s.Id
}

// HsUserIdentityGetReq 功能获取请求参数
type HsUserIdentityGetReq struct {
     Id int `uri:"id"`
}
func (s *HsUserIdentityGetReq) GetId() interface{} {
	return s.Id
}

// HsUserIdentityDeleteReq 功能删除请求参数
type HsUserIdentityDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsUserIdentityDeleteReq) GetId() interface{} {
	return s.Ids
}
