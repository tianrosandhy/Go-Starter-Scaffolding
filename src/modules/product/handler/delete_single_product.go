package handler

import (
	"skeleton/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// DeleteSingleProduct func
// @securityDefinitions.basic BearerAuth
// @Summary Get single product.
// @Description Get single product.
// @Produce  application/json
// @Param id path int true "Product ID"
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /v1/product/{id} [delete]
// @Tags Products
func (p *ProductHandler) DeleteSingleProduct(c *fiber.Ctx) error {
	productID, _ := strconv.Atoi(c.Params("id"))
	product := p.ProductRepository.GetByID(productID)
	if product == nil {
		return response.ErrorMessage(c, "Product not found", 404)
	}

	err := p.ProductRepository.Delete(product.ID)
	if err != nil {
		return response.Error(c, err, 400)
	}

	return response.OK(c, "Product has been deleted successfully")
}
