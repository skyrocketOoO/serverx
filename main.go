package main

import (
	"github.com/skyrocketOoO/serverx/cmd"
	_ "github.com/skyrocketOoO/serverx/docs/openapi"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
