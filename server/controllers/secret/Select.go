package secret

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	controllers.KubectlFunc(c, "secret", "get")
}

func List(c *gin.Context) {
	controllers.KubectlFunc(c, "secret", "list")
}
