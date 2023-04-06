package user_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/response"
	"go.opentelemetry.io/otel/attribute"
)

func filterAttrs(req *request.Filter) []attribute.KeyValue {
	attrs := make([]attribute.KeyValue, 0)

	if req.Limit != nil {
		attrs = append(attrs, attribute.Int("req.Limit", *req.Limit))
	}
	if req.Offset != nil {
		attrs = append(attrs, attribute.Int("req.Offset", *req.Offset))
	}
	if req.ID != nil {
		attrs = append(attrs, attribute.Int("req.ID", int(*req.ID)))
	}
	if req.Username != nil {
		attrs = append(attrs, attribute.String("req.Username", *req.Username))
	}

	return attrs
}

func resAttrs(header string, res response.UserView) []attribute.KeyValue {
	attrs := []attribute.KeyValue{
		attribute.String(header+".Username", res.Username),
		attribute.String(header+".Role", res.Role),
	}

	if res.LastLogin != nil {
		attrs = append(attrs, attribute.String(header+".LastLogin", (*res.LastLogin).String()))
	}

	return attrs
}
