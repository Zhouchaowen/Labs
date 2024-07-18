package transport

import (
	"Labs/lab_go-kit/ch_5/service"
	"context"
	"fmt"
	"go.uber.org/zap"
)

type LogErrorHandler struct {
	logger *zap.Logger
}

func NewZapLogErrorHandler(logger *zap.Logger) *LogErrorHandler {
	return &LogErrorHandler{
		logger: logger,
	}
}

func (h *LogErrorHandler) Handle(ctx context.Context, err error) {
	h.logger.Warn(fmt.Sprint(ctx.Value(service.ContextReqUUid)), zap.Error(err))
}
