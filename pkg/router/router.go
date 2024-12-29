package router

import (
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/auth"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/cart"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/favorite"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/member"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/order"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/product"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/store"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/tag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Group all auth-related routes
	g := e.Group("/api/auth")
	{
		g.POST("/login", auth.Login)
	}

	// Group all member-related routes
	g = e.Group("/api/member")
	{
		g.POST("/register", member.Register)
		g.GET("/info", member.Info, config.JWTMiddleware)
		g.PUT("/update", member.Update, config.JWTMiddleware)
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

	// Group all cart-related routes
	g = e.Group("/api/cart")
	{
		g.POST("/add", cart.Add, config.JWTMiddleware)
		g.GET("/list", cart.List, config.JWTMiddleware)
		g.PUT("/update", cart.Update, config.JWTMiddleware)
		g.DELETE("/delete", cart.Delete, config.JWTMiddleware)
	}

	// Group all store-related routes
	g = e.Group("/api/store")
	{
		g.GET("/info", store.Info, config.JWTMiddleware)
	}

	// Group all order-related routes
	g = e.Group("/api/order")
	{
		g.POST("/create", order.Create, config.JWTMiddleware)
		g.GET("/list", order.List, config.JWTMiddleware)
	}

	// Serve static files
	e.Static("/static", "static")
}
