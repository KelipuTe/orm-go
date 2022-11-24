package v20

import (
	"context"
	"database/sql"
	"orm-go/v20/metadata"
	"orm-go/v20/result"
)

// s6Monitor 控制器：构造 I9Session 的时候，控制[结果集、方言]的使用哪个实现。
type s6Monitor struct {
	// i9Registry 元数据注册中心
	i9Registry metadata.I9Registry
	// f8NewI9Result 处理"用数据库返回的查询结果构造结构体"
	f8NewI9Result result.F8NewI9Result
	// 处理方言
	i9Dialect I9Dialect

	s5f8Middleware []F8Middleware
}

func f8DoEXEC(ctx context.Context, i9Session I9Session, p7s6Monitor *s6Monitor, p7s6Context *S6QueryContext) S6Result {
	var f8HandleFunc F8MiddlewareHandle = func(ctx context.Context, p7s6Context *S6QueryContext) *S6QueryResult {
		// 查询构造器构造查询
		p7s6Query, err := p7s6Context.i9Builder.F8BuildQuery()
		if nil != err {
			return &S6QueryResult{
				I9Err: err,
			}
		}
		// 执行查询
		sqlResult, err2 := i9Session.f8DoExecContext(ctx, p7s6Query.SQLString, p7s6Query.S5Value...)
		return &S6QueryResult{
			AnyResult: sqlResult,
			I9Err:     err2,
		}
	}

	// 中间件套娃
	for i := len(p7s6Monitor.s5f8Middleware) - 1; 0 <= i; i-- {
		f8HandleFunc = p7s6Monitor.s5f8Middleware[i](f8HandleFunc)
	}
	// 执行套娃
	p7s6Result := f8HandleFunc(ctx, p7s6Context)

	// 从中间件的 S6QueryResult 里面把结果捞出来
	var i9SQLResult sql.Result = nil
	if nil != p7s6Result.AnyResult {
		i9SQLResult = p7s6Result.AnyResult.(sql.Result)
	}
	return S6Result{I9SQLResult: i9SQLResult, I9Err: p7s6Result.I9Err}
}
