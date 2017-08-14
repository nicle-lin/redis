package redis

import (
	gredis "github.com/garyburd/redigo/redis"
)

//BLPOP key [key ...] timeout
func (r *RedisPool)BLPop(args ...interface{})([]interface{},error)  {
	return gredis.Values(r.Do("BLPOP",args...))
}
