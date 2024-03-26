package apiserver

// структура нашего сервера
type APIServer struct {
	config *Config
}

// создание нового сервера
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// запуск сервера
func (s *APIServer) Start() error {

	return nil
}
