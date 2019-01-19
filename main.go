package main

import (
	"project/echo-firefly-web/bootstrap"
	"project/echo-firefly-web/routes"
)

func main() {
	app := bootstrap.New("go-echo-firefly-web", "wood")
	app.Bootstrap()
	app.Configure(routes.Configure)
	app.Listen(":8081")
}
