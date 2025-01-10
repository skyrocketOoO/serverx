package boot

import "github.com/skyrocketOoO/web-server-template/docs/openapi"

func InitSwagger() {
	openapi.SwaggerInfo.Title = "Swagger API"
	openapi.SwaggerInfo.Version = "1.0"
	openapi.SwaggerInfo.Host = "petstore.swagger.io"
	openapi.SwaggerInfo.BasePath = "/v2"
	openapi.SwaggerInfo.Schemes = []string{"http"}
}
