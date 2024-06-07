package handler

import (
	"skeleton/pkg/response"
	"skeleton/src/modules/example/transformer"

	"github.com/gofiber/fiber/v2"
)

// GetExampleHandler func
// @securityDefinitions.basic BearerAuth
// @Summary Get list of example.
// @Description Get list of example.
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]dto.ExampleResponse} "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /v1/example [get]
// @Tags Examples
func (p *ExampleHandler) GetExampleHandler(c *fiber.Ctx) error {
	examples := p.ExampleRepository.GetAll()
	exampleResponses := transformer.TransformBatchExample(examples)
	return response.OK(c, "Get Example Lists", exampleResponses)
}
