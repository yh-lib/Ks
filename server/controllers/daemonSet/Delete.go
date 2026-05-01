package daemonSet

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	controllers.KubectlFunc(c, "daemonSet", "delete")
}

func DeleteList(c *gin.Context) {
	controllers.KubectlFunc(c, "daemonSet", "deleteList")
}
