package middleware

import (
	"AiServer/internal/models"
	"AiServer/pkg/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginUser struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var r LoginRequest
	if err := c.ShouldBind(&r); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	var u models.User
	err := utils.DB.Where("username", r.Username).First(&u).Error
	if err != nil || u.Password != r.Password {
		return nil, jwt.ErrFailedAuthentication
	}
	return &u, nil
}

func Authorizator(data interface{}, c *gin.Context) bool {
	_, ok := data.(*LoginUser)
	return ok
}

func Payload(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.User); ok {
		logrus.Infof("[PayloadFunc]%v", v)
		return jwt.MapClaims{
			"id":   v.Id,
			"name": v.Username,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)

	if claims["id"] == nil || claims["name"] == nil {
		return nil
	}
	id := int64(claims["id"].(float64))
	user := &LoginUser{Id: id, Username: claims["name"].(string)}

	c.Set("user", user)
	return user
}

func NewJwtMiddleware() *jwt.GinJWTMiddleware {
	m, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:           []byte("AI_TRAIN"),
		Timeout:       time.Hour * 24 * 7,
		MaxRefresh:    time.Hour * 24 * 7,
		Authenticator: Authenticator,
		Authorizator:  Authorizator,
		PayloadFunc:   Payload,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  message,
			})
		},
		IdentityHandler: IdentityHandler,
	})
	return m
}
