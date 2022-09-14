package main

import "github.com/ferriantitus/alterra-agmc/day2/static-api/routes"

func main() {
	e := routes.New()
	e.Start(":8080")

	e.Logger.Fatal(e.Start("8000"))
}