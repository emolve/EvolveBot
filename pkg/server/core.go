package server

import (
	"github.com/emolve/EvolveBot/pkg/config"
	"github.com/emolve/EvolveBot/pkg/selenium"
)

func Start(config config.Config) {
	sel := selenium.NewSeleniumControl(config)
	sel.Login(config)

}
