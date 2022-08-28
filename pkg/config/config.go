package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

var envConfig *Config

type Config struct {
	Token string `json:"token"`
}

func GetConfig() (*Config, error) {
	if envConfig == nil {
		var err error
		envConfig, err = parseConfig()
		if err != nil {
			return nil, err
		}
	}

	return envConfig, nil
}

func parseConfig() (*Config, error) {
	currentDir := getCurrentDir()
	path := filepath.Join(currentDir, "..", "..", ".env")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	jsonProperties := make([]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		keyValue := strings.Split(line, "=")
		jsonProperties = append(jsonProperties, `"`+keyValue[0]+`":"`+keyValue[1]+`"`)
	}

	jsonString := "{" + strings.Join(jsonProperties, ",") + "}"
	config := &Config{}

	err = json.Unmarshal([]byte(jsonString), config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func getCurrentDir() string {
	_, dir, _, _ := runtime.Caller(0)
	return filepath.Dir(dir)
}
