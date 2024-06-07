package database

import (
	exampleEntity "skeleton/src/modules/example/entity"
)

var EntityMigrations []interface{} = []interface{}{
	//  specify the entities that you want to migrate here
	&exampleEntity.Example{},
}
