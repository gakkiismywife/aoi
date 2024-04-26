package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cores(c *gin.Context) {
	//if c.Request.Header.Get("origin") != "" {
	//	c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("origin"))
	//}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Referer-Policy", "strict-origin-when-cross-origin")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")
	c.Set("content-type", "application/json")

	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, "ok")
	}
	c.Next()
}
