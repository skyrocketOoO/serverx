package boot

import (
	"github.com/skyrocketOoO/serverx/docs/openapi"
)

func InitSwagger() {
	openapi.SwaggerInfo.Title = "OpenAPI"
	openapi.SwaggerInfo.Schemes = []string{"http"}
}
