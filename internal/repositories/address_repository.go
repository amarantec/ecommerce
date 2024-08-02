package repositories

import (
  "github.com/amarantec/e-commerce/internal/models"
  "github.com/jackc/pgx/v5"
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
