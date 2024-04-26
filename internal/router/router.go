package router

import (
	"AiServer/internal/middleware"
	"AiServer/pkg/log"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.LoggerWithWriter(log.Logger.Writer()))

	//设置中间件
	r.Use(gin.RecoveryWithWriter(log.Logger.Writer()))
	r.Use(middleware.Cores)

	//注册路由
	register(r)
	return r
}

func register(r *gin.Engine) {
	jwt := middleware.NewJwtMiddleware()
	r.POST("/login", jwt.LoginHandler)
	r.POST("/logout", jwt.LogoutHandler)
	r.POST("/refresh", jwt.RefreshHandler)

	apiRouter := r.Group("/api")
	apiRouter.Use(jwt.MiddlewareFunc())
}
