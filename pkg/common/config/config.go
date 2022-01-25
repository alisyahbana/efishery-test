package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
)

func GetAppURL() url.URL {
	return url.URL{
		Scheme: os.Getenv("APP_SCHEME"),
		Host:   os.Getenv("APP_HOST"),
	}
}

const appName = "efishery-test"

func LoadConfiguration(config interface{}, service string, env string) {
	configPathMap := make(map[string]string)
	configPathMap["development"] = fmt.Sprintf("config/%s/%s.development.json", service, service)
	configPathMap["staging"] = fmt.Sprintf("/etc/%s/%s.staging.json", appName, service)
	configPathMap["production"] = fmt.Sprintf("/etc/%s/%s.production.json", appName, service)

	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	fullpath := ""

	if env == "development" {
		if configDir, exists := os.LookupEnv("CONFIG_DIR"); exists && configDir != "" {
			fullpath = fmt.Sprintf("%s/%s/%s.development.json", configDir, appName, service)
		} else {
			fullpath = basePath + "/../../../" + configPathMap[env]
		}
	} else {
		fullpath = configPathMap[env]
	}

	configFile, err := os.Open(fullpath)
	defer configFile.Close()
	if err != nil {
		panic(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
}
