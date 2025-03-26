package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Api is now working",
		})
	})
	Login(router, db)
	router.Run()
}
