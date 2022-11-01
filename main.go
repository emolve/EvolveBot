package main

import (
	"fmt"
	"github.com/emolve/EvolveBot/pkg/config"
	"github.com/emolve/EvolveBot/pkg/server"
)

func main() {
	config := config.GetConfig()
	fmt.Println("EvolveBot Start.")
	fmt.Println("App Mode: ", config.App.Mode)
	server.Start(config)

	forever := make(chan uint)
	<-forever
}
