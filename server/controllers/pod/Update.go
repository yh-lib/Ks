package pod

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	controllers.KubectlFunc(c, "pod", "update")
}
