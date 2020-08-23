package main

import EComApp "github.com/codedv8/go-ecom-app"

func main() {
	app := EComApp.NewApplication()
	app.SysInit()
	app.Init()
	app.Run()
	app.Done()
}
