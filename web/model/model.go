package model

import (
	"blog-console/common"
	"blog-console/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID          uint32 `gorm:"primary_key" json:"id"`
	CreateUser  string `json:"create_user"`
	CreateTime  uint32 `json:"create_time"`
	UpdateUser  string `json:"update_user"`
	UpdateTime  uint32 `json:"update_time"`
	DeletedTime uint32 `json:"delete_time"`
	IsDel       uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *common.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
