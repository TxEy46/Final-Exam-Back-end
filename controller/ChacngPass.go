package controller

import (
	"go-backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ฟังก์ชันสำหรับการเปลี่ยนรหัสผ่าน
func ChangePassword(router *gin.Engine, db *gorm.DB) {
	routes := router.Group("/auth")
	{
		routes.POST("/change-password", func(c *gin.Context) {
			// ดึงข้อมูล JSON ที่ส่งมาใน request body
			var changeData map[string]string
			if err := c.ShouldBindJSON(&changeData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
				return
			}

			// ค้นหาผู้ใช้จาก email ที่ส่งมา
			var user model.Customer
			result := db.Where("email = ?", changeData["email"]).First(&user)
			if result.Error != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email"})
				return
			}

			// เปรียบเทียบรหัสผ่านเก่าที่ส่งมา
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changeData["old_password"]))
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Old password is incorrect"})
				return
			}

			// เข้ารหัสรหัสผ่านใหม่
			newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(changeData["new_password"]), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash new password"})
				return
			}

			// อัพเดตรหัสผ่านในฐานข้อมูล
			user.Password = string(newPasswordHash)
			if err := db.Save(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update password"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Password changed successfully",
			})
		})
	}
}
