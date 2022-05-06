package main

import (
	"fmt"
	"getir-assignment/config"
	"getir-assignment/internal/client"
	inmemory "getir-assignment/internal/in_memory"
	"getir-assignment/internal/record"
	"log"
	"net/http"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "local"
	}
	conf, err := config.New(".config", appEnv)
	if err != nil {
		panic(err)
	}
	logger := buildLogger(conf.LogLevel)
	client := client.ConnectMongoDB(conf.MongoDB)
	recordRepository := record.NewRepository(client, conf.MongoDB)
	recordService := record.NewService(recordRepository, logger)
	recordHandler := record.NewHandler(recordService, logger)

	inMemoryRepository := inmemory.NewRepository()
	inMemoryService := inmemory.NewService(inMemoryRepository, logger)
	inMemoryHandler := inmemory.NewHandler(inMemoryService, logger)

	mux := http.NewServeMux()
	mux.Handle("/record", recordHandler)
	mux.Handle("/in-memory", inMemoryHandler)
	if os.Getenv("PORT") != "" {
		http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), mux)
	} else {
		http.ListenAndServe(fmt.Sprintf(":%d", conf.Server.Port), mux)
	}

}

func buildLogger(logLevel string) *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(getLogLevel(logLevel))
	loggerConfig.OutputPaths = []string{"stdout", "logs/output.log"}
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	return logger
}

func getLogLevel(level string) zapcore.Level {
	switch levelFromConfig := strings.TrimSpace(level); {
	case strings.EqualFold(levelFromConfig, "debug"):
		return zapcore.DebugLevel
	case strings.EqualFold(levelFromConfig, "error"):
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
