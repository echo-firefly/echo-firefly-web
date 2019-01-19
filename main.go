package main

import (
	"project/echo_web/bootstrap"
	"project/echo_web/routes"
)

func main() {
	app := bootstrap.New("go-echo-firefly-web", "wood")
	app.Bootstrap()
	app.Configure(routes.Configure)
	app.Listen(":8081")
	/*app.Configure(routes.Configure)
	//注册
	app.Listen(":8081")*/
}
