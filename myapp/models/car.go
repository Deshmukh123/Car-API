package models

import "time"

type Car struct {
	Name      string    `json:"name" bson:"name"`
	Model     string    `json:"model" bson:"model"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
