package discovery

import (
	"context"
	"testing"
	"time"
)

func TestServiceDiscovery(t *testing.T) {
	var endpoints = []string{"localhost:2379"}
	ctx := context.Background()
	ser := NewServiceDiscovery(&ctx, endpoints)
	defer ser.Close()
	ser.WatchService("/web/", func(key, value string) {}, func(key, value string) {})
	ser.WatchService("/gRPC/", func(key, value string) {}, func(key, value string) {})
	for {
		select {
		case <-time.Tick(10 * time.Second):
		}
	}
}
