package statefulSet

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	controllers.KubectlFunc(c, "statefulSet", "delete")
}

func DeleteList(c *gin.Context) {
	controllers.KubectlFunc(c, "statefulSet", "deleteList")
}
