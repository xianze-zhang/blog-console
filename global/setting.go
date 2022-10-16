package global

import (
	"blog-console/common"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *common.ServerSettingS
	AppSetting      *common.AppSettingS
	DatabaseSetting *common.DatabaseSettingS
	Logger          *common.Logger
	DBEngine        *gorm.DB
)
