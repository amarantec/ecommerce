package models

type Address struct {
  Id        int64   `json:"id"`
  UserId    int64   `json:"user_id"`
  FirstName string  `json:"first_name"`
  LastName  string  `json:"last_name"`
  Street    string  `json:"street"`
  City      string  `json:"city"`
  State     string  `json:"state"`
  Zip       string  `json:"zip"`
  Country   string  `json:"country"`
}

