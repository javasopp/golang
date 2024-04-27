package config

type Config struct {
	AppConfig      App
	DatabaseConfig Database
	RedisConfig    Redis
}

type App struct {
	Port     int
	Host     string
	LogLevel string
	LogFile  string
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

type Redis struct {
	Host     string
	Port     int
	Password string
	Db       int
}
