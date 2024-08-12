package handler

import (
	"skeleton/pkg/response"
	"skeleton/src/modules/example/transformer"

	"github.com/labstack/echo/v4"
)

// ListsExample func
// @securityDefinitions.basic BearerAuth
// @Summary Get list of example.
// @Description Get list of example.
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]dto.ExampleResponse} "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /api/example/v1/lists [get]
// @Tags Examples
func (p *ExampleHandler) ListsExample(c echo.Context) error {
	examples := p.ExampleRepository.GetAll(c)
	exampleResponses := transformer.TransformBatchExample(examples)
	return response.OK(c, "Get Example Lists", exampleResponses)
}
