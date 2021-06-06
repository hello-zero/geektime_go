package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func userController(task_id string)  interface{}{

	val, err := listTaskOwnerService(task_id)
	var message string
	result := make(map[string]interface{})
	if err != nil {
		fmt.Printf("%+v\n", err)
		if errors.Is(err, sql.ErrNoRows) {
			message = "数据不存在"
		}else{
			message = err.Error()
		}
		result["code"] = 5000
		result["message"] = message
	}else {
		result["code"] = 200
		result["data"] = val
	}
	return result
}
