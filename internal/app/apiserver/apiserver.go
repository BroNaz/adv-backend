package apiserver

import "github.com/sirupsen/logrus"

// APIServer - управляющая сущность
type APIServer struct {
	config *Config
	logger *logrus.Logger
}

// New - возвращает дефолтный APIServer
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start - точка входа сервиса и конфигурация его
func (s *APIServer) Start() error {
	if err := s.configLogger(); err != nil {
		return err
	}
	s.logger.Info("Start API server")

	return nil
}

func (s *APIServer) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}
