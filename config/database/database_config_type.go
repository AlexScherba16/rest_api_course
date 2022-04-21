package database

type DBConfig struct {
	Driver              string
	Host                string
	Port                int
	DBName              string
	SSLMode             string
	Username            string
	Password            string
	MaxOpenConnects     int
	MaxIdleConnects     int
	MaxLifeTimeConnects int
}
