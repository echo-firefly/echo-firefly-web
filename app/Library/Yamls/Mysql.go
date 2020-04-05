package Yamls

type Mysql struct {
	//user
	User struct {
		Host   string `yaml:"DB_USER_HOST"`
		Port   int    `yaml:"DB_USER_PORT"`
		User   string `yaml:"DB_USER_USER"`
		Pwd    string `yaml:"DB_USER_PASS"`
		DbName string `yaml:"DB_USER_NAME"`
	}
}
