package models

import (
	"time"
)

type Role string

const (
	Administrator Role = "Administrator"
	AuthorizedUser Role = "AuthorizedUser"
)

type User struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Salt      string    `json:"salt"`
	Hash      string    `json:"hash"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
