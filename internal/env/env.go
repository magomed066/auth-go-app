package env

import (
	"fmt"
	"os"
)

func GetPort() string {
	val := os.Getenv("PORT")
	
	if val == "" {
		return ":8080"
	}

	return fmt.Sprint(":", val)
}