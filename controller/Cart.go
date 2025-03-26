package controller

import (
	"go-backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ฟังก์ชันสำหรับการเพิ่มสินค้าลงในรถเข็น
func AddToCart(router *gin.Engine, db *gorm.DB) {
	routes := router.Group("/cart")
	{
		routes.POST("/add", func(c *gin.Context) {
			// ดึงข้อมูลจาก JSON request body
			var request struct {
				CustomerID int    `json:"customer_id"`
				CartName   string `json:"cart_name"`
				ProductID  int    `json:"product_id"`
				Quantity   int    `json:"quantity"`
			}
			if err := c.ShouldBindJSON(&request); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
				return
			}

			// 1. ค้นหารถเข็นที่มีชื่อที่ลูกค้ากำหนด
			var cart model.Cart
			result := db.Where("customer_id = ? AND cart_name = ?", request.CustomerID, request.CartName).First(&cart)

			// 2. ถ้ารถเข็นไม่พบ ให้สร้างรถเข็นใหม่
			if result.Error != nil {
				cart = model.Cart{
					CustomerID: request.CustomerID,
					CartName:   request.CartName,
				}
				if err := db.Create(&cart).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
					return
				}
			}

			// 3. ค้นหาว่ามีสินค้านี้ในรถเข็นแล้วหรือไม่
			var cartItem model.CartItem
			result = db.Where("cart_id = ? AND product_id = ?", cart.CartID, request.ProductID).First(&cartItem)

			// 4. ถ้ามีสินค้าแล้วให้เพิ่มจำนวนสินค้า
			if result.Error == nil {
				cartItem.Quantity += request.Quantity
				if err := db.Save(&cartItem).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item quantity"})
					return
				}
			} else {
				// 5. ถ้าไม่มีสินค้าในรถเข็น ให้เพิ่มสินค้าใหม่
				cartItem = model.CartItem{
					CartID:    cart.CartID,
					ProductID: request.ProductID,
					Quantity:  request.Quantity,
				}
				if err := db.Create(&cartItem).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product to cart"})
					return
				}
			}

			c.JSON(http.StatusOK, gin.H{"message": "Product added to cart successfully"})
		})
	}
}
