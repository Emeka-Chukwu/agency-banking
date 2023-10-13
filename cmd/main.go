package main

import (
	server "agency-banking/app"
	"agency-banking/util"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Cannot set config: %v", err.Error())
		log.Fatal(err.Error())
	}

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatalf("Cannot set config: %v", err.Error())
		log.Fatal(err.Error())
	}
	_, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	fmt.Println("server running at: ", config.HTTPServerAddress)
	// server.StartGin(config.HTTPServerAddress)
	server.StartChi(config.HTTPServerAddress)

}
