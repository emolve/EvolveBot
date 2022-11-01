package config

type Config struct {
	App struct {
		Mode string
	}
	User struct {
		LoginPath string
		Username  string
		Password  string
	}
	Selenium struct {
		DriverPath string
		Port       int
	}
}
