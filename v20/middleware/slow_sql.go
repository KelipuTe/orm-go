package middleware

import (
	"context"
	"log"
	v20 "orm-go/v20"
	"time"
)

// SlowLogMiddlewareBuild 计算查询执行时间，用于捕获慢 SQL
func SlowLogMiddlewareBuild() v20.F8Middleware {
	return func(next v20.F8MiddlewareHandle) v20.F8MiddlewareHandle {
		return func(ctx context.Context, p7s6Context *v20.S6QueryContext) *v20.S6QueryResult {
			timeStart := time.Now()
			t4 := next(ctx, p7s6Context)
			timeEnd := time.Now()
			timeCost := timeEnd.Sub(timeStart).Milliseconds()
			log.Printf("time pass %d ms\r\n", timeCost)
			if 200 < timeCost {
				log.Printf("slow sql, time pass %d ms\r\n", timeCost)
			}
			return t4
		}
	}
}
