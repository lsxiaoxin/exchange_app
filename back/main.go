package main

import (
	"exchange_app/config"
	"exchange_app/rounter"
)

func main() {
	config.InitDB()
	config.InitRedis()
	r := rounter.SetupRouter()
	r.Run()
}