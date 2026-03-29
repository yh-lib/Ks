package pod

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	controllers.KubectlFunc(c, "pod", "delete")
}

func DeleteList(c *gin.Context) {
	controllers.KubectlFunc(c, "pod", "deleteList")
}
