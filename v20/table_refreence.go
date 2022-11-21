package v20

// i9TableReference 对应 table_references
// SELECT 语句 FROM 后面的：表、JOIN 查询、子查询
type i9TableReference interface {
	f8BuildTableReference(p7s6qb *s6QueryBuilder) error
	f8GetTableReferenceAlies() string
	f8GetTableReferenceEntity() []any
}
