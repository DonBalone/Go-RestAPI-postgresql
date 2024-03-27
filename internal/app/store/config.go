package store

type Config struct {
	Storage string `yaml:"storage_path"`
}

/*type Storage struct {
	storageURL string `yaml:"storage_path"`
}*/

func NewConfig() *Config {
	return &Config{}
}
