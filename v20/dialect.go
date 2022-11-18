package v20

var S6MySQLDialect I9Dialect = &s6MySQLDialect{}
var S6SQLite3Dialect I9Dialect = &s6SQLite3Dialect{}

type I9Dialect interface {
	// quoter 返回一个引号，引用列名，表名的引号
	f8GetQuoter() byte
	f8BuildOnConflict(*s6QueryBuilder, *S6Conflict) error
}

// #### MySQL ####

type s6MySQLDialect struct {
}

func (p7this *s6MySQLDialect) f8GetQuoter() byte {
	return '`'
}

func (p7this *s6MySQLDialect) f8BuildOnConflict(p7s6QueryBuilder *s6QueryBuilder, p7s6Conflict *S6Conflict) error {
	p7s6QueryBuilder.sqlString.WriteString(" ON DUPLICATE KEY UPDATE ")
	for i, t4value := range p7s6Conflict.S5Assignment {
		if 0 < i {
			p7s6QueryBuilder.sqlString.WriteByte(',')
		}
		switch t4value2 := t4value.(type) {
		case S6Column:
			p7s6QueryBuilder.f8WrapWithQuote(t4value2.Name)
			p7s6QueryBuilder.sqlString.WriteString("=VALUES(")
			p7s6QueryBuilder.f8WrapWithQuote(t4value2.Name)
			p7s6QueryBuilder.sqlString.WriteString(")")
		case S6Assignment:
			p7s6QueryBuilder.f8WrapWithQuote(t4value2.Name)
			p7s6QueryBuilder.sqlString.WriteString("=")
			return p7s6QueryBuilder.F8BuildExpression(t4value2.Expr)
		default:
			return NewErrUnsupportedExpressionType(t4value2)
		}
	}
	return nil
}

// #### SQLite3 ####

type s6SQLite3Dialect struct {
}

func (p7this *s6SQLite3Dialect) f8GetQuoter() byte {
	return '`'
}

func (p7this *s6SQLite3Dialect) f8BuildOnConflict(p7s6QueryBuilder *s6QueryBuilder, p7s6Conflict *S6Conflict) error {
	p7s6QueryBuilder.sqlString.WriteString(" ON CONFLICT ")

	if 0 < len(p7s6Conflict.S5ConflictColumn) {
		p7s6QueryBuilder.sqlString.WriteByte('(')
		for i, t4value := range p7s6Conflict.S5ConflictColumn {
			if 0 < i {
				p7s6QueryBuilder.sqlString.WriteByte(',')
			}
			err := t4value.f8BuildColumn(p7s6QueryBuilder)
			if nil != err {
				return err
			}
		}
		p7s6QueryBuilder.sqlString.WriteByte(')')
	}

	p7s6QueryBuilder.sqlString.WriteString(" DO UPDATE SET ")

	for i, t4value := range p7s6Conflict.S5Assignment {
		if 0 < i {
			p7s6QueryBuilder.sqlString.WriteByte(',')
		}
		switch t4value2 := t4value.(type) {
		case S6Column:
			p7s6QueryBuilder.f8WrapWithQuote(t4value2.Name)
			p7s6QueryBuilder.sqlString.WriteString("=excluded.")
			p7s6QueryBuilder.f8WrapWithQuote(t4value2.Name)
		case S6Assignment:
			p7s6QueryBuilder.f8WrapWithQuote(t4value2.Name)
			p7s6QueryBuilder.sqlString.WriteString("=")
			return p7s6QueryBuilder.F8BuildExpression(t4value2.Expr)
		default:
			return NewErrUnsupportedExpressionType(t4value2)
		}
	}

	return nil
}
