package main

import EComApp "github.com/CodeDV8/go-ecom-app"

func main() {
	app := EComApp.NewApplication()
	app.Init()
	app.Run()
	app.Done()
}
