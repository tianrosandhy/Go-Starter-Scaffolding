package repository

import (
	"skeleton/bootstrap"
	"skeleton/pkg/merger"
	"skeleton/src/modules/example/dto"
	"skeleton/src/modules/example/entity"

	"github.com/labstack/echo/v4"
)

type ExampleRepository struct {
	App *bootstrap.Application
}

func NewExampleRepository(app *bootstrap.Application) *ExampleRepository {
	return &ExampleRepository{
		App: app,
	}
}

func (p *ExampleRepository) GetAll(c echo.Context) []entity.Example {
	var examples []entity.Example
	p.App.DB.Find(&examples)
	return examples
}

func (p *ExampleRepository) GetByID(c echo.Context, id int) *entity.Example {
	var example entity.Example
	p.App.DB.First(&example, id)
	if example.ID == 0 {
		return nil
	}
	return &example
}

func (p *ExampleRepository) Create(c echo.Context, exampleRequest dto.ExampleRequest) (*entity.Example, error) {
	example := entity.Example{}
	merger.Merge(exampleRequest, &example)

	err := p.App.DB.Create(&example).Error
	if err != nil {
		return nil, err
	}
	return &example, nil
}

func (p *ExampleRepository) Update(c echo.Context, example entity.Example, exampleRequest dto.ExampleRequest) (*entity.Example, error) {
	id := example.ID
	merger.Merge(exampleRequest, &example)
	example.ID = id

	err := p.App.DB.Save(&example).Error
	if err != nil {
		return nil, err
	}

	return &example, nil
}

func (p *ExampleRepository) Delete(c echo.Context, id int) error {
	return p.App.DB.Delete(&entity.Example{}, id).Error
}
