package cache

import (
	"context"
	"fmt"
)

const (
	LuaCompareAndIncrClientID = "LuaCompareAndIncrClientID"
)

type luaPart struct {
	LuaScript string
	Sha       string
}

var luaScriptTable map[string]*luaPart = map[string]*luaPart{
	LuaCompareAndIncrClientID: {
		LuaScript: "if redis.call('exists', KEYS[1]) == 0 then redis.call('set', KEYS[1], 0) end;if redis.call('get', KEYS[1]) == ARGV[1] then redis.call('incr', KEYS[1]);redis.call('expire', KEYS[1], ARGV[2]); return 1 else return -1 end",
	},
}

// 初始化lua脚本
func initLuaScript(ctx context.Context) {
	for name, part := range luaScriptTable {
		cmd := rdb.ScriptLoad(ctx, part.LuaScript)
		if cmd == nil {
			panic(fmt.Sprintf("lua init failed lua=%s", name))
		}
		if cmd.Err() != nil {
			panic(cmd.Err())
		}
		part.Sha = cmd.Val()
	}
}
