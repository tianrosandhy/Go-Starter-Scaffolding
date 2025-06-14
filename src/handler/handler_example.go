package handler

import (
	"skeleton/pkg/response"
	"skeleton/src/dto"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
)

// Example func
// @Summary Example of single handler / controller.
// @Description Example of single handler / controller.
// @Produce  application/json
// @Success 200 {object} response.Response{data=dto.ExampleResponse} "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /api/v1/example [get]
// @Tags Examples
func (p *Handler) Example(c echo.Context) error {
	now := strfmt.DateTime(time.Now())
	exampleResponse := dto.ExampleResponse{
		ID:        1,
		Name:      "Contoh",
		Price:     2500,
		CreatedAt: now,
		UpdatedAt: now,
	}
	p.App.Log.Printf("Example log : %+v", exampleResponse)
	return response.OK(c, "Handler or controller response example", exampleResponse)
}
