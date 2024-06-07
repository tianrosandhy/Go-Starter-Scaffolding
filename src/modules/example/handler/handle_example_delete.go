package handler

import (
	"skeleton/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// DeleteSingleExample func
// @securityDefinitions.basic BearerAuth
// @Summary Get single example.
// @Description Get single example.
// @Produce  application/json
// @Param id path int true "Example ID"
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /v1/example/{id} [delete]
// @Tags Examples
func (p *ExampleHandler) DeleteSingleExample(c *fiber.Ctx) error {
	exampleID, _ := strconv.Atoi(c.Params("id"))
	example := p.ExampleRepository.GetByID(exampleID)
	if example == nil {
		return response.ErrorMessage(c, "Example not found", 404)
	}

	err := p.ExampleRepository.Delete(example.ID)
	if err != nil {
		return response.Error(c, err, 400)
	}

	return response.OK(c, "Example has been deleted successfully")
}
