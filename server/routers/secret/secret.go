package secret

import (
	"server/controllers/secret"

	"github.com/gin-gonic/gin"
)

func create(secretGroup *gin.RouterGroup) {
	secretGroup.POST("/create", secret.Create)
}
func delete(secretGroup *gin.RouterGroup) {
	secretGroup.POST("/delete", secret.Delete)
}
func deleteList(secretGroup *gin.RouterGroup) {
	secretGroup.POST("/deleteList", secret.DeleteList)
}
func update(secretGroup *gin.RouterGroup) {
	secretGroup.POST("/update", secret.Update)
}
func get(secretGroup *gin.RouterGroup) {
	secretGroup.GET("/get", secret.Get)
}
func list(secretGroup *gin.RouterGroup) {
	secretGroup.GET("/list", secret.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	secretGroup := g.Group("secret")
	create(secretGroup)
	delete(secretGroup)
	deleteList(secretGroup)
	update(secretGroup)
	get(secretGroup)
	list(secretGroup)
}
