package middleware

import (
	"context"
	v20 "orm-go/v20"
	"time"
)

// SlowLogTriggerMiddlewareBuild 触发慢 SQL 用的
func SlowLogTriggerMiddlewareBuild() v20.F8Middleware {
	return func(next v20.F8MiddlewareHandle) v20.F8MiddlewareHandle {
		return func(ctx context.Context, p7s6Context *v20.S6QueryContext) *v20.S6QueryResult {
			time.Sleep(500 * time.Millisecond)
			return next(ctx, p7s6Context)
		}
	}
}
