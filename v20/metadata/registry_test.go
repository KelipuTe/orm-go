package metadata

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type S6APPUserModel struct {
	Id   int
	Name string
	Age  int8
	Sex  int8
}

func TestI9RegistryF8Get(p7tt *testing.T) {
	s5s6case := []struct {
		name    string
		input   any
		wantRes *S6Model
		wantErr error
	}{
		{
			// 指针
			name:  "pointer",
			input: &S6APPUserModel{},
			wantRes: &S6Model{
				TableName: "s6_test_model",
				M3StructToField: map[string]*S6ModelField{
					"Id": {
						StructName: "Id",
						I9Type:     reflect.TypeOf(int(0)),
						Offset:     0,
						FieldName:  "id",
					},
					"s6Column": {
						StructName: "s6Column",
						I9Type:     reflect.TypeOf(""),
						Offset:     8,
						FieldName:  "name",
					},
					"Age": {
						StructName: "Age",
						I9Type:     reflect.TypeOf(int8(0)),
						Offset:     24,
						FieldName:  "age",
					},
					"Sex": {
						StructName: "Sex",
						I9Type:     reflect.TypeOf(int8(0)),
						Offset:     25,
						FieldName:  "sex",
					},
				},
				M3FieldToStruct: map[string]*S6ModelField{
					"id": {
						StructName: "Id",
						I9Type:     reflect.TypeOf(int(0)),
						Offset:     0,
						FieldName:  "id",
					},
					"name": {
						StructName: "s6Column",
						I9Type:     reflect.TypeOf(""),
						Offset:     8,
						FieldName:  "name",
					},
					"age": {
						StructName: "Age",
						I9Type:     reflect.TypeOf(int8(0)),
						Offset:     24,
						FieldName:  "age",
					},
					"sex": {
						StructName: "Sex",
						I9Type:     reflect.TypeOf(int8(0)),
						Offset:     25,
						FieldName:  "sex",
					},
				},
			},
		},
	}

	i9registry := F8NewI9Registry()
	for _, s6case := range s5s6case {
		p7tt.Run(s6case.name, func(p7tt *testing.T) {
			model, err := i9registry.F8Get(s6case.input)
			assert.Equal(p7tt, s6case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7tt, s6case.wantRes, model)
		})
	}
}
