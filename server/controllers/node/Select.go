package node

import (
	"server/utils/logs"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	logs.Info(nil, "详情逻辑")
}

func List(c *gin.Context) {
	logs.Info(nil, "列表逻辑")
}
