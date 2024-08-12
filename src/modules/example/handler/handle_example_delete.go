package handler

import (
	"skeleton/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

// DeleteExample func
// @securityDefinitions.basic BearerAuth
// @Summary Delete single example.
// @Description Delete single example.
// @Produce  application/json
// @Param id path int true "Example ID"
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /api/example/v1/delete/{id} [post]
// @Tags Examples
func (p *ExampleHandler) DeleteExample(c echo.Context) error {
	exampleID, _ := strconv.Atoi(c.Param("id"))
	example := p.ExampleRepository.GetByID(c, exampleID)
	if example == nil {
		return response.ErrorMessage(c, "Example not found", 404)
	}

	err := p.ExampleRepository.Delete(c, example.ID)
	if err != nil {
		return response.Error(c, err, 400)
	}

	return response.OK(c, "Example has been deleted successfully")
}
