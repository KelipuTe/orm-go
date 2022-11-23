package v20_demo

import v20 "orm-go/v20"

func main() {
	p7s6DB := v20.F8NewS6DB(nil)
	v20.F8NewS6Insert[v20.S6APPUserModel](p7s6DB).
		F8SetEntity(
			&v20.S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1},
			&v20.S6APPUserModel{Id: 22, Name: "bb", Age: 33, Sex: 2},
		)

	v20.F8NewS6Insert[v20.S6APPUserModel](p7s6DB).
		F8SetEntity(&v20.S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
		F8OnConflictBuilder().
		F8SetUpdate(v20.F8NewS6Column("Id"), v20.F8NewS6Column("s6Column"))
}
