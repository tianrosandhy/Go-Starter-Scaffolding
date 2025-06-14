package dto

import "github.com/go-openapi/strfmt"

type ExampleResponse struct {
	ID        int             `json:"id"`
	Name      string          `json:"name" validate:"required"`
	Price     float64         `json:"price" validate:"required"`
	CreatedAt strfmt.DateTime `json:"created_at" format:"date-time" swaggertype:"string"`
	UpdatedAt strfmt.DateTime `json:"updated_at" format:"date-time" swaggertype:"string"`
}
