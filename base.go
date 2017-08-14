package redis

import (
	gredis "github.com/garyburd/redigo/redis"
	"log"
)

type RedisPool struct {
	pool     *gredis.Pool
	conn     gredis.Conn
	password string
	ip       string
	port     string
}

func NewRedisPool(ip, port, password string) *RedisPool {
	pool := &gredis.Pool{
		MaxIdle:   100,
		MaxActive: 100000,
		Dial: func() (gredis.Conn, error) {
			host := ip + ":" + port
			conn, err := gredis.Dial("tcp", host)
			if err != nil {
				log.Printf("connect to %s errror:%s", host, err)
				return nil, err
			}
			if password != "" {
				err := conn.Send("AUTH", password)
				if err != nil {
					log.Printf("redis password wrong:%s", err)
				}
				err = conn.Flush()
				if err != nil {
					log.Printf("redis password wrong:%s", err)
				}
			}
			return conn, err
		},
	}
	return &RedisPool{pool:pool, password:password, ip:ip, port:port}
}

func (r *RedisPool) getConn() gredis.Conn {
	return r.pool.Get()
}

func (r *RedisPool) GetConn(){
	r.conn = r.getConn()
}

func (r *RedisPool) Close()error{
	return r.conn.Close()
}

func (r *RedisPool) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := r.getConn()
	defer conn.Close()
	return conn.Do(commandName, args...)
}

func (r *RedisPool) Send(commandName string, args ...interface{}) error {
	return r.conn.Send(commandName, args...)
}

func (r *RedisPool) Flush() error{
	return r.conn.Flush()
}

func (r *RedisPool) Receive() (reply interface{},err error) {
	return r.conn.Receive()
}
