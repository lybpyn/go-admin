package models

import (
    "time"

	"go-admin/common/models"

)

type HsCarouselBanner struct {
    models.Model
    
    Title string `json:"title" gorm:"type:varchar(200);comment:标题"` 
    ImageUrl string `json:"imageUrl" gorm:"type:varchar(500);comment:图片地址"` 
    LinkUrl string `json:"linkUrl" gorm:"type:varchar(500);comment:跳转地址"` 
    SortOrder string `json:"sortOrder" gorm:"type:int(11);comment:排序，小的优先"` 
    StartTime time.Time `json:"startTime" gorm:"type:datetime(3);comment:开始展示时间"` 
    EndTime time.Time `json:"endTime" gorm:"type:datetime(3);comment:结束展示时间"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:0=下线,1=上线"` 
    models.ModelTime
    models.ControlBy
}

func (HsCarouselBanner) TableName() string {
    return "hs_carousel_banner"
}

func (e *HsCarouselBanner) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsCarouselBanner) GetId() interface{} {
	return e.Id
}