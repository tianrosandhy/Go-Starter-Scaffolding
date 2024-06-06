package database

import (
	productEntity "skeleton/src/modules/product/entity"
)

var EntityMigrations []interface{} = []interface{}{
	//  specify the entities that you want to migrate here
	&productEntity.Product{},
}
