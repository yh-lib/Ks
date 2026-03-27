package node

import (
	"server/utils/logs"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	logs.Info(nil, "删除逻辑")
}
