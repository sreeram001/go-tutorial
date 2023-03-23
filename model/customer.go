package model

import (
	"time"
)

type Customer struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
