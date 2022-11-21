package v20

import (
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
}
