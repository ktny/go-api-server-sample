package main

import (
	"flag"
	"fmt"

	"github.com/ktny/go-api-server-sample/echo/config"
	"github.com/ktny/go-api-server-sample/echo/route"
)

func main() {
	setConfig()

	c := config.Config.Database
	fmt.Printf("DBユーザー::%s", c.User)

	e := route.Init()
	e.Logger.Fatal(e.Start(":8080"))
}

func setConfig() {
	env := "dev"

	flag.Parse()
	if args := flag.Args(); 0 < len(args) && args[0] == "pro" {
		env = "production"
	}
	config.SetEnvironment(env)
}
