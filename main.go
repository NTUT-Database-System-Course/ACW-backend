// main.go

package main

import (
    "log"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    echoSwagger "github.com/swaggo/echo-swagger"
    _ "github.com/NTUT-Database-System-Course/ACW-Backend/docs" // Import generated docs
    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router"
)

// @contact.name   API Support
// @contact.url    https://github.com/NTUT-Database-System-Course/ACW-Backend/issues
// @contact.email  ericncnl3742@gmail.com
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @title          ACW-Backend API
// @version        0.0.1
// @description    This is an API server for ACW-Backend
// @host           localhost:8080
// @BasePath       /v2
func main() {
    // Initialize Swagger Info
    config.NewSwaggerInfo()

    // Initialize the database connection
    config.InitDB()

    // Create a new Echo instance
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
    router.NewRouter(e)

    // Swagger endpoint
    e.GET("/swagger/*", echoSwagger.WrapHandler)

    // Start the server
    if err := e.Start("0.0.0.0:8080"); err != nil {
        log.Fatalf("Shutting down the server with error: %v", err)
    }
}