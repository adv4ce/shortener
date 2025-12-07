package handlers

import (
	"context"
	"net/http"
	"shortener/internal/database"
	"shortener/internal/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type Gurl struct {
	Url string `json:"url"`
}

func health() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	}
}

func gCode(r database.DBRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var gurl Gurl

		if err := ctx.ShouldBindJSON(&gurl); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		if !services.IsValidUrl(gurl.Url) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid or unsafe URL"})
			return
		}
		code, err := r.GetCode(ctx.Request.Context(), gurl.Url)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"url":  gurl,
			"code": code,
		})
	}
}

func gURL(db database.DBRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Param("code")

		if code == "" || strings.Contains(code, "/") || strings.Contains(code, ".") {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid short code"})
			return
		}

		url, err := db.GetURL(context.Background(), code)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, url.URL)
	}
}
