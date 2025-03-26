package controller

import (
	"go-backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ฟังก์ชันสำหรับการค้นหาสินค้าตามรายละเอียดและช่วงราคา
func SearchProducts(router *gin.Engine, db *gorm.DB) {
	routes := router.Group("/product")
	{
		routes.GET("/search", func(c *gin.Context) {
			// รับค่าพารามิเตอร์จาก query string
			description := c.DefaultQuery("description", "")
			minPrice := c.DefaultQuery("min_price", "0")
			maxPrice := c.DefaultQuery("max_price", "1000000")

			// แปลงค่าราคาเป็นตัวเลข
			minPriceFloat, err := strconv.ParseFloat(minPrice, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minimum price"})
				return
			}
			maxPriceFloat, err := strconv.ParseFloat(maxPrice, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maximum price"})
				return
			}

			// สร้าง query เงื่อนไขสำหรับการค้นหาสินค้า
			var products []model.Product
			result := db.Where("description LIKE ? AND price BETWEEN ? AND ?", "%"+description+"%", minPriceFloat, maxPriceFloat).Find(&products)

			if result.Error != nil || len(products) == 0 {
				c.JSON(http.StatusNotFound, gin.H{"message": "No products found within the price range"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message":  "Products found",
				"products": products,
			})
		})
	}
}
