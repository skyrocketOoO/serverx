package boot

import (
	"github.com/skyrocketOoO/serverx/docs/openapi"
	"github.com/skyrocketOoO/serverx/internal/domain"
)

func InitSwagger() {
	openapi.SwaggerInfo.Title = "OpenAPI"
	openapi.SwaggerInfo.Version = domain.ApiVersion
	// openapi.SwaggerInfo.Host = ""
	openapi.SwaggerInfo.BasePath = "/" + domain.ApiVersion
	openapi.SwaggerInfo.Schemes = []string{"http"}
}
