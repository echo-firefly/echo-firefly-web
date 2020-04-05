package database

type Db struct {
	Mysql MySqlDriver
	Mgo   MgoDriver
}

//实例化
func BaseDb() *Db {
	var Db Db
	return &Db
}

func Load() error {
	var err error
	//重新加载mysql配置
	err = BaseDb().Mysql.Load()
	err = BaseDb().Mgo.Load()
	//加载MongoDB配置
	return err
}
