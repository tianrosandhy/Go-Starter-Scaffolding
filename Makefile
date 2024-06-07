.PHONY: swagger

swagger:
	swag init -pd -o ./docs

module:
	bash module_generator.sh $(name)