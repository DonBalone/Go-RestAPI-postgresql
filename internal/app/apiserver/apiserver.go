package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// структура нашего сервера
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// создание нового сервера
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// запуск сервера
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("Starting API server")

	return http.ListenAndServe(s.config.HTTPServer.Address, s.router)
}

// создание логгера на указанном уровне
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

// создание роутера
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/huesosina", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello!")
	}
}
