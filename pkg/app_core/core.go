package app_core

import (
	"bufio"
	"fmt"
	"github.com/emolve/EvolveBot/pkg/config"
	"github.com/emolve/EvolveBot/pkg/selenium"
	"os"
	"strings"
)

var sel selenium.SeleniumControl

type Command map[string]func(...interface{})

type Core struct {
	Cmd  Command
	Exit chan os.Signal
}

func NewApp() (Core, error) {
	app := Core{
		Cmd: Command{
			"help": getAppHelpInfo,
			"exit": func(options ...interface{}) {
				// TODO
			},
			"mode":  getMode,
			"login": loginFishPi,
		},
	}

	return app, nil
}

func (app Core) Run() {
	config := config.GetConfig()
	sel = selenium.NewSeleniumControl(config)

	app.Cmd["help"]("")
	fmt.Print("/")

	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		input := strings.Split(buf.Text(), " ")
		inputCommand := input[0]
		inputOption := input[1:]
		if _, ok := app.Cmd[inputCommand]; !ok {
			fmt.Printf("[Error] Unknown Command: %s\n", inputCommand)
			app.Cmd["help"]("")
		} else {
			app.Cmd[inputCommand](inputOption)
		}
		fmt.Print("/")
	}
}
