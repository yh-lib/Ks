package pod

import (
	"server/controllers/pod"

	"github.com/gin-gonic/gin"
)

func create(podGroup *gin.RouterGroup) {
	podGroup.POST("/create", pod.Create)
}
func delete(podGroup *gin.RouterGroup) {
	podGroup.GET("/delete", pod.Delete)
}
func update(podGroup *gin.RouterGroup) {
	podGroup.POST("/update", pod.Update)
}
func get(podGroup *gin.RouterGroup) {
	podGroup.GET("/get", pod.Get)
}
func list(podGroup *gin.RouterGroup) {
	podGroup.GET("/list", pod.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	podGroup := g.Group("pod")
	create(podGroup)
	delete(podGroup)
	update(podGroup)
	get(podGroup)
	list(podGroup)
}
