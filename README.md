# web-server-template
The basic server template

## Term
- Init: Initialize someting
- New: New instance

## Dev environment
- Linux
- Code first (not data first)

## Directory structure
- api: API definition
  - grpc: Store the generated grpc object
  - rest: Route binding with controller function
- benchmark: Benchmark and api test with k6
- docs: Documentation
  - grpc: Document generated from proto-gen-doc
  - rest: Document generated from go-swagger
- internal: Go function to not export the main logic
  - boot: The initial process
  - cmd: Command to run start and multiple functions
  - global:
    - domain: General domain like errorcode...etc
  - controller: The first layer received from api, used for parse request and response, validation...etc
  - usecase: Primary business logic
  - repository: Orm, Redis, Dao... 
  - i18n
- manifest: Config...
  - config
  - deploy
  - docker
  - protobuf
- scripts: 
- utils: Utility functions

## Function
- Protocol
  - Grpc
  - Rest

## Package
- Router: Gin
- Unit test: Testify
- Config: Viper
- ORM: Gorm
- Logging: Zerolog
- Cmd: Cobra
- Error trace: Eris
- Rest api doc: swaggo/swag
- Grpc doc: pseudomuto/protoc-gen-doc

## Tool
- DB:
    - Postgres
    - Sqlite
- Benchmark: K6
- format
  - https://github.com/mvdan/gofumpt
- Long line format
  - https://github.com/segmentio/golines

#### Multi-repo vs Mono-repo
https://goframe.org/pages/viewpage.action?pageId=87246750