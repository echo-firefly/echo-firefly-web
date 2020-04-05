package main

import (
	"echo-firefly-web/app/Library/Yamls"
	"fmt"
	"echo-firefly-web/Apollo"
	"echo-firefly-web/bootstrap"
	"echo-firefly-web/routes"
)

/*
          ,_---~~~~~----._
   _,,_,*^____      _____``*g*\"*,
  / __/ /'     ^.  /      \ ^@q   f
 [  @f | @))    |  | @))   l  0 _/
  \`/   \~____ / __ \_____/    \
   |           _l__l_           I
   }          [______]           I
   ]            | | |            |
   ]             ~ ~             |
   |                            |
    |                           |
*/

/*
* 用于框架初始化调用
 */
func init() {
	Yamls.LoadYaml()    //初始化 yaml 配置
	Apollo.InitApollo() //初始化阿波罗配置
}

/*
* 框架入口函数
* 加载 bootstrap 公共类
* 加载router文件
* 监听端口
 */
func main() {

	app := bootstrap.New("go-echo-firefly-web", "wood")
	app.Bootstrap()
	app.Configure(routes.Configure)
	fmt.Println("框架启动成功")
	app.Listen(":8001")
}
