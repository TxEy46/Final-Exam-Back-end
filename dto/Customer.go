package dto

// Customer DTO ใช้ในการส่งข้อมูลลูกค้า
type Customer struct {
	CustomerID  int    `json:"customer_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}
