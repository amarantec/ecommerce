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

func (r *RepositoryPostgres) ValidateCredentials(ctx context.Context, user models.UserRegister) (int64, error) {
	var retriviedPassword string

	err := r.Conn.QueryRow(
		ctx,
		`SELECT id, password FROM users WHERE  email = $1 `, user.Email,
	).Scan(&user.Id, &retriviedPassword)
	if err != nil {
		return 0, err
	}
	passwordIsValid := utils.CheckPasswordHash(user.Password, retriviedPassword)
	if !passwordIsValid {
		return 0, errors.New("credentials invalid")
	}

	return user.Id, nil
}
