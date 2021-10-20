package dao

type User struct {
	Id       int    `db:"user_id" json:"user_id"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
}

type Todo struct {
	UserId int    `db:"user_id" json:"user_id"`
	Id     int    `db:"id" json:"id"`
	Title  string `db:"title" json:"title"`
	Status int    `db:"status" json:"status"`
}
