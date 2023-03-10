package model

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// LocalTime 自定义数据类型1开始
type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	//格式化秒
	seconds := t.Unix()
	return []byte(strconv.FormatInt(seconds, 13)), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type BaseEntity struct {
	// 重写id字段
	Id        int64     `json:"id" gorm:"id;primaryKey;autoIncrement;comment:主键id"`
	CreatedAt LocalTime `json:"createdAt" gorm:"created_at;comment:创建时间;"`
	//UpdatedAt LocalTime      `json:"updatedAt" gorm:"updated_at;comment:更新时间"`
	//update的莫名其妙bug
	DeletedAt gorm.DeletedAt `json:"-" gorm:"deleted_at;comment:删除时间"` // 查询这个字段但是不返回这个字段
}
