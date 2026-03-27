package node

import (
	"server/utils/logs"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	logs.Info(nil, "创建逻辑")
}
