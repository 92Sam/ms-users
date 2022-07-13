package models

import "time"

type Product struct {
	Id          string
	Name        string
	Description string
	Rating      float64
	CreateAt    time.Time
}
