package v20

// i9TableReference 对应 table_references
// SELECT 语句 FROM 后面的：表、JOIN 查询、子查询
type i9TableReference interface {
	f8BuildTableReference(p7s6Builder *s6QueryBuilder) error
	// f8CheckColumn 校验结构体属性存不存在，如果存在，就转换成数据库列名
	f8CheckColumn(p7s6Builder *s6QueryBuilder, s6Column S6Column) (string, error)
	f8GetTableReferenceAlies() string
}
