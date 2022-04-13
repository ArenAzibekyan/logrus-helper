package logger

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ctxKey struct{}

var key ctxKey

func NewContext(ctx context.Context, log *logrus.Entry) context.Context {
	if log == nil {
		return ctx
	}
	return context.WithValue(ctx, key, log)
}

func FromContext(ctx context.Context) *logrus.Entry {
	log, ok := ctx.Value(key).(*logrus.Entry)
	if ok {
		return log
	}
	return nil
}

func FromRequest(req *http.Request) *logrus.Entry {
	return FromContext(req.Context())
}
