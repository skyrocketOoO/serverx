package boot

import (
	"github.com/skyrocketOoO/web-server-template/docs/openapi"
	"github.com/skyrocketOoO/web-server-template/internal/global"
)

func InitSwagger() {
	openapi.SwaggerInfo.Title = "Swagger API"
	openapi.SwaggerInfo.Version = "1.0"
	openapi.SwaggerInfo.Host = ""
	openapi.SwaggerInfo.BasePath = global.ApiVersion
	openapi.SwaggerInfo.Schemes = []string{"http"}
}
