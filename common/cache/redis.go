package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/hardcore-os/plato/common/config"
)

// 声明一个全局的rdb变量，这是一个单机client
var rdb *redis.Client

func InitRedis(ctx context.Context) {
	if rdb != nil {
		return
	}
	endpoints := config.GetCacheRedisEndpointList()
	opt := &redis.Options{Addr: endpoints[0], PoolSize: 10000}
	rdb = redis.NewClient(opt)
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(err)
	}
	initLuaScript(ctx)
}
func GetBytes(ctx context.Context, key string) ([]byte, error) {
	cmd := rdb.Conn().Get(ctx, key)
	if cmd == nil {
		return nil, errors.New("redis GetBytes cmd is nil")
	}
	data, err := cmd.Bytes()
	if redis.Nil == err {
		return nil, nil
	}
	return data, err
}

func GetUInt64(ctx context.Context, key string) (uint64, error) {
	cmd := rdb.Conn().Get(ctx, key)
	if cmd == nil {
		return 0, errors.New("redis GetUInt64 cmd is nil")
	}
	return cmd.Uint64()
}

func SetBytes(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	cmd := rdb.Set(ctx, key, value, ttl)
	if cmd == nil {
		return errors.New("redis SetBytes cmd is nil")
	}
	return cmd.Err()
}

func Del(ctx context.Context, key ...string) error {
	cmd := rdb.Conn().Del(ctx, key...)
	if cmd == nil {
		return errors.New("redis Del cmd is nil")
	}
	return cmd.Err()
}

func SADD(ctx context.Context, key string, member interface{}) error {
	cmd := rdb.SAdd(ctx, key, member)
	if cmd == nil {
		return errors.New("redis SADD cmd is nil")
	}
	return cmd.Err()
}

func SREM(ctx context.Context, key string, members ...interface{}) error {
	cmd := rdb.Conn().SRem(ctx, key, members...)
	if cmd == nil {
		return errors.New("redis SREM cmd is nil")
	}
	return cmd.Err()
}

func SmembersStrSlice(ctx context.Context, key string) ([]string, error) {
	cmd := rdb.Conn().SMembers(ctx, key)
	if cmd == nil {
		return nil, errors.New("redis SmembersUint64StructMap cmd is nil")
	}
	return cmd.Result()
}

func Incr(ctx context.Context, key string, ttl time.Duration) error {
	_, err := rdb.Conn().Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Incr(ctx, key)
		p.Expire(ctx, key, ttl)
		return nil
	})
	return err
}

func SetString(ctx context.Context, key string, value string, ttl time.Duration) error {
	cmd := rdb.Set(ctx, key, value, ttl)
	if cmd == nil {
		return errors.New("redis SetString cmd is nil")
	}
	return cmd.Err()
}

func GetString(ctx context.Context, key string) (string, error) {
	cmd := rdb.Get(ctx, key)
	if cmd == nil {
		return "", errors.New("redis GetString cmd is nil")
	}
	return cmd.String(), cmd.Err()
}

func RunLuaInt(ctx context.Context, name string, keys []string, args ...interface{}) (int, error) {
	if part, ok := luaScriptTable[name]; !ok {
		return -1, errors.New("lua not registered")
	} else {
		cmd := rdb.EvalSha(ctx, part.Sha, keys, args...)
		if cmd == nil {
			return -1, errors.New("redis RunLua cmd is nil")
		}

		return cmd.Int()
	}
}

func GetKeys(ctx context.Context, key string) ([]string, error) {
	return rdb.Keys(ctx, key).Result()
}

// redis cache
type redisCache struct {
}

func newRedisCache(opt *Options) *redisCache {
	return &redisCache{}
}

func (r *redisCache) MSet(keys map[string]interface{}) {

}

func (r *redisCache) MGet(key []string) map[string]interface{} {
	return nil
}

func (r *redisCache) MDel(key []string) {

}
