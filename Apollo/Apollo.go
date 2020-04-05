package Apollo

/**
* apollo配置中心初始化
* 主要是用于从远端Apollo拉取,mysql redis MongoDB和其他配置
* 采用热更新的方式，实时的监控配置中心，有新的改动时会注销连接池，重新建立连接
 */

import (
	"fmt"
	"github.com/zouyx/agollo"
	"os"
	"echo-firefly-web/app/Library"
	"echo-firefly-web/app/Library/Yamls"
	"echo-firefly-web/config/database"
)

type Apollo struct {
	Off           string //1开启 2关闭
	APPID         string
	CLUSTER       string
	NAMESPACENAME string
	IP            string
}

//初始化Apollo配置
func InitApollo() {
	conf := getApolloConf()
	if conf.Off == "0" {
		readyConfig := &agollo.AppConfig{
			AppId:            conf.APPID,
			Cluster:          conf.CLUSTER,
			NamespaceName:    conf.NAMESPACENAME,
			Ip:               conf.IP,
			BackupConfigPath: "/tmp",
		}
		agollo.InitCustomConfig(func() (*agollo.AppConfig, error) {
			return readyConfig, nil
		})

		//启动Apollo服务
		err := agollo.Start()
		if err != nil {
			fmt.Println("启动Apollo服务失败", err)
		}

		//热更新加载
		event := agollo.ListenChangeEvent()
		go func() {
			for {
				select {
				case e := <-event:
					destruct(e)
				}
			}
		}()
	}
	loadConf()
}

/**
* 销毁连接
* 目前只支持变动后销毁所有的连接
* event暂时不用
 */
func destruct(event *agollo.ChangeEvent) {
	if database.Load() != nil {
		fmt.Println("数据配置重新加载失败!")
	}
	//销毁mysql对象
	Library.DestructionMysql()
	//销毁redis对象
	Library.DestructionRedis()
	//销毁MongoDB
	Library.DestructionMgo()

}

func loadConf() {
	if database.Load() != nil {
		fmt.Println("数据配置加载失败!")
	}
}

//加载Apollo配置
func getApolloConf() Apollo {
	apolloConf := Yamls.GetConf().Apollo
	var apollo Apollo
	var appid, cluster, namespacename, ip string
	siteEnv := os.Getenv("ACTIVE")
	//只有在yaml配置 开启 && 非线上环境下才执行
	if apolloConf.OPEN == "1" && siteEnv != "pro" {
		if apolloConf.OPEN == "1" {
			apollo.Off = "1"
			return apollo
		}
		appid = apolloConf.APPID
		cluster = apolloConf.CLUSTER
		namespacename = apolloConf.NAMESPACENAME
		ip = apolloConf.IP
	} else {
		appid = os.Getenv("APPID")
		cluster = os.Getenv("CLUSTER")
		namespacename = os.Getenv("NAMESPACENAME")
		ip = os.Getenv("IP")
	}
	apollo = Apollo{
		Off:           "1",
		APPID:         appid,
		CLUSTER:       cluster,
		NAMESPACENAME: namespacename,
		IP:            ip,
	}
	return apollo
}
