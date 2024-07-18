// endpoint 是 service 的调用层，还可以添加各种中间件，验证，限流，日志。
package endpoint

import (
	"Labs/lab_go-kit/ch_5/pb"
	"Labs/lab_go-kit/ch_5/service"
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// 把Service的方法封装到Endpoint中
type Server struct {
	LoginEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc service.Service, log *zap.Logger, limit *rate.Limiter, limiter ratelimit.Limiter) Server {
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)
		loginEndPoint = LoggingMiddleware(log)(loginEndPoint)
		loginEndPoint = NewGolangRateAllowMiddleware(limit)(loginEndPoint)
	}
	return Server{LoginEndPoint: loginEndPoint}
}

func (s Server) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	res, err := s.LoginEndPoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.LoginAck), nil
}

func MakeLoginEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.Login)
		return s.Login(ctx, req)
	}
}
