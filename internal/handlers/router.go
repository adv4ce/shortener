package handlers

import (
	"fmt"
	"net/http"
	"shortener/internal/database"
	"github.com/gin-contrib/cors"
	"time"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRouter(con *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(rateLimitMiddleware())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %s - %s %s\n", param.ClientIP, param.Method, param.Path)
	}))
	db := database.InitDBRepo(con)

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000", "https://yourfrontenddomain.com"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	router.GET("/health", health())
	router.POST("/link", gCode(*db))
	router.GET("/:code", gURL(*db))
	return router
}

func rateLimitMiddleware() gin.HandlerFunc {
	limiter := tollbooth.NewLimiter(50.0/60.0, nil)
	
	return func (ctx *gin.Context)  {
		httpError := tollbooth.LimitByRequest(limiter, ctx.Writer, ctx.Request)
		if httpError != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests. Try again later."})
			return
		}
		ctx.Next()
	}
}