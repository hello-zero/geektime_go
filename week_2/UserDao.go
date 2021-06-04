package week_2

import sql2 "database/sql"

func ListUser() (interface{}, error)  {
	db, err := sql2.Open("mysql", "127.0.0.1")
	if err != nil {
		panic(err)
	}
	sql := "select * from user"
	err := db.QueryRow(sql).Scan("&amp;name")
}
