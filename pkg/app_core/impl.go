package app_core

import (
	"fmt"

	config "github.com/emolve/EvolveBot/pkg/config"
)

func getMode(options ...interface{}) {
	cfg := config.GetConfig()
	fmt.Printf("App Mode: %s\n", cfg.App.Mode)
}

func getAppHelpInfo(options ...interface{}) {
	fmt.Printf("%s\n", GetAppHelpInfo())
}

func loginFishPi(options ...interface{}) {
	cfg := config.GetConfig()
	sel.Login(cfg)
}
