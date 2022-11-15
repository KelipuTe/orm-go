package v20

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
