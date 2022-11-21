package v20

import (
	"fmt"
)

func NewErrUnsupportedExpressionType(e any) error {
	return fmt.Errorf("orm: 不支持的表达式 %v", e)
}

func f8NewErrUnknowStructField(name string) error {
	return fmt.Errorf("orm: 结构体中不存在 %v 属性。", name)
}
