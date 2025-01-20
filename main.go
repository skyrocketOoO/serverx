package main

import (
	"github.com/skyrocketOoO/serverx/cmd"
	_ "github.com/skyrocketOoO/serverx/docs/openapi"
)

// @title           Swagger Example API
// @description     This is a sample server celler server.
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
