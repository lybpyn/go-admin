package models

import (

	"go-admin/common/models"

)

type OrdRankLeaderboard struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20);comment:UserId"` 
    Avatar string `json:"avatar" gorm:"type:varchar(255);comment:Avatar"` 
    Username string `json:"username" gorm:"type:varchar(50);comment:Username"` 
    Amount string `json:"amount" gorm:"type:decimal(10,2);comment:Amount"` 
    models.ModelTime
    models.ControlBy
}

func (OrdRankLeaderboard) TableName() string {
    return "ord_rank_leaderboard"
}

func (e *OrdRankLeaderboard) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrdRankLeaderboard) GetId() interface{} {
	return e.Id
}