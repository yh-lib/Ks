package cluster

import "github.com/gin-gonic/gin"

func Get(c *gin.Context) {
	returnData.Message = "查看详情成功"
	c.JSON(200, returnData)
}
