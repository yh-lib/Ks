package cluster

import "github.com/gin-gonic/gin"

func Update(c *gin.Context) {
	returnData.Message = "更新成功"
	c.JSON(200, returnData)
}
