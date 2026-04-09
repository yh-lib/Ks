package secret

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	controllers.KubectlFunc(c, "secret", "delete")
}

func DeleteList(c *gin.Context) {
	controllers.KubectlFunc(c, "secret", "deleteList")
}
