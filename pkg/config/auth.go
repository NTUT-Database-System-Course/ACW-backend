package config

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/golang-jwt/jwt/v4"
)

// JWTMiddleware validates the JWT token
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{
                "error": "Missing token",
            })
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            return c.JSON(http.StatusUnauthorized, map[string]string{
                "error": "Invalid token",
            })
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            return c.JSON(http.StatusUnauthorized, map[string]string{
                "error": "Invalid token claims",
            })
        }

        c.Set("user_id", int(claims["user_id"].(float64)))
        return next(c)
    }
}