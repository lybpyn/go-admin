package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsUserSocialTokensGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_user_social_tokens" comment:""`
    HsUserSocialTokensOrder
}

type HsUserSocialTokensOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_user_social_tokens"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_user_social_tokens"`
    Provider string `form:"providerOrder"  search:"type:order;column:provider;table:hs_user_social_tokens"`
    AccessToken string `form:"accessTokenOrder"  search:"type:order;column:access_token;table:hs_user_social_tokens"`
    RefreshToken string `form:"refreshTokenOrder"  search:"type:order;column:refresh_token;table:hs_user_social_tokens"`
    ExpiresIn string `form:"expiresInOrder"  search:"type:order;column:expires_in;table:hs_user_social_tokens"`
    TokenType string `form:"tokenTypeOrder"  search:"type:order;column:token_type;table:hs_user_social_tokens"`
    Scope string `form:"scopeOrder"  search:"type:order;column:scope;table:hs_user_social_tokens"`
    IssuedAt string `form:"issuedAtOrder"  search:"type:order;column:issued_at;table:hs_user_social_tokens"`
    
}

func (m *HsUserSocialTokensGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsUserSocialTokensInsertReq struct {
    Id int `json:"-" comment:""` // 
    UserId string `json:"userId" comment:""`
    Provider string `json:"provider" comment:"google, apple, twitter, etc."`
    AccessToken string `json:"accessToken" comment:""`
    RefreshToken string `json:"refreshToken" comment:""`
    ExpiresIn string `json:"expiresIn" comment:"有效期（秒）"`
    TokenType string `json:"tokenType" comment:""`
    Scope string `json:"scope" comment:""`
    IssuedAt time.Time `json:"issuedAt" comment:""`
    common.ControlBy
}

func (s *HsUserSocialTokensInsertReq) Generate(model *models.HsUserSocialTokens)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.Provider = s.Provider
    model.AccessToken = s.AccessToken
    model.RefreshToken = s.RefreshToken
    model.ExpiresIn = s.ExpiresIn
    model.TokenType = s.TokenType
    model.Scope = s.Scope
    model.IssuedAt = s.IssuedAt
}

func (s *HsUserSocialTokensInsertReq) GetId() interface{} {
	return s.Id
}

type HsUserSocialTokensUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    UserId string `json:"userId" comment:""`
    Provider string `json:"provider" comment:"google, apple, twitter, etc."`
    AccessToken string `json:"accessToken" comment:""`
    RefreshToken string `json:"refreshToken" comment:""`
    ExpiresIn string `json:"expiresIn" comment:"有效期（秒）"`
    TokenType string `json:"tokenType" comment:""`
    Scope string `json:"scope" comment:""`
    IssuedAt time.Time `json:"issuedAt" comment:""`
    common.ControlBy
}

func (s *HsUserSocialTokensUpdateReq) Generate(model *models.HsUserSocialTokens)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.Provider = s.Provider
    model.AccessToken = s.AccessToken
    model.RefreshToken = s.RefreshToken
    model.ExpiresIn = s.ExpiresIn
    model.TokenType = s.TokenType
    model.Scope = s.Scope
    model.IssuedAt = s.IssuedAt
}

func (s *HsUserSocialTokensUpdateReq) GetId() interface{} {
	return s.Id
}

// HsUserSocialTokensGetReq 功能获取请求参数
type HsUserSocialTokensGetReq struct {
     Id int `uri:"id"`
}
func (s *HsUserSocialTokensGetReq) GetId() interface{} {
	return s.Id
}

// HsUserSocialTokensDeleteReq 功能删除请求参数
type HsUserSocialTokensDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsUserSocialTokensDeleteReq) GetId() interface{} {
	return s.Ids
}
