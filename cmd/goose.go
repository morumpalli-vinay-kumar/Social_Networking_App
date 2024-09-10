package main

import (
	_ "app/migrations"
	"flag"
	"fmt"
	"log"
	"os"

	_ "app/migrations"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", ".", "directory with migration files")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()
	fmt.Println(len(args))
	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	dsn := "postgres://postgres:postgres@localhost:5432/socialnetworkapp?sslmode=disable"
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

//
