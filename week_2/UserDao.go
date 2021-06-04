package main

import (
	sql2 "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// sql2.ErrNoRows的错误含义是没有查询结果可返回， 需要区分是否正常匹配无结果
// 1：查询匹配结果为空，即正常执行查询没有得到匹配条件的结果，此时应该正常返回结果，无需向上层返回error
// 2：如果异常匹配无结果，则要wrap向上层返回

func ListUser() (interface{}, error)  {
	db, err := sql2.Open("mysql", "root:44805809pjy@tcp(localhost:6606)/d_infosec")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(11111)
	var name string
	sql := "select hash from webscan_task_bug where task_id=?"
	err = db.QueryRow(sql,"75bd1062a12811e9b562fa163e6ef7e2").Scan(&name)
	if err != nil{
		if err != sql2.ErrNoRows {
			fmt.Println(123,err)
			return nil, err
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println(err)
	fmt.Println(name)
	//return nil, nil

	rows, err2 := db.Query(sql,"75bd1062a12811e9b562fa163e6ef7e2")
	if err2 != nil{
		if err2 != sql2.ErrNoRows {
			fmt.Println(123,err2)
			return nil, err2
		} else {
			fmt.Println(err2)
		}
	}

	for rows.Next(){
		err3 := rows.Scan(&name)
		fmt.Println(err3)
	}

	return nil, nil
}
