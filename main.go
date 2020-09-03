package main

import ecomapp "github.com/codedv8/go-ecom-app"

func main() {
	app := ecomapp.NewApplication()
	app.SysInit()
	app.Init()
	app.Run()
	app.Done()
}
