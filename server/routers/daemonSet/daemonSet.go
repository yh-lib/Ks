package daemonSet

import (
	"server/controllers/daemonSet"

	"github.com/gin-gonic/gin"
)

func create(daemonSetGroup *gin.RouterGroup) {
	daemonSetGroup.POST("/create", daemonSet.Create)
}
func delete(daemonSetGroup *gin.RouterGroup) {
	daemonSetGroup.POST("/delete", daemonSet.Delete)
}
func deleteList(daemonSetGroup *gin.RouterGroup) {
	daemonSetGroup.POST("/deleteList", daemonSet.DeleteList)
}
func update(daemonSetGroup *gin.RouterGroup) {
	daemonSetGroup.POST("/update", daemonSet.Update)
}
func get(daemonSetGroup *gin.RouterGroup) {
	daemonSetGroup.GET("/get", daemonSet.Get)
}
func list(daemonSetGroup *gin.RouterGroup) {
	daemonSetGroup.GET("/list", daemonSet.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	daemonSetGroup := g.Group("daemonSet")
	create(daemonSetGroup)
	delete(daemonSetGroup)
	deleteList(daemonSetGroup)
	update(daemonSetGroup)
	get(daemonSetGroup)
	list(daemonSetGroup)
}
