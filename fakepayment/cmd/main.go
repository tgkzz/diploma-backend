package main

import (
	"fakepayment/config"
	"fakepayment/internal/handler"
	"fakepayment/internal/repository"
	"fakepayment/internal/server"
	"fakepayment/internal/service"
	"fakepayment/logger"
	"fmt"
	"os"
)

func main() {
	infoLog, errLog, err := logger.NewLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	var cfgPath string

	switch len(os.Args[1:]) {
	case 1:
		cfgPath = os.Args[1]
	case 0:
		cfgPath = "./.env"
	default:
		errLog.Print("USAGE: go run [CONFIG_PATH]")
		return
	}

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		errLog.Print(err)
		return
	}

	db, err := repository.NewDB(cfg)
	if err != nil {
		errLog.Print(err)
		return
	}

	r := repository.NewRepository(db)

	s := service.NewService(*r, cfg.SecretKey)

	h := handler.NewHandler(s, infoLog, errLog)

	if err := server.RunServer(cfg, h.Routes(), infoLog); err != nil {
		errLog.Print(err)
		return
	}
}
