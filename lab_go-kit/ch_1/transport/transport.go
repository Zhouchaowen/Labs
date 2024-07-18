// transport 定义 Request 的编码和 Response 的解码。
package transport

import (
	mendpoint "Labs/lab_go-kit/ch_1/endpoint"
	"Labs/lab_go-kit/ch_1/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"strconv"
)

func NewHttpHandler(endpoint mendpoint.Server) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder), //程序中的全部报错都会走这里面
	}

	m := http.NewServeMux()
	m.Handle("/sum", httptransport.NewServer(
		endpoint.AddEndPoint,
		decodeHTTPADDRequest,      //解析请求值
		encodeHTTPGenericResponse, //返回值
		options...,
	))
	return m
}
func decodeHTTPADDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("transport.decodeHTTPADDRequest")
	var (
		in  service.Add
		err error
	)
	in.A, err = strconv.Atoi(r.FormValue("a"))
	in.B, err = strconv.Atoi(r.FormValue("b"))
	if err != nil {
		return in, err
	}
	return in, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("transport.encodeHTTPGenericResponse", response)
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	fmt.Println("transport.errorEncoder", err.Error())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

type errorWrapper struct {
	Error string `json:"errors"`
}
