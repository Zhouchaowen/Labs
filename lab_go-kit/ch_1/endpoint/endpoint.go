// endpoint 是 service 的调用层，还可以添加各种中间件，验证，限流，日志。
package endpoint

import (
	"Labs/lab_go-kit/ch_1/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

// 把Service的方法封装到Endpoint中
type Server struct {
	AddEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc service.Service) Server {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
	}
	return Server{AddEndPoint: addEndPoint}
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
