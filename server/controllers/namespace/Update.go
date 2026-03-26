package namespace

import (
	"context"
	"fmt"
	"server/config"
	"server/controllers"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Update(c *gin.Context) {
	logs.Info(nil, "更新逻辑")
	var ns corev1.Namespace
	returnData := config.ReturnData{}
	clientSet, _, err := controllers.BasicInit(c, &ns)
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "clientSet 初始化失败")
		returnData.Message = err.Error()
		returnData.Status = 400
		c.JSON(200, returnData)
		return
	}
	fmt.Println("测试数据：", ns)
	_, err = clientSet.CoreV1().Namespaces().Update(context.TODO(), &ns, metav1.UpdateOptions{})
	if err != nil {
		returnData.Status = 400
		returnData.Message = "更新 namespace: " + ns.Name + " 失败:" + err.Error()
		logs.Error(map[string]any{"Error": err}, "更新 namespace: "+ns.Name+" 失败:")
	} else {
		returnData.Status = 200
		returnData.Message = "更新 namespace: " + ns.Name + " 成功:"
	}
	c.JSON(200, returnData)
}
