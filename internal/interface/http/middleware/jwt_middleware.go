package middleware

import (
    "bytes"
    "context"
    "encoding/json"
    "io"
    "net/http"
    "strings"

    "cpru-api/internal/application/auth" 
    "github.com/labstack/echo/v4"
)

type ContextKey string

const KeyUserID ContextKey = "userID"

type graphQLRequest struct {
    OperationName string `json:"operationName"`
}

// ส่ง AuthService เข้ามาด้วย
func JWTMiddleware(authService *auth.AuthService) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // อ่าน body มาเก็บใน buffer แล้ว parse operationName
            var buf bytes.Buffer
            reader := io.TeeReader(c.Request().Body, &buf)

            var req graphQLRequest
            if err := json.NewDecoder(reader).Decode(&req); err != nil {
                return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid GraphQL request"})
            }
            c.Request().Body = io.NopCloser(&buf)

            // Skip register, login, IntrospectionQuery
            if req.OperationName == "register" ||
                req.OperationName == "login" ||
                req.OperationName == "IntrospectionQuery" { // Allow Docs in Playground
                return next(c)
            }

            // ตรวจ header Authorization
            authHeader := c.Request().Header.Get("Authorization")
            if authHeader == "" {
                return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing Authorization header"})
            }
            parts := strings.SplitN(authHeader, " ", 2)
            if len(parts) != 2 || parts[0] != "Bearer" {
                return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Authorization format"})
            }
            tokenString := parts[1]

            // ใช้ AuthService.ValidateToken แทน parse เอง
            userID, err := authService.ValidateToken(tokenString)
            if err != nil {
                return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
            }

            // ใส่ userID ลง context
            newCtx := context.WithValue(c.Request().Context(), KeyUserID, userID)
            c.SetRequest(c.Request().WithContext(newCtx))
            return next(c)
        }
    }
}
