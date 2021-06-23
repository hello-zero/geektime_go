package biz

import "log"

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepo interface {
	Register(u *User) (*User, error)
	Login(u *User) (string, error)
}

type UserUseCase struct {
	repo UserRepo
	logger log.Logger
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase  {
	return &UserUseCase{
		repo: repo,
		logger: logger,
	}
}

func (uc *UserUseCase) Register(u *User) (*User, error)  {
	return uc.repo.Register(u)
}

func (uc *UserUseCase) Login(u *User) (string, error)  {
	return uc.repo.Login(u)
}
