package handler

import (
	"skeleton/pkg/response"
	"skeleton/src/modules/product/transformer"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetSingleProductHandler func
// @securityDefinitions.basic BearerAuth
// @Summary Get single product.
// @Description Get single product.
// @Produce  application/json
// @Param id path int true "Product ID"
// @Success 200 {object} response.Response{data=dto.ProductResponse} "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /v1/product/{id} [get]
// @Tags Products
func (p *ProductHandler) GetSingleProductHandler(c *fiber.Ctx) error {
	productID, _ := strconv.Atoi(c.Params("id"))
	product := p.ProductRepository.GetByID(productID)
	if product == nil {
		return response.ErrorMessage(c, "Product not found", 404)
	}

	productResponse := transformer.TransformSingleProduct(*product)
	return response.OK(c, "Get Single Product", productResponse)
}
