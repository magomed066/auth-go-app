package env

import (
	"fmt"
	"os"
)

type ENVConfig struct {
	PORT string
	DB_URL string
}

func getPort() string {
	val := os.Getenv("PORT")
	
	if val == "" {
		return ":8080"
	}

	return fmt.Sprint(":", val)
}

func GetEnvs() ENVConfig {
	dbUrl := os.Getenv("DB_URL")

	envs := ENVConfig{
		PORT: getPort(),
		DB_URL: dbUrl,
	}

	return envs
}