package router

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/users"
)

func NewRouter(e *echo.Echo) {
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Group all user-related routes under /api/users
    g := e.Group("/api/users")
    {
        g.GET("/info", users.Info)
		g.POST("/register", users.Register)
        // Add more user-related routes here
    }
}