package v20

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
			err := t4value.F8BuildColumn(p7s6QueryBuilder)
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
