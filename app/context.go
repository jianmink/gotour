package app

import (
	"context"
)

type Req struct{
	kv map[string]string
}

func (r *Req)Get(k string)string{
	return r.kv[k]
}

func (r *Req)Set(k,v string) {
	r.kv[k] = v
}

type key int
const requestIDKey key = 0

func newContextWithRequestID(ctx context.Context, req *Req) context.Context {
	reqID := req.Get("X-Request-ID")
	if reqID == "" {
		reqID = "123456"
	}

	return context.WithValue(ctx, requestIDKey, reqID)
}

func requestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}