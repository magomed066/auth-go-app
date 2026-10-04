package env

import (
	"fmt"
	"os"
	"strconv"

	"charm.land/log/v2"
)

type ENVConfig struct {
	PORT 			 	  string
	DB_URL 				  string
	JWT_ACCESS_TOKEN	  string
	JWT_ACCESS_EXPIRES_IN int
}

func getPort() string {
	val := os.Getenv("PORT")
	
	if val == "" {
		return ":8080"
	}

	return fmt.Sprintf(":%s", val)
}

func getEnvVal(str string) string {
	val := os.Getenv(str)

	if val == "" {
		log.Fatal("missing required environment variable: %s", str)
	}

	return val
}

func getIntEnvVal(str string) int {
	val := getEnvVal(str)
	parsedValue, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("could not parse environment variable: %s", str)
	}
	return parsedValue
}

func GetEnvs() ENVConfig {
	dbUrl := getEnvVal("DB_URL")
	jwtAccessSecret := getEnvVal("JWT_ACCESS_TOKEN")
	jwtAccessExpiresIn := getIntEnvVal("JWT_ACCESS_EXPIRES_IN")
	if jwtAccessExpiresIn <= 0 {
		log.Fatal("JWT_ACCESS_EXPIRES_IN must be greater than zero (hours)")
	}

	envs := ENVConfig{
		PORT: getPort(),
		DB_URL: dbUrl,
		JWT_ACCESS_TOKEN: jwtAccessSecret,
		JWT_ACCESS_EXPIRES_IN: jwtAccessExpiresIn,
	}

	return envs
}