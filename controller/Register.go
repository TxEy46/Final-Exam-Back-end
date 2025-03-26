package controller

import (
	"go-backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ฟังก์ชันสำหรับการลงทะเบียนผู้ใช้ใหม่
func RegisterCustomer(router *gin.Engine, db *gorm.DB) {
	routes := router.Group("/auth/register")
	{
		routes.POST("", func(c *gin.Context) {
			var newCustomer model.Customer
			if err := c.ShouldBindJSON(&newCustomer); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
				return
			}

			// เข้ารหัสรหัสผ่านก่อนบันทึก
			if err := newCustomer.HashPassword(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
				return
			}

			// บันทึกข้อมูลผู้ใช้ลงฐานข้อมูล
			result := db.Create(&newCustomer)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
				return
			}

			// ส่งข้อมูลการลงทะเบียนสำเร็จ
			c.JSON(http.StatusOK, gin.H{
				"message": "User registered successfully",
				"user":    newCustomer,
			})
		})
	}
}
