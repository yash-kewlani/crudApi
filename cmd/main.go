package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/yash-kewlani/crudApi/cmd/api"
	"github.com/yash-kewlani/crudApi/config"
	"github.com/yash-kewlani/crudApi/db"
)

func main() {
	database, err := db.NewSqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	db.InitSqlStorage(database)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", database)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
