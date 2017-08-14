package main

import "github.com/nicle-lin/redis"

func main() {
	red := redis.NewRedisPool("127.0.0.1","6379","dghpgyss")

	_,err := red.Do("SET","keykey","value")
	if err != nil{
		panic(err)
	}
	println("err:",err)
	str, err := red.Set("keykyey","valuevalue","N")
	println("str:",str)
	if err != nil{
		panic(err)
	}
}
