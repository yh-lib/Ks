package pod

import (
	"context"
	"server/config"
	"server/controllers"
	"server/utils/logs"
	"strings"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(c *gin.Context) {
	logs.Info(nil, "删除逻辑")
	returnData := config.ReturnData{}
	clientSet, basicInfo, err := controllers.BasicInit(c, nil)
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "clientSet 初始化失败")
		returnData.Message = err.Error()
		returnData.Status = 400
		c.JSON(200, returnData)
		return
	}
	for _, podName := range basicInfo.DeleteList {
		err = clientSet.CoreV1().Pods(basicInfo.NameSpace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
		if err != nil {
			logs.Error(map[string]any{"Error": err}, "删除 pod: "+podName+" 失败:")
			returnData.Status = 400
			returnData.Message = "删除 pod: " + podName + " 失败:" + err.Error()
			c.JSON(200, returnData)
			return
		}
	}
	returnData.Status = 200
	returnData.Message = "删除 pod: " + strings.Join(basicInfo.DeleteList, ",") + " 成功"
	c.JSON(200, returnData)
}
