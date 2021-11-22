package models

import (
	"time"
)

type Account struct {
	ID        string `json:"id"`
	Name      string `json:"name" validate:"required"`
	Balance   int `json:"balance"`
	AccountNo int `json:"account_no`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}


type AccountFetchParam struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AccountNo string `json:"account_no`
}