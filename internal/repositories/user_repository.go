package repositories

import (
	"context"
	"errors"

	"github.com/amarantec/e-commerce/internal/models"
	"github.com/amarantec/e-commerce/internal/utils"
)

func (r *RepositoryPostgres) Save(ctx context.Context, user models.UserRegister) (models.UserRegister, error) {
	hashedPassword, err := utils.HashPassword(string(user.Password))
	if err != nil {
		return models.UserRegister{}, err
	}
	err = r.Conn.QueryRow(
		ctx,
		`INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email`,
		user.Email, hashedPassword).Scan(&user.Id, &user.Email)
	if err != nil {
		return models.UserRegister{}, err
	}
	return user, nil
}

func (r *RepositoryPostgres) ValidateCredentials(ctx context.Context, user models.UserRegister) error {
	var retriviedPassword string

	err := r.Conn.QueryRow(
		ctx,
		`SELECT email, password FROM users WHERE  email = $1 `, user.Email,
	).Scan(&retriviedPassword)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CheckPasswordHash(user.Password, retriviedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
