package week_2

func  userController()  interface{}{
	val, err := listUserService()
	result := make(map[string]interface{})
	if err != nil {
		result["code"] = 5000
		result["message"] = err
	}else {
		result["code"] = 200
		result["data"] = val
	}

}
