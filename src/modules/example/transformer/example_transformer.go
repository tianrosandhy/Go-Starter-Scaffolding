package transformer

import (
	"skeleton/pkg/merger"
	"skeleton/src/modules/example/dto"
	"skeleton/src/modules/example/entity"
)

func TransformSingleExample(example entity.Example) (resp dto.ExampleResponse) {
	merger.Merge(example, &resp)
	return resp
}

func TransformBatchExample(examples []entity.Example) (resp []dto.ExampleResponse) {
	merger.Merge(examples, &resp)
	return resp
}
