package v20

type S6TestModel struct {
	Id   int
	Name string
	Age  int8
	Sex  int8
}

func (p7this S6TestModel) F8TableName() string {
	return "test_model"
}
