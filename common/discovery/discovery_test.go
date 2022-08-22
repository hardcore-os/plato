package discovery

import (
	"context"
	"testing"
	"time"
)

func TestServiceDiscovery(t *testing.T) {
	ctx := context.Background()
	ser := NewServiceDiscovery(&ctx)
	defer ser.Close()
	ser.WatchService("/web/", func(key, value string) {}, func(key, value string) {})
	ser.WatchService("/gRPC/", func(key, value string) {}, func(key, value string) {})
	for {
		select {
		case <-time.Tick(10 * time.Second):
		}
	}
}
