package config

import (
	"github.com/NTUT-Database-System-Course/ACW-Backend/docs"
)

func NewSwaggerInfo() {
	docs.SwaggerInfo.Title = "ACW-Backend API"
	docs.SwaggerInfo.Description = "This is a api server for ACW-Backend"
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
