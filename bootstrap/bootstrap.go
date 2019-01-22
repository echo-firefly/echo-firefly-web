package bootstrap

/**
* desc: 使用Go内建的嵌入机制(匿名嵌入)，允许类型之前共享代码和数据
* Bootstrapper继承和共享 echo.Application
* auther: wood
*/

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"log"
	"time"
)

type Bootstrapper struct {
	*echo.Echo
	AppName		string
	AppOwner     string
	AppSpawnDate time.Time
}
type Configurator func(*Bootstrapper)

type TemplateRenderer struct {
	templates *template.Template
}

/**
* 封装 echo.new() 定义一些公共的参数 - 后期可考虑读取配置文件的方式
* appName string 网站的appname
* appOwner string 网站的开发团队
*/
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Echo:  echo.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

//echo Configure方法
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func (b *Bootstrapper) SetupViews() {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("resources/views/*.html")),
	}
	b.Renderer = renderer
}



/**
* 框架初始化的工作都在这里
* 后续这里会丰富大量的公共服务和三方模块，这里可优化的点非常多
*/
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	//加载前端模板
	b.Static("/public", "public")
	b.SetupViews()
	b.Use(middleware.Logger())
	b.Use(middleware.Recover())
	//TODO 加载http异常捕获模块
	return b
}

//开始监听 - 后续继续优化http配置
func (b *Bootstrapper) Listen(addr string) {
	log.Fatal(b.Start(addr))
}






