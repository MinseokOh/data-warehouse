package db

import "fmt"

type Config struct {
	Host       string `json:"host"`
	DriverName string `json:"driver_name"`
	User       string `json:"user"`
	Password   string `json:"password"`
	DBName     string `json:"db_name"`
	DBPort     int    `json:"db_port"`
	IdleConn   int    `json:"idle_conn"`
	MaxConn    int    `json:"max_conn"`
}

func (config Config) GetDriverName() string {
	return config.DriverName
}

func (config Config) GetDBSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.User,
		config.Password,
		config.Host,
		config.DBPort,
		config.DBName,
	)
}
