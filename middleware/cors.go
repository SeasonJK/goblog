package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc{
	return cors.New(
		cors.Config{
			AllowAllOrigins: 	true,
			AllowMethods: 		[]string{"POST","GET","DELETE","PUT","OPTIONS"},
			AllowHeaders: 		[]string{"*"},
			ExposeHeaders: 		[]string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
			AllowCredentials: 	true,
			MaxAge: 			12 * time.Hour,
		})
}
