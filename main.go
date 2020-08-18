package main

import EComApp "github.com/codedv8/go-ecom-app"

func main() {
	app := EComApp.NewApplication()
	app.Init()
	app.Run()
	app.Done()
}
