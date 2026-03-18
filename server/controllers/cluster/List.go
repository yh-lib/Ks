package cluster

import (
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	returnData.Message = "获取list成功"
	c.JSON(200, returnData)
}
