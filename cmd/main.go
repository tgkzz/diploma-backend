package main

import (
	redis "github.com/redis/go-redis/v9"
	"log"
	"os"
	"server/internal/config"
	"server/internal/handler"
	"server/internal/logger"
	"server/internal/repository"
	"server/internal/server"
	"server/internal/service"
)

func main() {
	infoLog, errLog, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("NewLogger: %s", err)
	}

	var cfgPath string

	switch len(os.Args[1:]) {
	case 1:
		cfgPath = os.Args[1]
	default:
		cfgPath = "./.env"
	}

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		errLog.Fatalf("LoadConfig %s", err)
	}

	db, err := repository.NewDB(cfg)
	if err != nil {
		errLog.Fatalf("NewDB %s", err)
	}

	mongoClient, err := repository.NewMongoDB(cfg.Mongo)
	if err != nil {
		errLog.Fatalf("NewMongoDB %s", err)
	}

	redisCli := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
	})

	r := repository.NewRepository(db, mongoClient)

	s := service.NewService(*r, cfg.SecretKey, cfg.MailSenderKey, redisCli)

	h := handler.NewHandler(s, infoLog, errLog, cfg.SecretKey)

	errLog.Fatal(server.RunServer(cfg, h.Routes(), infoLog))
}
