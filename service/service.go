package service

import (
	"log"
	"template-echo-gorm/config"
)

type Service struct {
	Logger *log.Logger
	Config *config.Config
}

func New() Service {
	return Service{}
}
