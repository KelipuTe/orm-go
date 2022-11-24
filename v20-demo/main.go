package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	v20 "orm-go/v20"
)

func main() {
	testInsert()
}

func testSelect() {

}

func testInsert() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := v20.F8NewS6DB(p7s6SQLDB)

	sqlResult, err := v20.F8NewS6InsertBuilder[v20.S6APPUserModel](p7s6DB).
		F8SetEntity(&v20.S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
		F8EXEC(context.Background())
	fmt.Println(sqlResult, err)

	sqlResult, err = v20.F8NewS6InsertBuilder[v20.S6APPUserModel](p7s6DB).
		F8SetEntity(
			&v20.S6APPUserModel{Id: 22, Name: "bb", Age: 33, Sex: 2},
			&v20.S6APPUserModel{Id: 33, Name: "cc", Age: 44, Sex: 1},
		).
		F8EXEC(context.Background())
	fmt.Println(sqlResult, err)

	sqlResult, err = v20.F8NewS6InsertBuilder[v20.S6APPUserModel](p7s6DB).
		F8SetEntity(&v20.S6APPUserModel{Id: 11, Name: "aaaa", Age: 22, Sex: 1}).
		F8OnConflictBuilder().
		F8SetUpdate(v20.F8NewS6Column("Name").ToAssignment("aaaa")).
		F8EXEC(context.Background())
	fmt.Println(sqlResult, err)

	sqlResult, err = v20.F8NewS6InsertBuilder[v20.S6APPUserModel](p7s6DB).
		F8SetEntity(&v20.S6APPUserModel{Id: 44, Name: "dd", Age: 55, Sex: 2}).
		F8OnConflictBuilder().
		F8SetUpdate(v20.F8NewS6Column("Name")).
		F8EXEC(context.Background())
	fmt.Println(sqlResult, err)

	sqlResult, err = v20.F8NewS6InsertBuilder[v20.S6APPUserModel](p7s6DB).
		F8SetEntity(&v20.S6APPUserModel{Id: 44, Name: "dddd", Age: 55, Sex: 2}).
		F8OnConflictBuilder().
		F8SetUpdate(v20.F8NewS6Column("Name")).
		F8EXEC(context.Background())
	fmt.Println(sqlResult, err)
}
