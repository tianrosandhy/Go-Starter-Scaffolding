package entity

import (
	"skeleton/pkg/baseentity"
)

type Product struct {
	baseentity.Base

	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (Product) TableName() string {
	return "products"
}
