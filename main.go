package main

import (
	"net/http"
	"os"

	"leakybucket/ratelimiter"

	"github.com/gin-gonic/gin"
)

func RateLimiterMiddleware(bucket ratelimiter.UserBucket) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if bucket.AllowAccess(1) {
			ctx.Next()
			return
		}
		ctx.AbortWithStatus(http.StatusTooManyRequests)
	}
}

func main() {
	server := gin.Default()

	user := ratelimiter.CreateUserBucket(1, 100)
	server.Use(RateLimiterMiddleware(user))

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "yoyo")
	})

	if err := server.Run(); err != nil {
		os.Exit(1)
	}
}
