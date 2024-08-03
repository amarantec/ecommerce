package services

import (
  "context"
  "github.com/amarantec/e-commerce/internal/models"
)

func (s Service) InsertAddress(ctx context.Context, address models.Address) (models.Address, error) {
  if address.UserId == 0 {
    return models.Address{}, ErrAddressUserIdEmpty
  }
  if address.FirstName == "" {
    return models.Address{}, ErrAddressFirstNameEmpty
  }
  if address.LastName == "" {
    return models.Address{}, ErrAddressLastNameEmpty
  }
  if address.Street == "" {
    return models.Address{}, ErrAddressStreetEmpty
  }
  if address.City == "" {
    return models.Address{}, ErrAddressCityEmpty
  }
  if address.State == "" {
    return models.Address{}, ErrAddressStateEmpty
  }
  if address.Zip == "" {
    return models.Address{}, ErrAddressZipEmpty
  }
  if address.Country == "" {
    return models.Address{}, ErrAddressCountryEmpty
  }

  return s.Repository.InsertAddress(ctx, address)
}

func (s Service) GetAddress(ctx context.Context, id int64) (models.Address, error) {
  return s.Repository.GetAddress(ctx, id)
}

func (s Service) UpdateAddress(ctx context.Context, id int64) error {
 err := s.Repository.UpdateAddress(ctx, id)

 if err != nil {
   return err
 }

 return nil
}
