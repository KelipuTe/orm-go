package v20

import (
	"context"
	"orm-go/v20/metadata"
)

// S6QueryContext 查询上下文，给中间件用的
type S6QueryContext struct {
	// 查询类型 SELECT、INSERT、UPDATE、DELETE
	QueryType string
	// 查询构造器 S6SELECT、S6INSERT、S6UPDATE、S6DELETE
	// 如果想在中间件里使用查询构造器，需要先进行类型断言
	i9Builder I9QueryBuilder
	// p7s6Model 映射模型
	p7s6Model *metadata.S6Model

	p7s6Query *S6Query
}

func (p7this *S6QueryContext) BuildQuery() (*S6Query, error) {
	var err error = nil

	if nil == p7this.p7s6Query {
		p7this.p7s6Query, err = p7this.i9Builder.F8BuildQuery()
	}
	return p7this.p7s6Query, err
}

// S6QueryResult 查询结果，给中间件用的
type S6QueryResult struct {
	// 查询结果，不同的查询，结果类型不一样，需要进行类型断言
	// S6SELECT.First() => *T
	// S6SELECT.List() => []*T
	// S6SELECT.Get() => map[string]any
	// INSERT.EXEC() => sql.Result
	// UPDATE.EXEC() => sql.Result
	// DELETE.EXEC() => sql.Result
	AnyResult any
	I9Err     error
}

type F8MiddlewareHandle func(ctx context.Context, p7s6Context *S6QueryContext) *S6QueryResult

type F8Middleware func(next F8MiddlewareHandle) F8MiddlewareHandle
