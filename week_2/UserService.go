package main

func listUserService() (interface{}, error)  {
	// 查询之前的逻辑
	_, _ = ListUser()
	// 查询之后的逻辑
	return nil, nil
}
