package main

import (
	"fmt"
	"go-orders-crud-mvc-mysql/src/app"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Printf("%s uses %s\n", os.Getenv("NAME"), os.Getenv("EDITOR"))

	app.StartApplication()
}
