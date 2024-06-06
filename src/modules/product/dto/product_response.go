package dto

import "github.com/go-openapi/strfmt"

type ProductResponse struct {
	ID        int             `json:"id"`
	Name      string          `json:"name" validate:"required"`
	Price     float64         `json:"price" validate:"required"`
	CreatedAt strfmt.DateTime `json:"created_at"`
	UpdatedAt strfmt.DateTime `json:"updated_at"`
}
