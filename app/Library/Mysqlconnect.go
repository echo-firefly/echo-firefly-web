package Library

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"project/echo-firefly-web/config/database"
	"sync"
	"time"
)

var (
	lock         sync.Mutex
)

var mysqlSlaveConnect = map[string]*xorm.Engine{}
var mysqlMasterConnect = map[string]*xorm.Engine{}

// 从库，单例 - 有点羞涩
func InstancetSlave(db string) *xorm.Engine {
	if _, ok := mysqlSlaveConnect[db]; ok{
		return mysqlSlaveConnect[db]
	}
	lock.Lock()
	defer lock.Unlock()

	if _, ok := mysqlSlaveConnect[db]; ok{
		return mysqlSlaveConnect[db]
	}
	var c database.MysqlDbConf
	switch db {
		case "test":
			c = database.MysqlTestDbConfig
		case "xin_credit":
			c = database.MysqlXinCreditSlave
		case "xin":
			c = database.MysqlXinSlave
	}
	engine, err := xorm.NewEngine(database.DriverName,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			c.User, c.Pwd, c.Host, c.Port, c.DbName))
	if err != nil {
		log.Fatal("dbhelper", "DbInstanceSlave", err)
		return nil
	}
	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(true)
	var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")
	engine.SetTZLocation(SysTimeLocation)
	mysqlSlaveConnect[db] = engine
	return mysqlSlaveConnect[db]
}

// 从库，单例
func InstancetMaster(db string) *xorm.Engine {
	if _, ok := mysqlMasterConnect[db]; ok{
		return mysqlMasterConnect[db]
	}
	lock.Lock()
	defer lock.Unlock()

	if _, ok := mysqlMasterConnect[db]; ok{
		return mysqlMasterConnect[db]
	}
	var c database.MysqlDbConf
	switch db {
	case "test":
		c = database.MysqlTestDbConfig
	case "xin_credit":
		c = database.MysqlXinCreditSlave
	case "xin":
		c = database.MysqlXinSlave
	}
	engine, err := xorm.NewEngine(database.DriverName,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			c.User, c.Pwd, c.Host, c.Port, c.DbName))
	if err != nil {
		log.Fatal("dbhelper", "DbInstanceSlave", err)
		return nil
	}
	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(true)
	var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")
	engine.SetTZLocation(SysTimeLocation)
	mysqlMasterConnect[db] = engine
	return mysqlMasterConnect[db]
}

//销毁mysql, TODO记得补充 master
func DestructionMysql(){

	if mysqlSlaveConnect == nil {
		fmt.Println("未进行从数据库连接")
	} else {
		for k,v := range mysqlSlaveConnect{
			fmt.Println("正在销毁 slave mysql数据库:",k)
			v.Close()
		}
	}
	if mysqlMasterConnect == nil {
		fmt.Println("未进行主数据库连接")
	} else {
		for k,v := range mysqlMasterConnect{
			fmt.Println("正在销毁 master mysql数据库:",k)
			v.Close()
		}
	}

}

