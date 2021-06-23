package data

import (
	"geektime_go/week_4/app/web_backend/internal/biz"
	"log"
)

type userRepo struct {
	data *Data
	logger *log.Logger
}

func NewUserRepo(data *Data, logger *log.Logger) biz.UserRepo  {
	return &userRepo{
		data: data,
		logger: logger,
	}
}

func (up *userRepo) Login(u *biz.User) (string, error)  {
	return "", nil
}

func (up *userRepo) Register(user *biz.User) (*biz.User, error)  {
	return user, nil
}
