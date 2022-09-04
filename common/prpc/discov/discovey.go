package discov

import (
	"context"
)

type Discovery interface {
	// Name 服务发现名字 eg etcd zk consul
	Name() string
	// Register 注册服务
	Register(ctx context.Context, service *Service)
	// UnRegister 取消注册服务
	UnRegister(ctx context.Context, service *Service)
	// GetService 获取服务节点信息
	GetService(ctx context.Context, name string) *Service
	// AddListener 增加监听者
	AddListener(ctx context.Context, f func())
	// NotifyListeners 通知所有的监听者
	NotifyListeners()
}
