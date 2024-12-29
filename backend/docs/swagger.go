// Package docs 提供API文档
package docs

import "github.com/swaggo/swag"

// @title IRT Exam System API
// @version 1.0
// @description API documentation for IRT Exam System
// @BasePath /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func init() {
	swag.Register(swag.Name, &swag.Spec{
		InfoInstanceName: "swagger",
		SwaggerTemplate:  docTemplate,
	})
}
