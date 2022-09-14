package main

import (
	"github.com/ferriantitus/day-1/config"
	"github.com/ferriantitus/day-1/routes"
)

func init() {
	config.InitDB()
	config.InitialMigration()
}

func main() {
	e := routes.New()

	e.Start(":8080")

	e.Logger.Fatal(e.Start("8000"))
}
