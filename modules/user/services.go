package user

import (
	"context"
	"server/modules/shared"
	"time"
)

type service struct {
	Repository
	timeout        time.Duration
	sharedServices shared.Service
}

func NewService(repository Repository, sharedServices shared.Service) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
		sharedServices,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, hashErr := s.sharedServices.HashPassword(req.Password)

	if hashErr != nil {
		return hashErr
	}

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	_, err := s.Repository.CreateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil
}
