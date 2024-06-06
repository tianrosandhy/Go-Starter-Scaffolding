package repository

import (
	"skeleton/bootstrap"
	"skeleton/pkg/merger"
	"skeleton/src/modules/product/dto"
	"skeleton/src/modules/product/entity"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB     *gorm.DB
	Redis  *redis.Client
	Config *viper.Viper
	Log    *logrus.Logger
}

func NewProductRepository(app *bootstrap.Application) *ProductRepository {
	return &ProductRepository{
		DB:     app.DB,
		Redis:  app.Redis,
		Config: app.Config,
		Log:    app.Log,
	}
}

func (p *ProductRepository) GetAll() []entity.Product {
	var products []entity.Product
	p.DB.Find(&products)
	return products
}

func (p *ProductRepository) GetByID(id int) *entity.Product {
	var product entity.Product
	p.DB.First(&product, id)
	if product.ID == 0 {
		return nil
	}
	return &product
}

func (p *ProductRepository) Create(productRequest dto.ProductRequest) (*entity.Product, error) {
	product := entity.Product{}
	merger.Merge(productRequest, &product)

	err := p.DB.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepository) Update(product entity.Product, productRequest dto.ProductRequest) (*entity.Product, error) {
	id := product.ID
	merger.Merge(productRequest, &product)
	product.ID = id

	err := p.DB.Save(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepository) Delete(id int) error {
	return p.DB.Delete(&entity.Product{}, id).Error
}
