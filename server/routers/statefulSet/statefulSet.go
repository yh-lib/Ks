package statefulSet

import (
	"server/controllers/statefulSet"

	"github.com/gin-gonic/gin"
)

func create(statefulSetGroup *gin.RouterGroup) {
	statefulSetGroup.POST("/create", statefulSet.Create)
}
func delete(statefulSetGroup *gin.RouterGroup) {
	statefulSetGroup.POST("/delete", statefulSet.Delete)
}
func deleteList(statefulSetGroup *gin.RouterGroup) {
	statefulSetGroup.POST("/deleteList", statefulSet.DeleteList)
}
func update(statefulSetGroup *gin.RouterGroup) {
	statefulSetGroup.POST("/update", statefulSet.Update)
}
func get(statefulSetGroup *gin.RouterGroup) {
	statefulSetGroup.GET("/get", statefulSet.Get)
}
func list(statefulSetGroup *gin.RouterGroup) {
	statefulSetGroup.GET("/list", statefulSet.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	statefulSetGroup := g.Group("statefulSet")
	create(statefulSetGroup)
	delete(statefulSetGroup)
	deleteList(statefulSetGroup)
	update(statefulSetGroup)
	get(statefulSetGroup)
	list(statefulSetGroup)
}
