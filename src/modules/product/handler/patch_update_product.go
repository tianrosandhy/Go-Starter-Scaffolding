package handler

import (
	"skeleton/pkg/response"
	"skeleton/src/modules/product/dto"
	"skeleton/src/modules/product/transformer"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// PatchUpdateProduct func
// @securityDefinitions.basic BearerAuth
// @Summary Update existing product data
// @Description Update existing product data
// @Produce  application/json
// @Param data body dto.ProductRequest true "Product request"
// @Success 200 {object} response.Response{data=dto.ProductResponse} "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /v1/product/{id} [patch]
// @Tags Products
func (p *ProductHandler) PatchUpdateProduct(c *fiber.Ctx) error {
	request := dto.ProductRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return response.Error(c, err, 400)
	}
	err = p.Validator.Struct(&request)
	if err != nil {
		return response.Error(c, err, 400)
	}

	productID, _ := strconv.Atoi(c.Params("id"))
	product := p.ProductRepository.GetByID(productID)
	if product == nil {
		return response.ErrorMessage(c, "Product not found", 404)
	}

	product, err = p.ProductRepository.Update(*product, request)
	if err != nil {
		return response.Error(c, err, 400)
	}

	productResponse := transformer.TransformSingleProduct(*product)
	return response.OK(c, "Update Product Data", productResponse)
}
