package settings

import "os"

type Config struct {
	FrontendAddr    string
	UserServiceAddr string
}

func GetConfig() *Config {
	frontendAddr := os.Getenv("TC_FRONTEND_ADDR")
	if frontendAddr == "" {
		frontendAddr = ":8081"
	}
	userServiceAddr := os.Getenv("TC_USER_SERVICE_ADDR")
	if userServiceAddr == "" {
		userServiceAddr = ":8082"
	}
	return &Config{
		FrontendAddr:    frontendAddr,
		UserServiceAddr: userServiceAddr,
	}
}