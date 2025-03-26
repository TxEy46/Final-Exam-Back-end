package controller

import (
	"go-backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ฟังก์ชันสำหรับดูรถเข็นทั้งหมดของลูกค้า
func GetAllCarts(c *gin.Context, db *gorm.DB) {
	// รับข้อมูล customer_id จาก request
	customerID := c.Param("customer_id")

	// ค้นหารถเข็นทั้งหมดของลูกค้า
	var carts []model.Cart
	result := db.Where("customer_id = ?", customerID).Find(&carts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch carts"})
		return
	}

	// สร้างตัวแปรสำหรับเก็บข้อมูลรถเข็นทั้งหมด
	var cartsData []map[string]interface{}

	// ดึงข้อมูลสินค้าจากแต่ละรถเข็น
	for _, cart := range carts {
		var cartItems []model.CartItem
		db.Where("cart_id = ?", cart.CartID).Find(&cartItems)

		var itemsData []map[string]interface{}
		for _, item := range cartItems {
			// ค้นหาข้อมูลสินค้าจาก product_id
			var product model.Product
			db.Where("product_id = ?", item.ProductID).First(&product)

			// แปลงราคาเป็น float64 ถ้าราคาเป็น string
			price, err := strconv.ParseFloat(product.Price, 64)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid price format"})
				return
			}

			// คำนวณราคา
			totalPrice := price * float64(item.Quantity)

			// เก็บข้อมูลสินค้าลงในรายการ
			itemsData = append(itemsData, map[string]interface{}{
				"product_name": product.ProductName,
				"quantity":     item.Quantity,
				"price":        price,
				"total_price":  totalPrice,
			})
		}

		// เก็บข้อมูลรถเข็นแต่ละคัน
		cartsData = append(cartsData, map[string]interface{}{
			"cart_id":   cart.CartID,
			"cart_name": cart.CartName,
			"items":     itemsData,
		})
	}

	// ส่งข้อมูลรถเข็นทั้งหมดกลับไปยัง client
	c.JSON(http.StatusOK, gin.H{
		"customer_id": customerID,
		"carts":       cartsData,
	})
}
