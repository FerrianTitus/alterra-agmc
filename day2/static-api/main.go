package main

import (
	l "github.com/ferriantitus/alterra-agmc/day2/static-api/middleware"
	"github.com/ferriantitus/alterra-agmc/day2/static-api/routes"
)

func main() {
	e := routes.New()

	l.LogsMiddleware(e)

	e.Start(":8080")

	e.Logger.Fatal(e.Start("8000"))
}