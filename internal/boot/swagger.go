package boot

import (
	"github.com/skyrocketOoO/serverx/docs/openapi"
	"github.com/skyrocketOoO/serverx/internal/global"
)

func InitSwagger() {
	openapi.SwaggerInfo.Title = "OpenAPI"
	openapi.SwaggerInfo.Version = global.ApiVersion
	// openapi.SwaggerInfo.Host = ""
	openapi.SwaggerInfo.BasePath = "/" + global.ApiVersion
	openapi.SwaggerInfo.Schemes = []string{"http"}
}
