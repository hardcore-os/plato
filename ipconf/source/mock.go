package source

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/hardcore-os/plato/common/discovery"
)

// func init() {
// 	ctx := context.Background()
// 	testServiceRegister(&ctx, "7896", "node1")
// 	testServiceRegister(&ctx, "7897", "node2")
// 	testServiceRegister(&ctx, "7898", "node3")
// }
func testServiceRegister(ctx *context.Context, port, node string) {
	// 模拟服务发现
	go func() {
		ed := discovery.EndpointInfo{
			IP:   "127.0.0.1",
			Port: port,
			MetaData: map[string]interface{}{
				"connect_num":   float64(rand.Int63n(12312321231231131)),
				"message_bytes": float64(rand.Int63n(1231232131556)),
			},
		}
		sr, err := discovery.NewServiceRegister(ctx, []string{"localhost:2379"}, fmt.Sprintf("/plato/ip_dispatcher/%s", node), &ed, time.Now().Unix())
		if err != nil {
			panic(err)
		}
		go sr.ListenLeaseRespChan()
		for {
			ed = discovery.EndpointInfo{
				IP:   "127.0.0.1",
				Port: port,
				MetaData: map[string]interface{}{
					"connect_num":   float64(rand.Int63n(12312321231231131)),
					"message_bytes": float64(rand.Int63n(1231232131556)),
				},
			}
			sr.UpdateValue(&ed)
			time.Sleep(1 * time.Second)
		}
	}()
}
