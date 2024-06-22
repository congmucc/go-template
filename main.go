package main

import "gotemplate/cmd"

// @BasePath /api/v1

// @title go-template API
// @version 1.0
// @description This is an enterprise-level Gin template.
// @termsOfService http://swagger.io/terms/
// @host petstore.swagger.io:8088
// @BasePath /v1
func main() {
	defer cmd.Clear()
	cmd.Start()
}
