package models

import (
	"time"
)

type Cost struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Properties interface{} `json:"properties"`
	Total      int         `json:"total"`
	Date       interface{} `json:"date"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt  *time.Time  `json:"deleted_at"`
}
