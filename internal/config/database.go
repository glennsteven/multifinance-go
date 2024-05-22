package config

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *sql.DB {
	driver := viper.GetString("database.driver")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	database := viper.GetString("database.name")
	idleConnection := viper.GetInt("database.pool.idle")
	maxConnection := viper.GetInt("database.pool.max")
	maxLifeTimeConnection := viper.GetInt("database.pool.lifetime")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	conn, err := sql.Open(driver, dsn)

	if err != nil {
		log.Printf("connection database got error: %v", err)
		return nil
	}

	conn.SetMaxIdleConns(idleConnection)
	conn.SetMaxOpenConns(maxConnection)
	conn.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return conn
}
