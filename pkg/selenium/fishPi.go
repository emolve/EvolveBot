package selenium

import (
	"fmt"
	"github.com/emolve/EvolveBot/pkg/config"
	"github.com/tebeka/selenium"
)

type fishPiEvolve struct {
	webDriver selenium.WebDriver
}

func (f fishPiEvolve) Login(config config.Config) {
	wd := f.webDriver
	wd.Get(config.User.LoginPath)
	we, _ := wd.FindElement(selenium.ByID, "nameOrEmail")
	we.SendKeys(config.User.Username)
	we, _ = wd.FindElement(selenium.ByID, "loginPassword")
	we.SendKeys(config.User.Password)
	// login button
	we, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[3]/div/div[1]/div/button[1]")
	we.Click()
	if err != nil {
		fmt.Println("[Error] Login FindElement error: ", err)
	}
}
