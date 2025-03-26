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
	// ส่ง db ไปให้ Customer
	Customer(router, db)
	RegisterCustomer(router, db)
	ChangePassword(router, db)
	AddToCart(router, db)
	SearchProducts(router, db)
	router.Run()
}
