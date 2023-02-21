package router

import (
	"BackendGoECOReWi/apps/economy/api"
	"BackendGoECOReWi/apps/economy/entity"
	"BackendGoECOReWi/apps/economy/services"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitDynamicEcoRouter(container *restful.Container) {
	// 操作日志
	s := &api.EcoDynamicApi{
		EcoDynamicApp: services.EcoDynamicModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/economy/ecoDynamic").Produces(restful.MIME_JSON)
	tags := []string{"ecoDynamic"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取操作日志列表").Handle(s.GetEcoDynamicList)
	}).
		Doc("获取操作动态经济列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("businessType", "businessType").DataType("string")).
		Param(ws.QueryParameter("operName", "operName").DataType("string")).
		Param(ws.QueryParameter("title", "title").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取操作动态经济信息").Handle(s.GetEcoDynamic)
	}).
		Doc("获取操作动态经济信息").
		Param(ws.PathParameter("id", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.EcoDynamic{}). // on the response
		Returns(200, "OK", entity.EcoDynamic{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加动态经济信息").Handle(s.InsertEcoDynamic)
	}).
		Doc("添加动态经济信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.EcoDynamic{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改动态经济信息").Handle(s.UpdateEcoDynamic)
	}).
		Doc("修改动态经济信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.EcoDynamic{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除操作动态经济信息").Handle(s.DeleteEcoDynamic)
	}).
		Doc("删除操作动态经济信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	ws.Route(ws.DELETE("/all").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("清空操作动态经济信息").Handle(s.DeleteAll)
	}).
		Doc("清空操作动态经济信息").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)
}
