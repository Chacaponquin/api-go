package user

import "context"

type User struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"email"`
}

type CreateUserReq struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"email"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
}

type Service interface {
	CreateUser(c context.Context, req *CreateUserReq) error
}
