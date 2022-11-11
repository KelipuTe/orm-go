package v20

var S6MySQLDialect Dialect = &s6MySQLDialect{}

type Dialect interface {
	// quoter 返回一个引号，引用列名，表名的引号
	f8GetQuoter() byte
}

type s6MySQLDialect struct {
}

func (p7this *s6MySQLDialect) f8GetQuoter() byte {
	return '`'
}
