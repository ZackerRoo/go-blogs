package global

import (
	"go-blogs/common/config"

	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
)
