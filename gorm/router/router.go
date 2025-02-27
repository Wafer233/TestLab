package router

import (
	"github.com/gin-gonic/gin"
	"gos/gorm/api"
)

func InitRouter(r *gin.Engine) {
	api.RegisterRouter(r)
}
