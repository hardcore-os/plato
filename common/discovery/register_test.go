package discovery

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestServiceRegiste(t *testing.T) {
	ctx := context.Background()
	ser, err := NewServiceRegister(&ctx, "/web/node1", &EndpointInfo{
		IP:   "127.0.0.1",
		Port: "9999",
	}, 5)
	if err != nil {
		log.Fatalln(err)
	}
	//监听续租相应chan
	go ser.ListenLeaseRespChan()
	select {
	case <-time.After(20 * time.Second):
		ser.Close()
	}
}
