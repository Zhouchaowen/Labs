// transport 定义 Request 的编码和 Response 的解码。
package transport

import (
	mendpoint "Labs/lab_go-kit/ch_4/endpoint"
	"Labs/lab_go-kit/ch_4/service"
	"Labs/lab_go-kit/ch_4/utils"
	"context"
	"encoding/json"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func NewHttpHandler(endpoint mendpoint.Server, log *zap.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			log.Warn(fmt.Sprint(ctx.Value(service.ContextReqUUid)), zap.Error(err))
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
		}), //程序中的全部报错都会走这里面
		//httptransport.ServerErrorHandler(NewZapLogErrorHandler(log)),
		httptransport.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			UUID := uuid.NewV5(uuid.NewV4(), "req_uuid").String()
			log.Debug("给请求添加uuid", zap.Any("UUID", UUID))
			ctx = context.WithValue(ctx, service.ContextReqUUid, UUID)
			ctx = context.WithValue(ctx, utils.JWT_CONTEXT_KEY, request.Header.Get("Authorization"))
			log.Debug("把请求中的token发到Context中", zap.Any("Token", request.Header.Get("Authorization")))

			return ctx
		}),
	}
	m := http.NewServeMux()
	m.Handle("/sum", httptransport.NewServer(
		endpoint.AddEndPoint,
		decodeHTTPADDRequest,      //解析请求值
		encodeHTTPGenericResponse, //返回值
		options...,
	))

	m.Handle("/login", httptransport.NewServer(
		endpoint.LoginEndPoint,
		decodeHTTPLoginRequest,    //解析请求值
		encodeHTTPGenericResponse, //返回值
		options...,
	))
	return m
}

func decodeHTTPLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var login service.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		return nil, err
	}
	utils.GetLogger().Debug(fmt.Sprint(ctx.Value(service.ContextReqUUid)), zap.Any(" 开始解析请求数据", login))
	return login, nil
}

func decodeHTTPADDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var (
		in  service.Add
		err error
	)
	in.A, err = strconv.Atoi(r.FormValue("a"))
	in.B, err = strconv.Atoi(r.FormValue("b"))
	if err != nil {
		return in, err
	}
	utils.GetLogger().Debug(fmt.Sprint(ctx.Value(service.ContextReqUUid)), zap.Any(" 开始解析请求数据", in))
	return in, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	utils.GetLogger().Debug(fmt.Sprint(ctx.Value(service.ContextReqUUid)), zap.Any("请求结束封装返回值", response))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorWrapper struct {
	Error string `json:"errors"`
}
