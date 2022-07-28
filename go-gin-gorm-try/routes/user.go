package routes

import (
	"github.com/ElifnurBenhur/go-gin-gorm-try/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
    
	router.GET("/", controller.UserGETController)
	router.POST("/", controller.UserPOSTController)
	router.DELETE("/:id", controller.UserDELETEController)
	router.PUT("/:id", controller.UserPUTController)
	
}
