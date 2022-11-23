package internal

import (
	"errors"
	"fmt"
)

// #### 元数据 ####

func F8NewErrInputOnlyStructPointer() error {
	return errors.New("orm: 只支持一级结构体指针作为输入\r\n")
}

func F8NewErrInvalidTagContent(tag string) error {
	return fmt.Errorf("orm: 标签 [%s] 格式错误\r\n", tag)
}

func F8NewErrUnknownField(name string) error {
	return fmt.Errorf("orm: 元数据中不存在 %v 属性。", name)
}

func F8NewErrUnknownColumn(column string) error {
	return fmt.Errorf("orm: 元数据中不存在 %s 列", column)
}

// #### 结果集 ####

var ErrNoRows = errors.New("orm: 未找到数据")
var ErrTooManyReturnedColumns = errors.New("orm: 返回的列过多")

// #### ORM ####

var ErrUpdateWithoutColumn = errors.New("UPDATE 没有设置更新的列")
var ErrUpdateWithoutWhere = errors.New("UPDATE 没有 WHERE")

var ErrDeleteWithoutWhere = errors.New("DELETE 没有 WHERE")

func NewErrUnsupportedExpressionType(e any) error {
	return fmt.Errorf("orm: 不支持的表达式 %v", e)
}
