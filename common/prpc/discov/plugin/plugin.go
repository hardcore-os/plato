package plugin

import (
	"errors"
	"fmt"

	"github.com/hardcore-os/plato/common/prpc/config"
	"github.com/hardcore-os/plato/common/prpc/discov"
	"github.com/hardcore-os/plato/common/prpc/discov/etcd"
)

// GetDiscovInstance 获取服务发现实例
func GetDiscovInstance() (discov.Discovery, error) {
	name := config.GetDiscovName()
	switch name {
	case "etcd":
		return etcd.NewETCDRegister(etcd.WithEndpoints(config.GetDiscovEndpoints()))
	}

	return nil, errors.New(fmt.Sprintf("not exist plugin:%s", name))
}
