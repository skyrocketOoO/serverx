package boot

import docs "github.com/skyrocketOoO/web-server-template/docs/rest"

func InitSwagger() {
	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
