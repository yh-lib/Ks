package cluster

import "server/config"

// 定义一个结构体，用于描述创建集群所用的配置信息
type ClusterInfo struct {
	Id       string `json:"id"`
	Alias    string `json:"alias"`
	Location string `json:"city"`
}

type ClusterConfig struct {
	ClusterInfo
	Kubeconfig string `json:"kubeconfig"`
}

// 定义一个结构体，用于描述集群的状态
type ClusterStatus struct {
	ClusterInfo
	Version string `json:"version"`
	Status  string `json:"kubeconfig"`
}

// 定义一个结构体的方法，用于判断集群的状态
func (c *ClusterConfig) getClusterStatus() (ClusterStatus, error) {
	// 判断集群是否是正常
	ClusterStatus := ClusterStatus{}
	ClusterStatus.ClusterInfo = c.ClusterInfo
	serverVersion, err := config.InClusterClientSet.Discovery().ServerVersion()
	if err != nil {
		return ClusterStatus, err
	}
	clusterVersion := serverVersion.String()
	ClusterStatus.Version = clusterVersion
	ClusterStatus.Status = "Active"
	return ClusterStatus, nil
}
