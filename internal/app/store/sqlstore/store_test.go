package sqlstore_test

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"testing"
)

var (
	StorageInfo string
	configPath  string
)

func TestMain(m *testing.M) {
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/apiserver.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	var config map[string]interface{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarshalling YAML: %v", err)
	}

	// Получение значения StorageInfo из YAML файла и запись его в переменную Username
	StorageInfo = "host=localhost port=5432 sslmode=disable dbname=restapi_test password=12345 user=postgres"

	os.Exit(m.Run())
}
