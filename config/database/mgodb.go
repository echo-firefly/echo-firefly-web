package database

import (
	"github.com/zouyx/agollo"
	"os"
	"echo-firefly-web/app/Library/Yamls"
)

type MgoDriver struct{}
type MgoDbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
	Replicaset string
}

var mgoConf map[string]MgoDbConf = nil

//初始化MongoDB连接
func (this *MgoDriver) Load() error {
	if mgoConf == nil {
		mgoConf = make(map[string]MgoDbConf)
	}

	apolloConf := Yamls.GetConf().Apollo
	siteEnv := os.Getenv("ACTIVE")
	//生产环境必须走Apollo
	if apolloConf.OPEN == "1" && siteEnv != "pro" { //yaml配置
		mongoConf := Yamls.GetConf().Mongo
		this.yamlConf(mongoConf)
	} else {
		this.apolloConf()
	}
	return nil
}

//选择数据库
func (this *MgoDriver) SwitchDb(name string) MgoDbConf {
	MgoDb := mgoConf[name]
	return MgoDb
}

//获取Apollo配置
func (this *MgoDriver) apolloConf() {
	mgoConf["user"] = MgoDbConf{
		Host:   agollo.GetStringValue("MONGODB_USER_HOST", ""),
		Port:   agollo.GetIntValue("MONGODB_USER_PORT", 27017),
		User:   agollo.GetStringValue("MONGODB_USER_USER", ""),
		Pwd:    agollo.GetStringValue("MONGODB_USER_PASS", ""),
		DbName: agollo.GetStringValue("MONGODB_USER_NAME", ""),
		Replicaset: "",
	}
}

//yaml配置
func (this *MgoDriver) yamlConf(mongoConf Yamls.Mongo) {
	mgoConf["user"] = MgoDbConf{
		Host:   mongoConf.user.Host,
		Port:   mongoConf.user.Port,
		User:   mongoConf.user.User,
		Pwd:    mongoConf.user.Pwd,
		DbName: mongoConf.user.DbName,
		Replicaset: "",
	}
}
