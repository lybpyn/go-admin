package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsCarouselBannerGetPageReq struct {
	dto.Pagination     `search:"-"`
    StartTime time.Time `form:"startTime"  search:"type:exact;column:start_time;table:hs_carousel_banner" comment:"开始展示时间"`
    EndTime time.Time `form:"endTime"  search:"type:exact;column:end_time;table:hs_carousel_banner" comment:"结束展示时间"`
    HsCarouselBannerOrder
}

type HsCarouselBannerOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_carousel_banner"`
    Title string `form:"titleOrder"  search:"type:order;column:title;table:hs_carousel_banner"`
    ImageUrl string `form:"imageUrlOrder"  search:"type:order;column:image_url;table:hs_carousel_banner"`
    LinkUrl string `form:"linkUrlOrder"  search:"type:order;column:link_url;table:hs_carousel_banner"`
    SortOrder string `form:"sortOrderOrder"  search:"type:order;column:sort_order;table:hs_carousel_banner"`
    StartTime string `form:"startTimeOrder"  search:"type:order;column:start_time;table:hs_carousel_banner"`
    EndTime string `form:"endTimeOrder"  search:"type:order;column:end_time;table:hs_carousel_banner"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_carousel_banner"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_carousel_banner"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_carousel_banner"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_carousel_banner"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_carousel_banner"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_carousel_banner"`
    
}

func (m *HsCarouselBannerGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsCarouselBannerInsertReq struct {
    Id int `json:"-" comment:""` // 
    Title string `json:"title" comment:"标题"`
    ImageUrl string `json:"imageUrl" comment:"图片地址"`
    LinkUrl string `json:"linkUrl" comment:"跳转地址"`
    SortOrder string `json:"sortOrder" comment:"排序，小的优先"`
    StartTime time.Time `json:"startTime" comment:"开始展示时间"`
    EndTime time.Time `json:"endTime" comment:"结束展示时间"`
    Status string `json:"status" comment:"0=下线,1=上线"`
    common.ControlBy
}

func (s *HsCarouselBannerInsertReq) Generate(model *models.HsCarouselBanner)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Title = s.Title
    model.ImageUrl = s.ImageUrl
    model.LinkUrl = s.LinkUrl
    model.SortOrder = s.SortOrder
    model.StartTime = s.StartTime
    model.EndTime = s.EndTime
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsCarouselBannerInsertReq) GetId() interface{} {
	return s.Id
}

type HsCarouselBannerUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Title string `json:"title" comment:"标题"`
    ImageUrl string `json:"imageUrl" comment:"图片地址"`
    LinkUrl string `json:"linkUrl" comment:"跳转地址"`
    SortOrder string `json:"sortOrder" comment:"排序，小的优先"`
    StartTime time.Time `json:"startTime" comment:"开始展示时间"`
    EndTime time.Time `json:"endTime" comment:"结束展示时间"`
    Status string `json:"status" comment:"0=下线,1=上线"`
    common.ControlBy
}

func (s *HsCarouselBannerUpdateReq) Generate(model *models.HsCarouselBanner)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Title = s.Title
    model.ImageUrl = s.ImageUrl
    model.LinkUrl = s.LinkUrl
    model.SortOrder = s.SortOrder
    model.StartTime = s.StartTime
    model.EndTime = s.EndTime
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsCarouselBannerUpdateReq) GetId() interface{} {
	return s.Id
}

// HsCarouselBannerGetReq 功能获取请求参数
type HsCarouselBannerGetReq struct {
     Id int `uri:"id"`
}
func (s *HsCarouselBannerGetReq) GetId() interface{} {
	return s.Id
}

// HsCarouselBannerDeleteReq 功能删除请求参数
type HsCarouselBannerDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsCarouselBannerDeleteReq) GetId() interface{} {
	return s.Ids
}
