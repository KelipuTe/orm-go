package v20

import "strings"

type s6QueryBuilder struct {
	s6Monitor
	// 这个东西 s6Monitor 里面有，拿到这里方便操作
	quote byte
	// sqlString 带有占位符的 SQL 语句
	sqlString strings.Builder
	// s5parameter SQL 语句中占位符对应的参数
	s5parameter []any
}

func (p7this *s6QueryBuilder) f8WrapWithQuote(name string) {
	p7this.sqlString.WriteByte(p7this.quote)
	p7this.sqlString.WriteString(name)
	p7this.sqlString.WriteByte(p7this.quote)
}

// addParameter 添加占位符对应的参数
func (p7this *s6QueryBuilder) addParameter(s5p ...any) {
	if nil == p7this.s5parameter {
		p7this.s5parameter = make([]any, 0, 2)
	}
	p7this.s5parameter = append(p7this.s5parameter, s5p...)
}
