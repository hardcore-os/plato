package resolver

import (
	"context"
	"fmt"

	"github.com/hardcore-os/plato/common/prpc/discov"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"
)

type DiscovBuilder struct {
	discov discov.Discovery
}

// NewDiscovBuilder ...
func NewDiscovBuilder(d discov.Discovery) resolver.Builder {
	return &DiscovBuilder{
		discov: d,
	}
}

func (d *DiscovBuilder) Scheme() string {
	return DiscovBuilderScheme
}

func (d *DiscovBuilder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	d.discov.GetService(context.TODO(), d.getServiceName(target))
	serviceName := d.getServiceName(target)
	listener := func() {
		service := d.discov.GetService(context.TODO(), serviceName)
		var addrs []resolver.Address
		for _, item := range service.Endpoints {
			attr := attributes.New("weight", item.Weight)
			addr := resolver.Address{
				Addr:       fmt.Sprintf("%s:%d", item.IP, item.Port),
				Attributes: attr,
			}

			addrs = append(addrs, addr)
		}

		cc.UpdateState(resolver.State{
			Addresses: addrs,
		})
	}

	d.discov.AddListener(context.TODO(), listener)
	listener()

	return d, nil
}

func (d *DiscovBuilder) getServiceName(target resolver.Target) string {
	return target.Endpoint
}

func (d *DiscovBuilder) Close() {
}

func (d *DiscovBuilder) ResolveNow(options resolver.ResolveNowOptions) {
}
