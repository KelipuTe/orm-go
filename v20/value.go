package v20

// s6value 对应查询语句里的占位符对应的参数
type s6value struct {
	value any
}

func (this s6value) doExpression() {}

// toParameter 把输入转换成查询语句里的占位符对应的参数
func toParameter(p any) s6value {
	return s6value{
		value: p,
	}
}
