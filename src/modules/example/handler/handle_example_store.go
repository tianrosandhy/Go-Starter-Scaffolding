package handler

import (
	"skeleton/pkg/response"
	"skeleton/src/modules/example/dto"
	"skeleton/src/modules/example/transformer"

	"github.com/labstack/echo/v4"
)

// StoreExample func
// @securityDefinitions.basic BearerAuth
// @Summary Store new example data.
// @Description Store new example data.
// @Produce  application/json
// @Param data body dto.ExampleRequest true "Example request"
// @Success 200 {object} response.Response{data=dto.ExampleResponse} "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /api/example/v1/store [post]
// @Tags Examples
func (p *ExampleHandler) StoreExample(c echo.Context) error {
	request := dto.ExampleRequest{}
	err := c.Bind(&request)
	if err != nil {
		return response.Error(c, err, 400)
	}
	err = p.Validator.Struct(&request)
	if err != nil {
		return response.Error(c, err, 400)
	}

	example, err := p.ExampleRepository.Create(request)
	if err != nil {
		return response.Error(c, err, 400)
	}

	exampleResponse := transformer.TransformSingleExample(*example)
	return response.OK(c, "Create Single Example", exampleResponse)
}
