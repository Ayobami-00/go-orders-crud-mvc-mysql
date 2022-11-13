package main

import (
	"fmt"
	"go-orders-crud-mvc-mysql/src/app"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "go-orders-crud-mvc-mysql"

func main() {

	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Printf("%s uses %s\n", os.Getenv("NAME"), os.Getenv("EDITOR"))

	app.StartApplication()
}
