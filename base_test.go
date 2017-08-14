package redis

import (
	"fmt"
	"reflect"
	"testing"
)

var Redis *RedisPool = NewRedisPool("127.0.0.1", "6379", "testpassword")

func TestRedisPool_Do(t *testing.T) {
	reply, err := Redis.Do("SET", "key", "haha")
	if err != nil {
		t.Errorf("Do SET error:%s\n", err)
	}
	if reply != "OK" {
		t.Errorf("Do SET reply isn't OK:%s\n", reply)
	}

	reply, err = Redis.Do("GET", "key")
	if err != nil {
		t.Errorf("Do GET error:%s\n", err)
	}
	if string(reply.([]uint8)) != "haha" {
		fmt.Println("reply:", reflect.TypeOf(reply))
		t.Errorf("Do GET reply isn't haha:%s\n", reply)
	}

}

func TestRedisPool_Send(t *testing.T) {
	Redis.GetConn()
	defer Redis.Close()
	Redis.Send("AUTH", "testwrongpassword")
	Redis.Flush()
	reply, err := Redis.Receive()
	if err.Error() != "ERR invalid password"{
		t.Errorf("receive:%s,expect:ERR invalid password\n",err)
	}
	if reply != nil{
		t.Errorf("recevie result:%s,expect:nil\n",reply)
	}

	Redis.Send("AUTH", "testpassword")
	Redis.Flush()
	reply, err = Redis.Receive()
	if err != nil{
		t.Errorf("receive:%s\n",err)
	}

	if reply != "OK"{
		t.Errorf("recevie result:%s,expect:OK\n",reply)
	}

}
