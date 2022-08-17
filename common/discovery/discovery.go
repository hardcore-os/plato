package discovery

import (
	"context"
	"sync"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
)

//ServiceDiscovery 服务发现
type ServiceDiscovery struct {
	cli  *clientv3.Client //etcd client
	lock sync.Mutex
	ctx  *context.Context
}

//NewServiceDiscovery  新建发现服务
func NewServiceDiscovery(ctx *context.Context, endpoints []string) *ServiceDiscovery {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logger.Fatal(err)
	}

	return &ServiceDiscovery{
		cli: cli,
		ctx: ctx,
	}
}

//WatchService 初始化服务列表和监视
func (s *ServiceDiscovery) WatchService(prefix string, set, del func(key, value string)) error {
	//根据前缀获取现有的key
	resp, err := s.cli.Get(*s.ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	for _, ev := range resp.Kvs {
		set(string(ev.Key), string(ev.Value))
	}
	//监视前缀，修改变更的server
	s.watcher(prefix, set, del)
	return nil
}

//watcher 监听前缀
func (s *ServiceDiscovery) watcher(prefix string, set, del func(key, value string)) {
	rch := s.cli.Watch(*s.ctx, prefix, clientv3.WithPrefix())
	logger.CtxInfof(*s.ctx, "watching prefix:%s now...", prefix)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT: //修改或者新增
				set(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE: //删除
				del(string(ev.Kv.Key), string(ev.Kv.Value))
			}
		}
	}
}

//Close 关闭服务
func (s *ServiceDiscovery) Close() error {
	return s.cli.Close()
}
