package models

import (
	"time"
)

type Role string

type User struct {
	UserID          int       `json:"user_id"`
	Username        string    `json:"username"`
	IsAdministrator bool      `json:"is_administrator"`
	Hash            string    `json:"-"`
	Salt            string    `json:"-"`
	CreatedAt       time.Time `json:"created_at"`
}
