package apiserver

// параметры конфига
type Config struct {
	LogLevel   string `yaml:"log_level" env-default:"debug"`
	HTTPServer `yaml:"http_server"`
}
type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}

// создание нового конфига
func NewConfig() *Config {
	return &Config{
		HTTPServer: HTTPServer{
			Address: "localhost:8080",
		},
		LogLevel: "debug",
	}
}
