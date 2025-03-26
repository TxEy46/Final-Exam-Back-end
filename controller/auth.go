package controller

import (
	"go-backend/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ฟังก์ชัน login สำหรับตรวจสอบข้อมูลจาก JSON
func Login(router *gin.Engine, db *gorm.DB) {
	router.POST("/auth/login", func(c *gin.Context) {
		login(c, db)
	})
}

func login(c *gin.Context, db *gorm.DB) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// อ่าน JSON ที่ส่งมา
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}

	// ดึงข้อมูลผู้ใช้จากฐานข้อมูล
	var customer model.Customer
	result := db.Where("email = ? AND password = ?", req.Email, req.Password).First(&customer)

	// ตรวจสอบผลลัพธ์จากการ query
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
		}
		return
	}

	// ถ้าผู้ใช้เข้าสู่ระบบสำเร็จ ส่งข้อมูลลูกค้ากลับไป
	c.JSON(200, gin.H{
		"customer": gin.H{
			"CustomerID":  customer.CustomerID,
			"FirstName":   customer.FirstName,
			"LastName":    customer.LastName,
			"Email":       customer.Email,
			"PhoneNumber": customer.PhoneNumber,
			"Address":     customer.Address,
			"CreateAt":    customer.CreatedAt,
			"UpdateAt":    customer.UpdatedAt,
		},
	})
}
