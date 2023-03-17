package router

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hardcore-os/plato/common/cache"
)

const (
	gatewayRotuerKey = "gateway_rotuer_%d"
	ttl7D            = 7 * 24 * 60 * 60
)

type Record struct {
	Endpoint string
	ConndID  uint64
}

func Init(ctx context.Context) {
	cache.InitRedis(ctx)
}

func AddRecord(ctx context.Context, did uint64, endpoint string, conndID uint64) error {
	key := fmt.Sprintf(gatewayRotuerKey, did)
	value := fmt.Sprintf("%s-%d", endpoint, conndID)
	return cache.SetString(ctx, key, value, ttl7D*time.Second)
}
func DelRecord(ctx context.Context, did uint64) error {
	key := fmt.Sprintf(gatewayRotuerKey, did)
	return cache.Del(ctx, key)
}
func QueryRecord(ctx context.Context, did uint64) (*Record, error) {
	key := fmt.Sprintf(gatewayRotuerKey, did)
	data, err := cache.GetString(ctx, key)
	if err != nil {
		return nil, err
	}
	ec := strings.Split(data, "-")
	conndID, err := strconv.ParseUint(ec[1], 10, 64)
	if err != nil {
		return nil, err
	}
	return &Record{
		Endpoint: ec[0],
		ConndID:  conndID,
	}, nil
}
