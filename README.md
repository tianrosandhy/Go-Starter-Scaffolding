# Skeleton Go Basic Scaffolding

This is a general basic golang scaffolding as starter project.

### Requirement
- Go 1.22

### Installation
```
cp .env.example .env
go mod tidy
go run .
```

### Generate Swagger Docs
```
make swagger

or 

swag init -pd
```

Docs can be accessed in your `/swag/docs` endpoint by default


### Notes 
- Port & Endpoint are defined in config. For example if you setup the port=9000 and endpoint="/api/v1", then your Base URL will be : http://localhost:9000/api/v1
- By default use gorm, so you can already use `postgre` , `mysql` or `sqlite` as DB_DRIVER
- Default Basic Authentication for all endpoints can be defined from .env `BASIC_AUTH=yourusername:yourpassword` or set blank to disable authentication.
- Default Basic Authentication for swagger can be defined from .env `SWAGGER_AUTH=yourswaggerusername:yourswaggerpassword` or set blank to disable authentication.
- Application modules stored in `./src/modules/{package_name}`.


### Module Autogenerate
- Run with bash script `bash module_generator.sh {modulename}` or run with make file : `make module name={modulename}`
- Register the new module in routes : `./src/routes/routes.go` add module registration in Handler function : `modulename.NewModuleNameModuleRegistration(app, api)`
- Optionally, register the entity migration or seeder to `./src/database` too (optional)
- Module scaffold finish, you can start modify the module based on your needs



### Base Directory Structure
- `./bootstrap` contains all application global scoped package that can be injected to any package
- `./config` contains application configuration key & default value that will be overriden with `.env`
- `./docs` contains swagger autogenerated documentation files
- `./pkg` contains standalone package that will be used in any module (we really discourage use `lib` or `utils`)
- `./src/database` contains migration & seeder handling
- `./src/modules` contains application main business module
    - `@modules/dto` contains struct data object that will be displayed / passed
    - `@modules/entity` contains module single entity / table representative
    - `@modules/handler` contains the business logic for each endpoints. all business logic will be handled here
    - `@modules/repository` contains all available methods for current module. all database logic will be handled here
    - `@modules/transformer` contains logic to transform data from dto to entity or vice versa
- `./src/routes` contains routing and middleware logic that connect to modules