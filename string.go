package redis

import (
	gredis "github.com/garyburd/redigo/redis"
)

//APPEND key value
func (r *RedisPool) Append(key, value string) (int, error) {
	return gredis.Int(r.Do("APPEND", key, value))
}

//BITCOUNT key [start end]
func (r *RedisPool) BitCount(args ...interface{}) (int, error) {
	return gredis.Int(r.Do("BITCOUNT", args...))
}

//BITFIELD key [GET type offset] [SET type offset value] [INCRBY type offset increment] [OVERFLOW WRAP|SAT|FAIL]
func (r *RedisPool) BitField(args ...interface{}) ([]int, error) {
	return gredis.Ints(r.Do("BITFIELD", args...))
}

//BITOP operation destkey key [key ...]
func (r *RedisPool) BitTop(args ...interface{}) (int, error) {
	return gredis.Int(r.Do("BITOP", args...))
}

//BITPOS key bit [start] [end]
func (r *RedisPool) BitTops(args ...interface{}) (int, error) {
	return gredis.Int(r.Do("BITOPS", args...))
}

//DECR key
func (r *RedisPool) Decr(key string) (int, error) {
	return gredis.Int(r.Do("DECR", key))
}

//DECRBY key decrement
func (r *RedisPool) DecrBy(key string, decrement int) (int, error) {
	return gredis.Int(r.Do("DECRBY", key, decrement))
}

//GET key
func (r *RedisPool) Get(key string) ([]string, error) {
	return gredis.Strings(r.Do("GET", key))
}

//GETBIT key offset
func (r *RedisPool)GetBit(key string,offset int)(int,error){
	return gredis.Int(r.Do("GETBIT",key,offset))
}

//GETRANGE key start end
func (r *RedisPool)GetRange(key string,start,end int)([]string,error)  {
	return gredis.Strings(r.Do("GETRANGE", key,start,end))
}

//GETSET key value
func (r *RedisPool)GetSet(key,value string)(string,error){
	return gredis.String(r.Do("GETSET",key,value))
}

//INCR key
func (r *RedisPool)Incr(key string)(int, error){
	return gredis.Int(r.Do("INCR",key))
}

//INCRBY key increment
func (r *RedisPool)IncrBy(key string,incremnet int)(int, error){
	return gredis.Int(r.Do("INCRBY",key,incremnet))
}

//INCRBYFLOAT key increment
func (r *RedisPool)IncrByFloat(key string,increment float64)(float64,error)  {
	return gredis.Float64(r.Do("INCRBYFLOAT",key,increment))
}

//MGET key [key ...]
func (r *RedisPool)MGet(args ...interface{})([]interface{},error){
	return gredis.Values(r.Do("MGET",args...))
}

//MSET key value [key value ...]
func (r *RedisPool)MSet(args ...interface{})([]interface{},error){
	return gredis.Values(r.Do("MSET",args...))
}

//MSETNX key value [key value ...]
func (r *RedisPool)MSetNX(args ...interface{})(bool,error){
	return gredis.Bool(r.Do("MSETNX",args...))
}

//SET key value [EX seconds] [PX milliseconds] [NX|XX]
func (r *RedisPool) Set(args ...interface{}) (string, error) {
	return gredis.String(r.Do("SET", args...))
}

//PSETEX key milliseconds value
func (r *RedisPool) PSetEX(key, value string, milliseconds int64) (string, error) {
	return gredis.String(r.Do("PSETEX", key, value, milliseconds))
}


//SETEX key seconds value
func (r *RedisPool) SetEX(key, value string, seconds int64) (string, error) {
	return gredis.String(r.Do("SETEX", key, value, seconds))
}

//SETNX key value
func (r *RedisPool) SetNX(key, value string) (int, error) {
	return gredis.Int(r.Do("SETNX", key, value))
}

//SETBIT key offset value
func (r *RedisPool) SetBit(key string, offset int, value string) (int, error) {
	return gredis.Int(r.Do("SETBIT", key, offset, value))
}

//SETRANGE key offset value
func (r *RedisPool) SetRange(key string, offset int, value string) (int, error) {
	return gredis.Int(r.Do("SETRANGE", key, offset, value))
}

//STRLEN key
func (r *RedisPool) StrLen(key string) (int, error) {
	return gredis.Int(r.Do("SETLEN", key))
}
