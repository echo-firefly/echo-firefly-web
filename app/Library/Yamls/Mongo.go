package Yamls

type Mongo struct {
	//user
	User struct {
		Host   string `yaml:"MONGODB_USER_HOST"`
		Port   int    `yaml:"MONGODB_USER_PORT"`
		User   string `yaml:"MONGODB_USER_USER"`
		Pwd    string `yaml:"MONGODB_USER_PASS"`
		DbName string `yaml:"MONGODB_USER_NAME"`
		Replicaset string `yaml:"MONGODB_USER_REPLICASET"`
	}
}
