package router

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/member"
    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
)

func NewRouter(e *echo.Echo) {
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Group all member-related routes
    g := e.Group("/api/member")
    {
        g.POST("/register", member.Register)
        g.POST("/login", member.Login)
        g.GET("/info", member.GetInfo, config.JWTMiddleware)
    }
}