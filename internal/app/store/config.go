package store

type Config struct {
	Storage string `yaml:"storage_path"`
}

/*type Storage struct {
	storageURL string `yaml:"storage_path"`
}*/

func NewConfig() *Config {
	return &Config{
		Storage: "host=localhost port=5432 sslmode=disable dbname=restapi_dev password=12345 user=postgres",
	}
}
