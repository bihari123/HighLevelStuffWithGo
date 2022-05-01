package main

import (
	"fmt"
	"github.com/bihari123/high_level_stuff_with_golang/config"
)

func init() {
	config.SetViper()
}

func main() {
	db, _ := config.DBConnect()
	defer db.Close()
	fmt.Println("Host:", config.Host, " User: ", config.User, " Password: ", config.Password, " database: ", config.Database)
}
