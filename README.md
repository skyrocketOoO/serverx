# serverx
A server template combines multiple basic tools

## How to use
1. Clone the repository
2. Change the module name to yours

## Convention
```sh
# Check conventions
python3 scripts/convention/check.py
```
- **camelCase**: file name, folder name
- **PascalCase**: json attribute name, db table name, db column name
- Lines must be less than 99 characters, but its soft limit
- Never use **init()**, except for 3rd-party specify
- Never use fmt.Prinxxx, use log.Debug() instead

## Dev environment
- Linux
- Code first (not data first)

## Directory structure
- <span style="color: pink;">api</span>: API definition
  - <span style="color: pink;">grpc</span>: Store the generated grpc object
  - <span style="color: pink;">rest</span>: Route binding with controller function
- <span style="color: pink;">benchmark</span>: Benchmark and api test with k6
- <span style="color: pink;">docs</span>: Documentation
  - <span style="color: pink;">grpc</span>: Document generated from proto-gen-doc
  - <span style="color: pink;">rest</span>: Document generated from go-swagger
- <span style="color: pink;">internal</span>: Go function to not export the main logic
  - <span style="color: pink;">boot</span>: The initial process
  - <span style="color: pink;">cmd</span>: Command to run start and multiple functions
  - <span style="color: pink;">global</span>:
    - <span style="color: pink;">domain</span>: General domain like errorcode...etc
  - <span style="color: pink;">controller</span>: The first layer received from api, used for parse request and response, validation...etc
  - <span style="color: pink;">usecase</span>: Primary business logic
  - <span style="color: pink;">repository</span>: Orm, Redis, Dao...
  - <span style="color: pink;">i18n</span>
- <span style="color: pink;">manifest</span>: Config...
  - <span style="color: pink;">config</span>
  - <span style="color: pink;">deploy</span>
  - <span style="color: pink;">docker</span>
  - <span style="color: pink;">protobuf</span>
- <span style="color: pink;">scripts</span>:
- <span style="color: pink;">utils</span>: Utility functions

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
