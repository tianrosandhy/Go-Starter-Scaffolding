package handler

import (
	"skeleton/pkg/response"
	"skeleton/src/modules/product/transformer"

	"github.com/gofiber/fiber/v2"
)

// GetProductHandler func
// @securityDefinitions.basic BearerAuth
// @Summary Get list of product.
// @Description Get list of product.
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]dto.ProductResponse} "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /v1/product [get]
// @Tags Products
func (p *ProductHandler) GetProductHandler(c *fiber.Ctx) error {
	products := p.ProductRepository.GetAll()
	productResponses := transformer.TransformBatchProduct(products)
	return response.OK(c, "Get Product Lists", productResponses)
}
