// https://www.hwholiday.com/2020/go_kit_v4/
package endpoint

import (
	"Labs/lab_go-kit/ch_4/service"
	"Labs/lab_go-kit/ch_4/utils"
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"time"
)

// Logger
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

// JWT
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

// 限流
func NewGolangRateWaitMiddleware(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if err = limit.Wait(ctx); err != nil {
				return "", errors.New("limit req  Wait")
			}
			return next(ctx, request)
		}
	}
}

// 基于golang.org/x/time/rate的限流中间件
func NewGolangRateAllowMiddleware(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return "", errors.New("limit req  Allow")
			}
			return next(ctx, request)
		}
	}
}

// 基于go.uber.org/ratelimit的限流中间件
func NewUberRateMiddleware(limit ratelimit.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			limit.Take() // 阻塞以确保满足RPS
			return next(ctx, request)
		}
	}
}
