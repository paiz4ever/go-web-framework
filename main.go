package main

import (
	"log"

	"github.com/paiz4ever/go-web-framework/global"
	"github.com/paiz4ever/go-web-framework/initialize"
)

func main() {
	if err := initialize.InitConfig(); err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}
	if err := initialize.InitDataBase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	if err := initialize.InitRedis(); err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}
	initialize.InitServer()
}
