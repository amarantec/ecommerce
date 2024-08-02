package repositories

import (
  "github.com/amarantec/e-commerce/internal/models"
  "github.com/jackc/pgx/v5"
)

func (r *RepositoryPostgres) InsertAddress(ctx context.Context, address models.Address) (models.Address, error) {
  err := r.Conn.QueryRow(
    ctx,
    `INSERT INTO addresses 
