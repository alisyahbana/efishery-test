package env

import (
	"os"
	"strings"
)

func GetEnv() string {
	env := strings.Trim(strings.ToLower(os.Getenv("Q_ENV")), " ")
	if env == "" {
		env = "development"
	}
	return env
}
