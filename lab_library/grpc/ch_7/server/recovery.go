package main

import (
	"fmt"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// RecoveryInterceptor panic时返回Unknown错误吗
func RecoveryInterceptor() grpc_recovery.Option {
	fmt.Println("run RecoveryInterceptor")
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		return grpc.Errorf(codes.Unknown, "panic triggered: %v", p)
	})
}
