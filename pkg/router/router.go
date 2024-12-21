package router

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/member"
    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/tag"
    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/product"
    "github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/favorite"
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
        g.GET("/info", member.Info, config.JWTMiddleware)
    }

    // Group all tag-related routes
    g = e.Group("/api/tag")
    {
        g.GET("/list", tag.List)
    }

    // Group all product-related routes
    g = e.Group("/api/product")
    {
        g.POST("/create", product.Create, config.JWTMiddleware)
        g.GET("/list", product.List)
        g.PUT("/update", product.Update, config.JWTMiddleware)
    }

    // group all favorite-related routes
    g = e.Group("/api/favorite")
    {
        g.POST("/add", favorite.Add, config.JWTMiddleware)
        g.GET("/list", favorite.List, config.JWTMiddleware)
        g.DELETE("/delete", favorite.Delete, config.JWTMiddleware)
    }
}