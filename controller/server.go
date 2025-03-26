package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ฟังก์ชันเริ่มต้นเซิร์ฟเวอร์
func StartServer(db *gorm.DB) {
	router := gin.Default()

	// กำหนดเส้นทางการเข้าถึง root ("/")
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

	// กำหนดเส้นทางสำหรับการดูรถเข็นทั้งหมด
	router.GET("/carts/:customer_id", func(c *gin.Context) {
		GetAllCarts(c, db) // เรียกฟังก์ชัน GetAllCarts และส่ง c เป็น *gin.Context
	})

	// เริ่มต้นเซิร์ฟเวอร์
	router.Run()
}
