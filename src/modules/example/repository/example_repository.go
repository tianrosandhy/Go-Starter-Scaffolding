package repository

import (
	"skeleton/bootstrap"
	"skeleton/pkg/merger"
	"skeleton/src/modules/example/dto"
	"skeleton/src/modules/example/entity"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ExampleRepository struct {
	DB     *gorm.DB
	Redis  *redis.Client
	Config *viper.Viper
	Log    *logrus.Logger
}

func NewExampleRepository(app *bootstrap.Application) *ExampleRepository {
	return &ExampleRepository{
		DB:     app.DB,
		Redis:  app.Redis,
		Config: app.Config,
		Log:    app.Log,
	}
}

func (p *ExampleRepository) GetAll() []entity.Example {
	var examples []entity.Example
	p.DB.Find(&examples)
	return examples
}

func (p *ExampleRepository) GetByID(id int) *entity.Example {
	var example entity.Example
	p.DB.First(&example, id)
	if example.ID == 0 {
		return nil
	}
	return &example
}

func (p *ExampleRepository) Create(exampleRequest dto.ExampleRequest) (*entity.Example, error) {
	example := entity.Example{}
	merger.Merge(exampleRequest, &example)

	err := p.DB.Create(&example).Error
	if err != nil {
		return nil, err
	}
	return &example, nil
}

func (p *ExampleRepository) Update(example entity.Example, exampleRequest dto.ExampleRequest) (*entity.Example, error) {
	id := example.ID
	merger.Merge(exampleRequest, &example)
	example.ID = id

	err := p.DB.Save(&example).Error
	if err != nil {
		return nil, err
	}

	return &example, nil
}

func (p *ExampleRepository) Delete(id int) error {
	return p.DB.Delete(&entity.Example{}, id).Error
}
