package config

import "fmt"

type Mysql struct {
	UserName     string `mapstructure:"user-name"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	DBName       string `mapstructure:"db-name"`
	MaxIdleConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
}

func (m *Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.UserName, m.Password, m.Host, m.Port, m.DBName)
}
