package main

import "github.com/hajbabaeim/hltchtodo/app"

func main() {
	a := app.NewApp()
	a.Init()
	a.InitModules()
	a.Start()
}
