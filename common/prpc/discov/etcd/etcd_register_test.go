package etcd

import (
	"context"
	"testing"
	"time"

	"github.com/hardcore-os/plato/common/prpc/discov"

	"github.com/stretchr/testify/assert"
)

func TestNewETCDRegister(t *testing.T) {
	_, err := NewETCDRegister()

	assert.Nil(t, err)
}

func TestRegister_Register(t *testing.T) {
	register, _ := NewETCDRegister()

	service := &discov.Service{
		Name: "test",
		Endpoints: []*discov.Endpoint{
			&discov.Endpoint{
				ServerName: "test",
				IP:         "127.0.0.1",
				Port:       9557,
				Weight:     100,
				Enable:     true,
			},
		},
	}
	register.Register(context.TODO(), service)
	time.Sleep(2 * time.Second)
	registerService := register.GetService(context.TODO(), "test")

	assert.Equal(t, *service.Endpoints[0], *registerService.Endpoints[0])
}
