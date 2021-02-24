package apiserver

import (
	"io/ioutil"
	"net/http"

	"github.com/BroNaz/adv-backend/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer - управляющая сущность
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	Store  *store.Store
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

	if err := s.configStore(); err != nil {
		//		s.logger.Error("Store ERROR ", err)
		return err
	}

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

func (s *APIServer) configStore() error {
	temp := store.New(s.config.Store)
	if err := temp.Open(); err != nil {
		return err
	}

	s.Store = temp
	return nil

}

func (s *APIServer) apiHello() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//идея засунуть сюда свегер и выводить доступные методы
		file, err := ioutil.ReadFile("static/apiHelper.html")
		if err != nil {
			s.logger.Error("Error apiHelper")
			rw.Write([]byte("501 - Not Implemented"))
			return
		}
		rw.Write(file)
	}
}
