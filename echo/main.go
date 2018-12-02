package main

import "github.com/ktny/go-api-server-sample/echo/route"

func main() {
	e := route.Init()
	e.Logger.Fatal(e.Start(":8080"))
}
