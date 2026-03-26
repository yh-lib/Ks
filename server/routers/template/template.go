package template

import (
	"server/controllers/template"

	"github.com/gin-gonic/gin"
)

func create(templateGroup *gin.RouterGroup) {
	templateGroup.POST("/create", template.Create)
}
func delete(templateGroup *gin.RouterGroup) {
	templateGroup.GET("/delete", template.Delete)
}
func update(templateGroup *gin.RouterGroup) {
	templateGroup.POST("/update", template.Update)
}
func get(templateGroup *gin.RouterGroup) {
	templateGroup.GET("/get", template.Get)
}
func list(templateGroup *gin.RouterGroup) {
	templateGroup.GET("/list", template.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	templateGroup := g.Group("template")
	create(templateGroup)
	delete(templateGroup)
	update(templateGroup)
	get(templateGroup)
	list(templateGroup)
}
