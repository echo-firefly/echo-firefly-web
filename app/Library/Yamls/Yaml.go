package Yamls

import (
"gopkg.in/yaml.v2"
"io/ioutil"
"log"
)

//用于解析yaml
type Base struct {
	Mysql  Mysql
	Mongo  Mongo
	Redis  Redis
	Apollo Apollo
	Other Other
}

var Yamls *Base
var YamlsOff int

//yaml 解析公共类
func LoadYaml() *Base {
	conf := new(Base)
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		//文件不存在则默认使用Apollo配置
		YamlsOff = 1
		return conf
		//log.Fatal("yaml加载失败:", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatal("yaml解析失败:", err)
	}
	Yamls = conf
	YamlsOff = 0
	return conf
}

func GetConf() *Base {
	if YamlsOff == 1 {
		var b = new(Base)
		b.Apollo.OPEN = "0"
		return b
	}
	return Yamls
}

