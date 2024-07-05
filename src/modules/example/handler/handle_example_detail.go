package handler

import (
	"skeleton/pkg/response"
	"skeleton/src/modules/example/transformer"
	"strconv"

	"github.com/labstack/echo/v4"
)

// DetailExample func
// @securityDefinitions.basic BearerAuth
// @Summary Get single example.
// @Description Get single example.
// @Produce  application/json
// @Param id path int true "Example ID"
// @Success 200 {object} response.Response{data=dto.ExampleResponse} "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /api/example/v1/detail/{id} [get]
// @Tags Examples
func (p *ExampleHandler) DetailExample(c echo.Context) error {
	exampleID, _ := strconv.Atoi(c.Param("id"))
	example := p.ExampleRepository.GetByID(exampleID)
	if example == nil {
		return response.ErrorMessage(c, "Example not found", 404)
	}

	exampleResponse := transformer.TransformSingleExample(*example)
	return response.OK(c, "Get Single Example", exampleResponse)
}
