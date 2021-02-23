package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer - управляющая сущность
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New - возвращает дефолтный APIServer
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start - точка входа сервиса и конфигурация его
func (s *APIServer) Start() error {
	if err := s.configLogger(); err != nil {
		return err
	}
	s.logger.Info("Start API server")
	s.logger.Info("")

	s.configRouter()

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configRouter() {

	// http://localhost:8080/api
	s.router.HandleFunc("/api", s.apiHello())
}

func (s *APIServer) apiHello() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//идея засунуть сюда свегер и выводить доступные методы
		io.WriteString(rw, "HELLO API")
	}
}
