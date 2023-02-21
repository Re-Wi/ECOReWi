package initialize

import (
	ecoEntity "BackendGoECOReWi/apps/economy/entity"
	logEntity "BackendGoECOReWi/apps/log/entity"
	"BackendGoECOReWi/pkg/global"
	"github.com/XM-GO/PandaKit/biz"
)

// 初始化时如果没有表创建表
func InitTable() {
	m := global.Conf.Server
	if m.IsInitTable {
		biz.ErrIsNil(
			global.Db.AutoMigrate(
				//casbin.CasbinRule{},
				logEntity.LogOper{},
				ecoEntity.EcoDynamic{},
			),
			"初始化表失败")
	}
}
