package consts

import (
	"fmt"
	"strings"
)

var (
	RedisKeyRefreshToken = RedisKey{PrefixKey: "rf_token"}
)

type RedisKey struct {
	PrefixKey string
}

// Key will generate a Redis key with the specified suffix.
// Eg: <Rediskey>.Key(abc-bcs)
func (k *RedisKey) Key(key ...string) string {
	keyString := strings.Join(key, "-")
	return fmt.Sprintf("%v:%v", k.PrefixKey, keyString)
}
