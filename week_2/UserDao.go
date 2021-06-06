package main

import (
	sql2 "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func ListTask(task_id string) (interface{}, error)  {
	db, err := sql2.Open("mysql", "root:44805809pjy@tcp(localhost:3306)/geektime")
	if err != nil {
		return nil, err
	}
	name := ""
	sql := "select owner from task where task_id=?"
	err = db.QueryRow(sql,task_id).Scan(&name)
	if err != nil{
		if err == sql2.ErrNoRows {
			return nil, errors.Wrap(err, "查找失败")
		} else {
			return nil, err
		}
	}
	return name, nil
}
