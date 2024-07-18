// transport 定义 Request 的编码和 Response 的解码。
package transport

import (
	mendpoint "Labs/lab_go-kit/ch_5/endpoint"
	"Labs/lab_go-kit/ch_5/pb"
	"Labs/lab_go-kit/ch_5/service"
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type grpcServer struct {
	login grpctransport.Handler
	pb.UnimplementedUserServer
}

func NewGRPCServer(endpoint mendpoint.Server, log *zap.Logger) pb.UserServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			ctx = context.WithValue(ctx, service.ContextReqUUid, md.Get(service.ContextReqUUid))
			return ctx
		}),
		grpctransport.ServerErrorHandler(NewZapLogErrorHandler(log)),
	}
	return &grpcServer{login: grpctransport.NewServer(
		endpoint.LoginEndPoint,
		RequestGrpcLogin,
		ResponseGrpcLogin,
		options...,
	)}
}

func (s *grpcServer) RpcUserLogin(ctx context.Context, req *pb.Login) (*pb.LoginAck, error) {
	_, rep, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginAck), nil
}

func RequestGrpcLogin(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Login)
	return &pb.Login{Account: req.GetAccount(), Password: req.GetPassword()}, nil
}

func ResponseGrpcLogin(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginAck)
	return &pb.LoginAck{Token: resp.Token}, nil
}
