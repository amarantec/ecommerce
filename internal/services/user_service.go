package services

import (
	"context"

	"github.com/amarantec/e-commerce/internal/models"
)

func (s Service) Save(ctx context.Context, user models.UserRegister) (models.UserRegister, error) {
	if user.Email == "" {
		return models.UserRegister{}, ErrUserEmailEmpty
	}
	if user.Password == "" {
		return models.UserRegister{}, ErrUserPasswordEmpty
	}

	return s.Repository.Save(ctx, user)
}

func (s Service) Login(ctx context.Context, user models.UserRegister) (int64, error) {
	return s.Repository.ValidateCredentials(ctx, user)
}
