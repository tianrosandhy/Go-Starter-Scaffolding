package handler

import (
	"skeleton/pkg/response"
	"skeleton/src/modules/product/dto"
	"skeleton/src/modules/product/transformer"

	"github.com/gofiber/fiber/v2"
)

// PostCreateProduct func
// @securityDefinitions.basic BearerAuth
// @Summary Create new product data.
// @Description Create new product data.
// @Produce  application/json
// @Param data body dto.ProductRequest true "Product request"
// @Success 200 {object} response.Response{data=dto.ProductResponse} "success"
// @Failure 500 {object} response.Response "internal error"
// @Router /v1/product [post]
// @Tags Products
func (p *ProductHandler) PostCreateProduct(c *fiber.Ctx) error {
	request := dto.ProductRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return response.Error(c, err, 400)
	}
	err = p.Validator.Struct(&request)
	if err != nil {
		return response.Error(c, err, 400)
	}

	product, err := p.ProductRepository.Create(request)
	if err != nil {
		return response.Error(c, err, 400)
	}

	productResponse := transformer.TransformSingleProduct(*product)
	return response.OK(c, "Create Single Product", productResponse)
}
