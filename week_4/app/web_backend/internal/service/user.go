package service

import "geektime_go/week_4/app/web_backend/internal/biz"

type UserService struct {
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService{
	return &UserService{
		uc: uc,
	}
}

func (us *UserService) Register(userName string, passwd string) (interface{}, error)  {
	us.uc.Register(&biz.User{
		Username: userName,
		Password: passwd,
	})
	return 200, nil
}
func (us *UserService) Login(userName string, passwd string) (interface{}, error) {
	us.uc.Login(&biz.User{
		Username: userName,
		Password: passwd,
	})
	return 200, nil
}