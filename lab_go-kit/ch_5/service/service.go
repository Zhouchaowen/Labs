// service 层一般处理业务逻辑和数据层交互等。
package service

import (
	"Labs/lab_go-kit/ch_5/pb"
	"Labs/lab_go-kit/ch_5/utils"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type Service interface {
	Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error)
}

type baseServer struct {
	logger *zap.Logger
}

func NewService(log *zap.Logger) Service {
	var server Service
	server = &baseServer{log}
	server = NewLogMiddlewareServer(log)(server)
	return server
}

func (s baseServer) Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error) {
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)), zap.Any("调用 Service", "Login 处理请求"))
	if in.Account != "imianba" || in.Password != "123456" {
		err = errors.New("用户信息错误")
		return
	}
	ack = &pb.LoginAck{}
	ack.Token, err = utils.CreateJwtToken(in.Account, 1)
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)), zap.Any("调用 Service", "Login 处理请求"), zap.Any("处理返回值", ack))
	return
}
