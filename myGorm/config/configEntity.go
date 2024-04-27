package config

type Config struct {
	App App
	Db  Db
}

type App struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	LogLevel string `yaml:"log_level"`
	LogFile  string `yaml:"log_file"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}
