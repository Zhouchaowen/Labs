// endpoint 是 service 的调用层，还可以添加各种中间件，验证，限流，日志。
package endpoint

import (
	"Labs/lab_go-kit/ch_4/service"
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// 把Service的方法封装到Endpoint中
type Server struct {
	AddEndPoint   endpoint.Endpoint
	LoginEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc service.Service, log *zap.Logger, limit *rate.Limiter, limiter ratelimit.Limiter) Server {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
		addEndPoint = LoggingMiddleware(log)(addEndPoint)
		addEndPoint = AuthMiddleware(log)(addEndPoint)
		addEndPoint = NewUberRateMiddleware(limiter)(addEndPoint)
	}
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)
		loginEndPoint = LoggingMiddleware(log)(loginEndPoint)
		loginEndPoint = NewGolangRateAllowMiddleware(limit)(loginEndPoint)
	}
	return Server{AddEndPoint: addEndPoint, LoginEndPoint: loginEndPoint}
}

func (s Server) Add(ctx context.Context, in service.Add) service.AddAck {
	res, _ := s.AddEndPoint(ctx, in)
	return res.(service.AddAck)
}

func MakeAddEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.Add)
		res := s.TestAdd(ctx, req)
		return res, nil
	}
}

func (s Server) Login(ctx context.Context, in service.Login) (service.LoginAck, error) {
	res, err := s.LoginEndPoint(ctx, in)
	return res.(service.LoginAck), err
}

func MakeLoginEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.Login)
		return s.Login(ctx, req)
	}
}
