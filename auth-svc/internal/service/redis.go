// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	IRedis interface {
		CheckKeyExist(ctx context.Context, key string) (err error)
		Set(ctx context.Context, key string, value interface{}) (err error)
		SetEx(ctx context.Context, key string, value interface{}, ttl uint) (err error)
		Get(ctx context.Context, key string) (*gvar.Var, error)
		HSet(ctx context.Context, key string, value g.Map) (err error)
		HSetEx(ctx context.Context, key string, value g.Map, ttl uint) (err error)
		HGet(ctx context.Context, key string) (value *gvar.Var, err error)
		HMSet(ctx context.Context, key string, value g.Map) (err error)
		HMSetEx(ctx context.Context, key string, value g.Map, ttl uint) (err error)
		HMGet(ctx context.Context, key string, keyValue ...string) (value []string, err error)
		Del(ctx context.Context, key string) (err error)
		Expire(ctx context.Context, key string, ttl time.Duration) (err error)
		FlushAll(ctx context.Context) (err error)
		Incr(ctx context.Context, key string) (num int64, err error)
	}
)

var (
	localRedis IRedis
)

func Redis() IRedis {
	if localRedis == nil {
		panic("implement not found for interface IRedis, forgot register?")
	}
	return localRedis
}

func RegisterRedis(i IRedis) {
	localRedis = i
}
