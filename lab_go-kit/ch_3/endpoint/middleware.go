package endpoint

import (
	"Labs/lab_go-kit/ch_3/service"
	"Labs/lab_go-kit/ch_3/utils"
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"time"
)

func LoggingMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Debug(fmt.Sprint(ctx.Value(service.ContextReqUUid)), zap.Any("调用 endpoint LoggingMiddleware", "处理完请求"), zap.Any("耗时毫秒", time.Since(begin).Milliseconds()))
			}(time.Now())
			logger.Debug(fmt.Sprint(ctx.Value(service.ContextReqUUid)), zap.Any("调用 endpoint LoggingMiddleware", "开始处理请求"))
			return next(ctx, request)
		}
	}
}

func AuthMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			token := fmt.Sprint(ctx.Value(utils.JWT_CONTEXT_KEY))
			if token == "" {
				err = errors.New("请登录")
				logger.Debug(fmt.Sprint(ctx.Value(service.ContextReqUUid)), zap.Any("[AuthMiddleware]", "token == empty"), zap.Error(err))
				return "", err
			}
			jwtInfo, err := utils.ParseToken(token)
			if err != nil {
				logger.Debug(fmt.Sprint(ctx.Value(service.ContextReqUUid)), zap.Any("[AuthMiddleware]", "ParseToken"), zap.Error(err))
				return "", err
			}
			if v, ok := jwtInfo["Name"]; ok {
				ctx = context.WithValue(ctx, "name", v)
			}
			return next(ctx, request)
		}
	}
}
