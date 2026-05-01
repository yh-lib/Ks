package daemonSet

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	controllers.KubectlFunc(c, "daemonSet", "get")
}

func List(c *gin.Context) {
	controllers.KubectlFunc(c, "daemonSet", "list")
}
