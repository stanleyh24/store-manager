package main

import (
	"log"
	"os"

	"store-manager/api"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = validateEnvironments()
	if err != nil {
		log.Fatal(err)
	}

	dbPool, err := newDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	httpServer := api.NewHttp(dbPool)

	log.Fatal(httpServer.Listen(":" + os.Getenv("SERVER_PORT")))

}
