package main

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"log"
)

func ZapInterceptor() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to initialize zap logger: %v", err)
	}
	grpc_zap.ReplaceGrpcLogger(logger)
	return logger
}
