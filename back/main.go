package main

import (
	"exchange_app/config"
	"exchange_app/rounter"
)

func main() {
	config.InitDB()
	r := rounter.SetupRouter()
	r.Run()
}