package Library

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var Redis redis.Conn
func InitRedis() redis.Conn {
	if Redis != nil{
		return Redis
	}
	var err error
	Redis, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err.Error())
	}
	return Redis
}

//销毁redis
func DestructionRedis(){

	if Redis == nil {
		fmt.Println("未进行redis连接")
	} else {
		fmt.Println("正在销毁redis连接")
		Redis.Close()
	}

}