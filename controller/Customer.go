package controller

import (
	"go-backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ฟังก์ชัน login สำหรับตรวจสอบข้อมูลจาก JSON
func Customer(router *gin.Engine, db *gorm.DB) {
	routes := router.Group("/auth/login")
	{
		routes.POST("", func(c *gin.Context) {
			// ดึงข้อมูล JSON ที่ส่งมาใน request body
			var loginData map[string]string
			if err := c.ShouldBindJSON(&loginData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
				return
			}

			// ค้นหาผู้ใช้จากฐานข้อมูล
			var user model.Customer
			result := db.Where("email = ?", loginData["email"]).First(&user)
			if result.Error != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
				return
			}

			// เปรียบเทียบรหัสผ่านที่เข้ารหัส
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData["password"]))
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
				return
			}

			// หากล็อกอินสำเร็จ ให้ส่งข้อมูลผู้ใช้โดยไม่มีรหัสผ่าน
			user.Password = "" // ลบรหัสผ่านออกก่อนส่ง
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
				"user":    user, // คืนข้อมูลผู้ใช้ (ไม่มีรหัสผ่าน)
			})
		})
	}
}
