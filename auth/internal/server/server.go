package server

import (
	"auth/config"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func RunServer(cfg config.Config, e *echo.Echo, infoLog *log.Logger) error {
	server := &http.Server{
		Addr:         cfg.Host + cfg.Port,
		Handler:      e,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("server is listening on http://%s%s", cfg.Host, cfg.Port)

	err := e.StartServer(server)

	return err
}
