package entity

import (
	"skeleton/pkg/baseentity"
)

type Example struct {
	baseentity.Base

	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (Example) TableName() string {
	return "examples"
}
