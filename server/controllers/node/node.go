package node

import (
	"context"
	"server/config"
	"server/utils/logs"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NodesNum() int {
	// 获取节点列表
	nodes, err := config.InClusterClientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(nil, "获取节点数量失败")
	}
	return len(nodes.Items)
}
