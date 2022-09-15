package main

import (
	"github.com/ferriantitus/alterra-agmc/day2/dynamic-api/config"
	"github.com/ferriantitus/alterra-agmc/day2/dynamic-api/routes"
	m "github.com/ferriantitus/alterra-agmc/day2/dynamic-api/middleware"
)

func main() {
	config.InitDB()
	e := routes.New()

	m.LogMiddleware(e)

	e.Start(":8080")

	e.Logger.Fatal(e.Start("8000"))
}
