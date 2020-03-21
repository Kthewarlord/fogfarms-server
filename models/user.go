package models

import (
	"time"
)

type Type string

const (
	Administrator Type = "Administrator"
)

type User struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Salt      string    `json:"salt"`
	Hash      string    `json:"hash"`
	Type      Type      `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
