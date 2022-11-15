package v20

var S6MySQLDialect I9Dialect = &s6MySQLDialect{}
var S6SQLite3Dialect I9Dialect = &s6SQLite3Dialect{}

type I9Dialect interface {
	// quoter 返回一个引号，引用列名，表名的引号
	f8GetQuoter() byte
	f8BuildOnConflict(*s6QueryBuilder, *S6Conflict) error
}
