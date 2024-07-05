package handler

import (
	"skeleton/pkg/response"
	"skeleton/src/modules/example/dto"
	"skeleton/src/modules/example/transformer"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UpdateExample func
// @securityDefinitions.basic BearerAuth
// @Summary Update existing example data
// @Description Update existing example data
// @Produce  application/json
// @Param data body dto.ExampleRequest true "Example request"
// @Success 200 {object} response.Response{data=dto.ExampleResponse} "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /api/example/v1/update/{id} [post]
// @Tags Examples
func (p *ExampleHandler) UpdateExample(c echo.Context) error {
	request := dto.ExampleRequest{}
	err := c.Bind(&request)
	if err != nil {
		return response.Error(c, err, 400)
	}
	err = p.Validator.Struct(&request)
	if err != nil {
		return response.Error(c, err, 400)
	}

	exampleID, _ := strconv.Atoi(c.Param("id"))
	example := p.ExampleRepository.GetByID(exampleID)
	if example == nil {
		return response.ErrorMessage(c, "Example not found", 404)
	}

	example, err = p.ExampleRepository.Update(*example, request)
	if err != nil {
		return response.Error(c, err, 400)
	}

	exampleResponse := transformer.TransformSingleExample(*example)
	return response.OK(c, "Update Example Data", exampleResponse)
}
