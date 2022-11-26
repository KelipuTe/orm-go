package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"orm-go/v20"
	"orm-go/v20/middleware"
)

func main() {
	//testSelectFirst()
	//testSelectGetList()
	//testSlowSQL()
	//testInsert()
	//testUpdate()
	//testDelete()
	//testRawSelectFirst()
	//testRawGetList()
	//testRawInsert()
	//testRawUpdate()
	//testRawDelete()
}

func testSelectFirst() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := v20.F8NewS6DB(p7s6SQLDB, v20.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

	sqlResult, err := v20.F8NewS6SelectBuilder[v20.S6APPUserModel](p7s6DB).
		F8Where(v20.F8NewS6Column("Id").F8Equal(11)).F8First(context.Background())
	fmt.Println(sqlResult, err)
}

func testSelectGetList() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := v20.F8NewS6DB(p7s6SQLDB, v20.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

	s5SQLResult, err := v20.F8NewS6SelectBuilder[v20.S6APPUserModel](p7s6DB).
		F8GetList(context.Background())
	fmt.Println(s5SQLResult, err)
	for _, t4value := range s5SQLResult {
		fmt.Println(t4value)
	}
}

func testSlowSQL() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := v20.F8NewS6DB(p7s6SQLDB, v20.F8DBWithMiddleware(
		middleware.SqlLogMiddlewareBuild(),
		middleware.SlowLogMiddlewareBuild(),
		middleware.SlowLogTriggerMiddlewareBuild(),
	))

	sqlResult, err := v20.F8NewS6SelectBuilder[v20.S6APPUserModel](p7s6DB).
		F8Where(v20.F8NewS6Column("Id").F8Equal(11)).F8First(context.Background())
	fmt.Println(sqlResult, err)
}

func testInsert() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := v20.F8NewS6DB(p7s6SQLDB, v20.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

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

func testUpdate() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := v20.F8NewS6DB(p7s6SQLDB, v20.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

	sqlResult, err := v20.F8NewS6UpdateBuilder[v20.S6APPUserModel](p7s6DB).
		F8SetEntity(&v20.S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
		F8SetUpdate(v20.F8NewS6Column("Name")).
		F8Where(v20.F8NewS6Column("Id").F8Equal(11)).
		F8EXEC(context.Background())
	fmt.Println(sqlResult, err)

	sqlResult, err = v20.F8NewS6UpdateBuilder[v20.S6APPUserModel](p7s6DB).
		F8SetEntity(&v20.S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
		F8SetUpdate(v20.F8NewS6Column("Age").ToAssignment(v20.F8NewS6Column("Age").F8Add(1))).
		F8Where(v20.F8NewS6Column("Id").F8Equal(11)).
		F8EXEC(context.Background())
	fmt.Println(sqlResult, err)
}

func testDelete() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := v20.F8NewS6DB(p7s6SQLDB, v20.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

	sqlResult, err := v20.F8NewS6DeleteBuilder[v20.S6APPUserModel](p7s6DB).
		F8Where(v20.F8NewS6Column("Id").F8Equal(55)).
		F8EXEC(context.Background())
	fmt.Println(sqlResult, err)
}

func testRawSelectFirst() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := v20.F8NewS6DB(p7s6SQLDB, v20.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

	sqlResult, err := v20.F8NewS6Raw[v20.S6APPUserModel](
		p7s6DB,
		"SELECT * FROM `app_user`;",
		nil,
	).F8First(context.Background())
	fmt.Println(sqlResult, err)
}

func testRawGetList() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := v20.F8NewS6DB(p7s6SQLDB, v20.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

	s5SQLResult, err := v20.F8NewS6Raw[v20.S6APPUserModel](
		p7s6DB,
		"SELECT * FROM `app_user`;",
		nil,
	).F8GetList(context.Background())
	fmt.Println(s5SQLResult, err)
	for _, t4value := range s5SQLResult {
		fmt.Println(t4value)
	}
}

func testRawInsert() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := v20.F8NewS6DB(p7s6SQLDB, v20.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

	sqlResult, err := v20.F8NewS6Raw[v20.S6APPUserModel](
		p7s6DB,
		"INSERT INTO `app_user`(`id`,`name`,`age`,`sex`) VALUES(?,?,?,?);",
		[]any{55, "ee", 66, 1},
	).F8EXEC(context.Background())
	fmt.Println(sqlResult, err)
}

func testRawUpdate() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := v20.F8NewS6DB(p7s6SQLDB, v20.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

	sqlResult, err := v20.F8NewS6Raw[v20.S6APPUserModel](
		p7s6DB,
		"UPDATE `app_user` SET `name`=? WHERE `id` = ?;",
		[]any{"eeee", 55},
	).F8EXEC(context.Background())
	fmt.Println(sqlResult, err)
}

func testRawDelete() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := v20.F8NewS6DB(p7s6SQLDB, v20.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

	sqlResult, err := v20.F8NewS6Raw[v20.S6APPUserModel](
		p7s6DB,
		"DELETE FROM `app_user` WHERE `id` = ?;",
		[]any{55},
	).F8EXEC(context.Background())
	fmt.Println(sqlResult, err)
}
