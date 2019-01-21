package Controllers

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
	"net/http"
	"project/echo-firefly-web/app/Library"
	"time"
)

//mysql操作
func(this *TestController) TestGetData(c echo.Context) error {
	testList,err:=this.Services.Test.GetUserList()
	if(err != nil){
		fmt.Println(err)
	}
	return c.Render(http.StatusOK, "test.html", map[string]interface{}{
		"name": "王老三m",
		"list": testList,
	})
}

//redis操作
func(this *TestController) TestRedis(c echo.Context) error{

	redis_key := "echo_test"
	redisO := Library.InitRedis()

	test1, _ := redis.String(redisO.Do("GET", redis_key))
	err, _ := redisO.Do("SET", redis_key, time.Now().Format(http.TimeFormat), "EX", "60")
	if(err != nil){
		fmt.Println(err)
	}
	test2, _ := redis.String(redisO.Do("GET", redis_key))

	return c.Render(http.StatusOK, "redis.html", map[string]interface{}{
		"test1": test1,
		"test2": test2,
	})
}