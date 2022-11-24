package middleware

import (
	"context"
	"log"
	v20 "orm-go/v20"
)

// SqlLogMiddlewareBuild 记录查询日志
func SqlLogMiddlewareBuild() v20.F8Middleware {
	return func(next v20.F8MiddlewareHandle) v20.F8MiddlewareHandle {
		return func(ctx context.Context, p7s6Context *v20.S6QueryContext) *v20.S6QueryResult {
			// 这里提前构造查询
			p7s6Query, err := p7s6Context.F8CTXBuildQuery()
			if nil != err {
				return &v20.S6QueryResult{
					I9Err: err,
				}
			}
			// 把查询语句和参数记到日志。注意，生产环境需要进行数据脱敏处理。
			log.Println(p7s6Query.SQLString, p7s6Query.S5Value)
			return next(ctx, p7s6Context)
		}
	}
}
