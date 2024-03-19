package ratelimit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func rateLimitPerIP(rps int, b int) gin.HandlerFunc {
	var (
		limiter = rate.NewLimiter(rate.Limit(rps), b)
	)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			return
		}
		c.Next()
	}
}

func Main() {
	r := gin.Default()

	// Áp dụng middleware này vào toàn bộ router hoặc bạn có thể áp dụng cho từng route cụ thể
	r.Use(rateLimitPerIP(1, 5)) // Giới hạn 1 request/giây với burst lên đến 5 request

	// Tạo route cho tính năng tạo mới user
	r.POST("/users", func(c *gin.Context) {
		// Logic tạo mới user ở đây
		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	})

	r.Run(":8080")
}
