package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var flags = flag.NewFlagSet("goose", flag.ExitOnError)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()
	fmt.Println(len(args))
	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	err1 := godotenv.Load("/home/ubuntu/app/goose.env")
	if err1 != nil {
		log.Fatalf("Error loading .env file")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	user := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	sslmode := os.Getenv("SSLMODE")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, databaseName, sslmode)

	db, err := goose.OpenDBWithDriver("postgres", dsn)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 2 {
		arguments = append(arguments, args[2:]...)
	}

	dir := "/home/ubuntu/app/migrations"
	if err := goose.Run(command, db, dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
