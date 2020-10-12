package yorm

import (
	"fmt"
	"github.com/timchine/pole/pkg/yorm/mysql"
	"testing"
	"time"
)

func TestDB_LoadSqlYaml(t *testing.T) {
	db, err := Open(mysql.NewMysql("127.0.0.1", 3306, "root", "123456", "pole"))
	fmt.Println(err)
	if err != nil {

	}
	//user := User{
	//	Id:   false,
	//	Name: "fhj",
	//}
	var u = UserA{
		UserId: 2,
	}
	var uo []UserA
	//users := map[string]User{"1":user, "2": user}
	fmt.Println(db.LoadSqlYaml("./yaml"))
	//fmt.Println(db.AddQuery("user", User{
	//	Id:   true,
	//	Name: "fhj",
	//}).AddQuery("users", users).Exec("user.getUserById").Scan(&u))
	//fmt.Println(db.mp)
	//var id int
	fmt.Println(db.AddQuery("user", u).Page(1, 1, "user.getUserById"))
	fmt.Println(uo)
}

type UserA struct {
	UserName string `yorm:"user_name"`
	UserId int `yorm:"user_id"`
	CreatedAt time.Time
	User
}

type User struct {
	Id *int `yorm:"user_age"`
	Name string
}
