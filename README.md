# web-server-template
The basic server template

## Term
- Init: Initialize someting
- New: New instance

## Dev environment
- Linux
- Code first (not data first)

## Directory structure
- <span style="color: yellow;">api</span>: API definition
  - <span style="color: yellow;">grpc</span>: Store the generated grpc object
  - <span style="color: yellow;">rest</span>: Route binding with controller function
- <span style="color: yellow;">benchmark</span>: Benchmark and api test with k6
- <span style="color: yellow;">docs</span>: Documentation
  - <span style="color: yellow;">grpc</span>: Document generated from proto-gen-doc
  - <span style="color: yellow;">rest</span>: Document generated from go-swagger
- <span style="color: yellow;">internal</span>: Go function to not export the main logic
  - <span style="color: yellow;">boot</span>: The initial process
  - <span style="color: yellow;">cmd</span>: Command to run start and multiple functions
  - <span style="color: yellow;">global</span>:
    - <span style="color: yellow;">domain</span>: General domain like errorcode...etc
  - <span style="color: yellow;">controller</span>: The first layer received from api, used for parse request and response, validation...etc
  - <span style="color: yellow;">usecase</span>: Primary business logic
  - <span style="color: yellow;">repository</span>: Orm, Redis, Dao...
  - <span style="color: yellow;">i18n</span>
- <span style="color: yellow;">manifest</span>: Config...
  - <span style="color: yellow;">config</span>
  - <span style="color: yellow;">deploy</span>
  - <span style="color: yellow;">docker</span>
  - <span style="color: yellow;">protobuf</span>
- <span style="color: yellow;">scripts</span>:
- <span style="color: yellow;">utils</span>: Utility functions

## Integration
- Router: Gin
- Config: Viper
- ORM: Gorm
- Logging: Zerolog
- Cmd: Cobra
- Error trace: erx
- Rest api doc: swaggo/swag
- Loger: Loki x grafana
- Session: JWT
- Validator: go-playground/validator
- DB:
    - Postgres
    - Sqlite
    - Mysql
- Benchmark: K6
- format
  - https://github.com/mvdan/gofumpt
- Long line format
  - https://github.com/segmentio/golines
