package repositories

import (
  "github.com/amarantec/e-commerce/internal/models"
  "github.com/jackc/pgx/v5"
  "context"
)

func (r *RepositoryPostgres) InsertAddress(ctx context.Context, address models.Address) (models.Address, error) {
  err := r.Conn.QueryRow(
    ctx,
    `INSERT INTO addresses 
      (user_id, first_name, last_name, street, city, state, zip, country)
      VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`,
     address.UserId,
     address.FirstName,
     address.LastName,
     address.Street,
     address.City,
     address.State,
     address.Zip,
     address.Country).Scan(&address.Id)  
  
  if err != nil {
    return models.Address{}, nil
  }

  return address, nil
}

func (r *RepositoryPostgres) GetAddress (ctx context.Context, id int64) (models.Address, error) {
  var address = models.Address{UserId: id}
  err := r.Conn.QueryRow(
    ctx,
    `SELECT first_name,
      last_name, street,
      city,
      state,
      zip,
      country
      FROM addresses WHERE user_id = $1;`, id).Scan(&address.FirstName,
                                                    &address.LastName,
                                                    &address.Street,
                                                    &address.City,
                                                    &address.State,
                                                    &address.Zip,
                                                    &address.Country)

  if err == pgx.ErrNoRows {
    return models.Address{}, err
  }

  if err != nil {
    return models.Address{}, err
  }

  return address, nil
}

func (r *RepositoryPostgres) UpdateAddress(ctx context.Context, id int64) error {
  var address = models.Address{UserId: id}
  _, err := r.Conn.Exec(
    ctx,
    `UPDATE addresses
      SET first_name = $2,
      last_name = $3,
      street = $4,
      city = $5,
      state = $6,
      zip = $7,
      country = $8
      WHERE user_id = $1;`,
      id,
      address.FirstName,
      address.LastName,
      address.Street,
      address.City,
      address.State,
      address.Zip,
      address.Country)

  if err != nil {
    return err
  }
  
  return nil
}

