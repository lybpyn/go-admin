package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// LocalTime 自定义时间类型，支持多种时间格式的解析
type LocalTime time.Time

// MarshalJSON 序列化为 RFC3339 格式
func (t LocalTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339))
	return []byte(formatted), nil
}

// UnmarshalJSON 反序列化，支持多种时间格式
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == "\"\"" {
		return nil
	}

	// 移除引号
	str := string(data)
	if len(str) > 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}

	// 尝试多种时间格式
	formats := []string{
		"2006-01-02 15:04:05",           // MySQL datetime 格式
		time.RFC3339,                     // 2006-01-02T15:04:05Z07:00
		"2006-01-02T15:04:05",           // 不带时区的 RFC3339
		"2006-01-02 15:04:05.999999999", // 带纳秒
		time.DateTime,                    // 2006-01-02 15:04:05
	}

	var err error
	var parsed time.Time
	for _, format := range formats {
		parsed, err = time.Parse(format, str)
		if err == nil {
			*t = LocalTime(parsed)
			return nil
		}
	}

	return fmt.Errorf("无法解析时间格式: %s, 错误: %v", str, err)
}

// Value 实现 driver.Valuer 接口，用于数据库写入
func (t LocalTime) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return time.Time(t), nil
}

// Scan 实现 sql.Scanner 接口，用于数据库读取
func (t *LocalTime) Scan(value interface{}) error {
	if value == nil {
		*t = LocalTime(time.Time{})
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*t = LocalTime(v)
	case []byte:
		parsed, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		*t = LocalTime(parsed)
	case string:
		parsed, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		*t = LocalTime(parsed)
	default:
		return fmt.Errorf("无法将 %T 转换为 LocalTime", value)
	}

	return nil
}

// ToTime 转换为标准 time.Time
func (t LocalTime) ToTime() time.Time {
	return time.Time(t)
}
