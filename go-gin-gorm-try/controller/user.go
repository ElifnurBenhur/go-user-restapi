package controller

import (
	"github.com/ElifnurBenhur/go-gin-gorm-try/config"
	"github.com/ElifnurBenhur/go-gin-gorm-try/models"
	"github.com/gin-gonic/gin"
)

func UserGETController(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)

}
func UserPOSTController(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}
func UserDELETEController(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)

}
func UserPUTController(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)

}
