package api

import (
	"BackendGoECOReWi/apps/economy/entity"
	"BackendGoECOReWi/apps/economy/services"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"log"
)

type EcoDynamicApi struct {
	EcoDynamicApp services.EcoDynamicModel
}

// GetEcoDynamicList 列表数据
func (l *EcoDynamicApi) GetEcoDynamicList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	item := restfulx.QueryParam(rc, "item")
	remark := restfulx.QueryParam(rc, "remark")
	list, total := l.EcoDynamicApp.FindListPage(pageNum, pageSize, entity.EcoDynamic{Item: item, Remark: remark})
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetEcoDynamic 获取
func (l *EcoDynamicApi) GetEcoDynamic(rc *restfulx.ReqCtx) {
	ecoId := restfulx.PathParamInt(rc, "id")
	rc.ResData = l.EcoDynamicApp.FindOne(int64(ecoId))
}

// InsertEcoDynamic 添加
func (p *EcoDynamicApi) InsertEcoDynamic(rc *restfulx.ReqCtx) {
	var data entity.EcoDynamic
	restfulx.BindQuery(rc, &data)
	p.EcoDynamicApp.Insert(data)
}

// UpdateEcoDynamic 修改
func (p *EcoDynamicApi) UpdateEcoDynamic(rc *restfulx.ReqCtx) {
	var data entity.EcoDynamic
	restfulx.BindQuery(rc, &data)

	p.EcoDynamicApp.Update(data)
}

// DeleteEcoDynamic 删除
func (l *EcoDynamicApi) DeleteEcoDynamic(rc *restfulx.ReqCtx) {
	ecoIds := restfulx.PathParam(rc, "id")
	group := utils.IdsStrToIdsIntGroup(ecoIds)
	log.Println("group", group)
	l.EcoDynamicApp.Delete(group)
}

func (l *EcoDynamicApi) DeleteAll(rc *restfulx.ReqCtx) {
	l.EcoDynamicApp.DeleteAll()
}
