package main

import (
	"github.com/met4tron/finances-api/internal/server"
	"github.com/met4tron/finances-api/pkg/logger"
)

func main() {
	logger.InitLogger()

	s := server.NewServer()

	s.Init()
}
