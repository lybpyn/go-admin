package service

import (
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
)

const (
	tenjinBizTypeOrder      = "order"
	tenjinBizTypeWithdrawal = "withdrawal"
)

func tenjinOrderEventName(platform string) string {
	if strings.ToLower(strings.TrimSpace(platform)) == "ios" {
		return "CardpartnerIOS_SuccessfulTransaction"
	}
	return "Cardpartner_sell_success"
}

func tenjinWithdrawalEventName(platform string) string {
	if strings.ToLower(strings.TrimSpace(platform)) == "ios" {
		return "CardpartnerIOS_SuccessfulWithdrawal"
	}
	return "Cardpartner_withdraw_success"
}

func insertTenjinReport(db *gorm.DB, bizType, bizId, eventName string, userId int64) (bool, error) {
	if db == nil {
		return false, errors.New("db is nil")
	}

	record := models.HsTenjinReportedAt{
		BizType:    bizType,
		BizId:      bizId,
		UserId:     userId,
		EventName:  eventName,
		ReportedAt: time.Now(),
	}

	err := db.Create(&record).Error
	if err == nil {
		return true, nil
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return false, nil
	}
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return false, nil
	}

	return false, err
}

func deleteTenjinReport(db *gorm.DB, bizType, bizId, eventName string) error {
	if db == nil {
		return errors.New("db is nil")
	}
	return db.Where("biz_type = ? AND biz_id = ? AND event_name = ?", bizType, bizId, eventName).
		Delete(&models.HsTenjinReportedAt{}).Error
}
