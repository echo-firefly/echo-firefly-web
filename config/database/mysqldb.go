package database

const DriverName = "mysql"


type MysqlDbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}
//本地
var MysqlTestDbConfig = MysqlDbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "123",
	DbName: "test",
}