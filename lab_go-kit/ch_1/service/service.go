// service 层一般处理业务逻辑和数据层交互等。
package service

import (
	"context"
	"fmt"
)

type Service interface {
	TestAdd(ctx context.Context, in Add) AddAck
}

type baseServer struct {
}

func NewService() Service {
	return &baseServer{}
}

func (s baseServer) TestAdd(ctx context.Context, in Add) AddAck {
	fmt.Println("service.TestAdd", in)
	return AddAck{Res: in.A + in.B}
}
