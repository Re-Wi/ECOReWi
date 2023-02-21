package global

import (
	"BackendGoECOReWi/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Log  *logrus.Logger // 日志
	Db   *gorm.DB       // gorm
	Conf *config.Config
)
