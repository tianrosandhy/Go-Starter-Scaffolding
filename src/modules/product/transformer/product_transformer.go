package transformer

import (
	"skeleton/pkg/merger"
	"skeleton/src/modules/product/dto"
	"skeleton/src/modules/product/entity"
)

func TransformSingleProduct(product entity.Product) (resp dto.ProductResponse) {
	merger.Merge(product, &resp)
	return resp
}

func TransformBatchProduct(products []entity.Product) (resp []dto.ProductResponse) {
	merger.Merge(products, &resp)
	return resp
}
