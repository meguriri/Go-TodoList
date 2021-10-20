package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlDb *sql.DB
var MysqlDbErr error
var id int

func InitDB() {
	MysqlDb, MysqlDbErr = sql.Open("mysql", "root:xyy001019@tcp(localhost:3306)/bubble?charset=utf8")
	if MysqlDbErr != nil {
		panic(MysqlDbErr.Error())
	} else {
		fmt.Println("connect success!!")
	}
	//
}

func CloseDB() {
	MysqlDb.Close()
}

func CheckUser(user *User) (bool, error) {
	var nuser User
	row := MysqlDb.QueryRow("SELECT * FROM user WHERE name=?", user.Name)
	if row.Err() != nil {
		return false, row.Err()
	}
	row.Scan(&nuser.Id, &nuser.Name, &nuser.Password)
	id = nuser.Id
	fmt.Println("user:", user.Name)
	fmt.Println("user:", user.Password)
	fmt.Println("db:", nuser.Name)
	fmt.Println("db:", nuser.Password)
	if nuser.Password == user.Password {
		return true, nil
	} else {
		return false, nil
	}
}

func GetList(name string) ([]Todo, int) {
	todos := make([]Todo, 0)
	rows, _ := MysqlDb.Query("SELECT * FROM todos where user_id in (SELECT user_id FROM user where name=?)", name)
	var todo Todo
	for rows.Next() {
		rows.Scan(&todo.Id, &todo.Title, &todo.Status, &todo.UserId)
		todos = append(todos, todo)
	}
	return todos, id
}

func CreatNewTodo(todo Todo) (bool, error) {
	_, err := MysqlDb.Exec("insert into todos(title, status,user_id)values(?,?,?)", todo.Title, todo.Status, todo.UserId)
	if err != nil {
		fmt.Println("exec failed", err)
		return false, err
	} else {
		fmt.Println("exec success!!!")
		return true, nil
	}
}

func UpdateTodo(title string, status string) (bool, error) {
	var newstatus int
	if status == "0" {
		newstatus = 1
	} else {
		newstatus = 0
	}
	_, err := MysqlDb.Exec("update todos set status=? where title=?", newstatus, title)
	if err != nil {
		fmt.Println("update failed", err)
		return false, err
	} else {
		fmt.Println("update success", err)
		return true, nil
	}
}

func DeleteTodo(title string) (bool, error) {
	_, err := MysqlDb.Exec("delete from todos where title=?", title)
	if err != nil {
		fmt.Println("delete failed", err)
		return false, err
	} else {
		fmt.Println("delete success", err)
		return true, nil
	}
}

func IsUserExsit(user *User) bool {
	row := MysqlDb.QueryRow("SELECT count(*) FROM user where name=? limit 1", user.Name)
	var a int
	row.Scan(&a)
	fmt.Println(a)
	if a == 0 {
		return true
	} else if a == 1 {
		return false
	}
	return false
}

func CreatNewUser(user *User) (bool, int64, error) {
	r, err := MysqlDb.Exec("insert into user (name,password)values(?,?)", user.Name, user.Password)
	id, _ := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed", err)
		return false, id, err
	} else {
		fmt.Println("exec success!!!")
		return true, 0, nil
	}
}
