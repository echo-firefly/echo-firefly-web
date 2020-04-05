package database

import (
	"github.com/zouyx/agollo"
	"os"
	"echo-firefly-web/app/Library/Yamls"
)

type MySqlDriver struct{}

const DriverName = "mysql"

//mysql公共结构体
type MysqlDbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

var dbConf map[string]MysqlDbConf

//初始化数据库连接
func (this *MySqlDriver) Load() error {
	if dbConf == nil {
		dbConf = make(map[string]MysqlDbConf)
	}
	dbConf["test.slave"] = MysqlDbConf{
		Host:   "127.0.0.1",
		Port:   3306,
		User:   "root",
		Pwd:    "123",
		DbName: "test",
	}
	apolloConf := Yamls.GetConf().Apollo
	siteEnv := os.Getenv("ACTIVE")
	//生产环境必须走Apollo
	if apolloConf.OPEN == "1" && siteEnv != "pro" { //yaml配置
		mysqlConf := Yamls.GetConf().Mysql
		this.yamlConf(mysqlConf)
	} else {
		this.apolloConf()
	}
	return nil
}

//选择数据库
func (this *MySqlDriver) SwitchDb(name string, types string) MysqlDbConf {
	XinCreditSlave := dbConf[name+"."+types]
	return XinCreditSlave
}

//获取Apollo配置
func (this *MySqlDriver) apolloConf() {
	dbConf["user.slave"] = MysqlDbConf{
		Host:   agollo.GetStringValue("DB_USER_HOST", ""),
		Port:   agollo.GetIntValue("DB_USER_PORT", 3306),
		User:   agollo.GetStringValue("DB_USER_USER", ""),
		Pwd:    agollo.GetStringValue("DB_USER_PASS", ""),
		DbName: agollo.GetStringValue("DB_USER_NAME", ""),
	}
}

//yaml配置
func (this *MySqlDriver) yamlConf(mysqlConf Yamls.Mysql) {
	dbConf["user.slave"] = MysqlDbConf{
		Host:   mysqlConf.user.Host,
		Port:   mysqlConf.user.Port,
		User:   mysqlConf.user.User,
		Pwd:    mysqlConf.user.Pwd,
		DbName: mysqlConf.user.DbName,
	}
}
