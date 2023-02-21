package services

import (
	"BackendGoECOReWi/apps/economy/entity"
	"BackendGoECOReWi/pkg/global"
	"github.com/XM-GO/PandaKit/biz"
)

type (
	EcoDynamicModel interface {
		Insert(data entity.EcoDynamic) *entity.EcoDynamic
		FindOne(infoId int64) *entity.EcoDynamic
		FindListPage(page, pageSize int, data entity.EcoDynamic) (*[]entity.EcoDynamic, int64)
		FindList(data entity.EcoDynamic) *[]entity.EcoDynamic
		Update(data entity.EcoDynamic) *entity.EcoDynamic
		Delete(infoId []int64)
		DeleteAll()
	}

	ecoDynamicModelImpl struct {
		table string
	}
)

var EcoDynamicModelDao EcoDynamicModel = &ecoDynamicModelImpl{
	table: `eco_dynamics`,
}

func (m *ecoDynamicModelImpl) Insert(data entity.EcoDynamic) *entity.EcoDynamic {
	global.Db.Table(m.table).Create(&data)
	return &data
}

func (m *ecoDynamicModelImpl) FindOne(operId int64) *entity.EcoDynamic {
	resData := new(entity.EcoDynamic)
	err := global.Db.Table(m.table).Where("id = ?", operId).First(resData).Error
	biz.ErrIsNil(err, "查询动态经济信息失败")
	return resData
}

func (m *ecoDynamicModelImpl) FindListPage(page, pageSize int, data entity.EcoDynamic) (*[]entity.EcoDynamic, int64) {
	list := make([]entity.EcoDynamic, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Item != "" {
		db = db.Where("item = ?", data.Item)
	}
	if data.Remark != "" {
		db = db.Where("remark like ?", "%"+data.Remark+"%")
	}
	if data.Sign != "" {
		db = db.Where("sign like ?", "%"+data.Sign+"%")
	}
	err := db.Where("delete_time IS NULL").Count(&total).Error
	err = db.Order("create_time desc").Limit(pageSize).Offset(offset).Find(&list).Error

	biz.ErrIsNil(err, "查询分页动态经济信息失败")
	return &list, total
}
func (m *ecoDynamicModelImpl) FindList(data entity.EcoDynamic) *[]entity.EcoDynamic {
	list := make([]entity.EcoDynamic, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询动态经济列表失败")
	return &list
}

func (m *ecoDynamicModelImpl) Update(data entity.EcoDynamic) *entity.EcoDynamic {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改动态经济失败")
	return &data
}

func (m *ecoDynamicModelImpl) Delete(ids []int64) {
	err := global.Db.Table(m.table).Delete(&entity.EcoDynamic{}, "id in (?)", ids).Error
	biz.ErrIsNil(err, "删除动态经济信息失败")
	return
}

func (m *ecoDynamicModelImpl) DeleteAll() {
	global.Db.Exec("DELETE FROM eco_dynamics")
}
