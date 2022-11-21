package v20

import "orm-go/v20/metadata"

// S6Column 对应 col_name
// 即语句中表示[表、JOIN、子查询]中列的部分
type S6Column struct {
	// i9From 列对应的[表、JOIN、子查询]
	i9From i9TableReference
	// fieldName 结构体属性名
	fieldName string
	// alias 数据库列名的别名
	alias string
}

// f8BuildColumn 构造列
// p7s6Builder 查询构造器
func (this S6Column) f8BuildColumn(p7s6Builder *s6QueryBuilder, isUseAlias bool) error {
	var ok bool = false
	var p7s6ModelField *metadata.S6ModelField
	var columnName string = ""

	// 处理表
	if nil != this.i9From {
		// 校验属性存不存在，如果存在，就转换成数据库列名
		s5Entity := this.i9From.f8GetTableReferenceEntity()
		for _, t4value := range s5Entity {
			p7s6Model, err := p7s6Builder.s6Monitor.i9Registry.F8Get(t4value)
			if nil == err {
				p7s6ModelField, ok = p7s6Model.M3FieldToColumn[this.fieldName]
				if ok {
					// 找到
					ok = true
					columnName = p7s6ModelField.ColumnName
					break
				}
			}
		}
		// 处理表的别名
		alies := this.i9From.f8GetTableReferenceAlies()
		if "" != alies {
			p7s6Builder.f8WrapWithQuote(alies)
			p7s6Builder.sqlString.WriteByte('.')
		}
	}
	// 上面的逻辑没找到属性，就走默认的校验一次
	if !ok {
		// 校验属性存不存在，存在转换成数据库列名
		p7s6ModelField, ok = p7s6Builder.p7s6Model.M3FieldToColumn[this.fieldName]
		if !ok {
			return f8NewErrUnknowStructField(this.fieldName)
		}
		columnName = p7s6ModelField.ColumnName
	}
	p7s6Builder.f8WrapWithQuote(columnName)
	// 处理列的别名
	if isUseAlias && "" != this.alias {
		p7s6Builder.sqlString.WriteString(" AS ")
		p7s6Builder.sqlString.WriteString(this.alias)
	}
	return nil
}

// F8As 给表设置别名
func (this S6Column) F8As(alias string) S6Column {
	return S6Column{
		i9From:    this.i9From,
		fieldName: this.fieldName,
		alias:     alias,
	}
}

func (this S6Column) f8BuildSelectExpr(p7s6qb *s6QueryBuilder) error {
	return this.f8BuildColumn(p7s6qb, true)
}

func (this S6Column) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	return this.f8BuildColumn(p7s6Builder, false)
}

func (this S6Column) F8BuildAssignment() error { return nil }

func (this S6Column) F8Equal(p any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorEqual,
		i9RightExpr: f8NewI9Expression(p),
	}
}

func (this S6Column) F8GreaterThan(p any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorGreaterThan,
		i9RightExpr: f8NewI9Expression(p),
	}
}

func (this S6Column) F8LessThan(p any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorLessThan,
		i9RightExpr: f8NewI9Expression(p),
	}
}

func (this S6Column) F8InQuery(sub S6SubQuery) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorIn,
		i9RightExpr: sub,
	}
}

func F8NewS6Column(name string) S6Column {
	return S6Column{
		fieldName: name,
		alias:     "",
	}
}
