package frsMySQL

import "time"

type MySQLConfig struct {
	ConnectionUrl      string
	MaxPoolSize        int
	MaxIdleConnections int
	ConnMaxIdleTime    time.Duration
	ConnMaxLifeTime    time.Duration
}
