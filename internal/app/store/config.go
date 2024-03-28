package store

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	StorageInfo string `yaml:"storage_info"`
}

func NewConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	var config map[string]interface{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarshalling YAML: %v", err)
	}

	// Получение значения username из YAML файла и запись его в переменную Username
	StorageInfo, ok := config["storage_info"].(string)
	if !ok {
		log.Fatalf("storage_info is not a string")
	}

	return &Config{
		StorageInfo: StorageInfo,
	}
}
