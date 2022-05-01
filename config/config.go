package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var (
	Host     string
	User     string
	Password string
	Database string
	Port     string
)

func SetViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error while reading viper")

	}

	runmode := cast.ToString(viper.Get("runmode"))
	databaseConfig := viper.Get("app." + runmode + ".databaseConfig").(map[string]interface{})
	Host = databaseConfig["host"].(string)
	User = databaseConfig["user"].(string)
	Password = databaseConfig["password"].(string)
	Database = databaseConfig["database"].(string)
	Port = cast.ToString(databaseConfig["port"].(int))
}

func DBConnect() (db *sql.DB, err error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", User, Database, Password, Host)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	/*
	Hoping our connection details are validated, next we’re going to call Ping() method on sql.DB object to test our connection. db.ping()
	will force open a database connection to confirm if we are successfully connected to the database.
	*/
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return
}
