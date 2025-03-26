package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	CustomerID  int       `gorm:"column:customer_id;primary_key;AUTO_INCREMENT" json:"customer_id"`
	FirstName   string    `gorm:"column:first_name;NOT NULL" json:"first_name"`
	LastName    string    `gorm:"column:last_name;NOT NULL" json:"last_name"`
	Email       string    `gorm:"column:email;NOT NULL" json:"email"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	Address     string    `gorm:"column:address" json:"address"`
	Password    string    `gorm:"column:password;NOT NULL" json:"password"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName กำหนดชื่อ Table ของฐานข้อมูล
func (m *Customer) TableName() string {
	return "customer"
}

// ฟังก์ชันสำหรับการเข้ารหัสรหัสผ่าน
func (m *Customer) HashPassword() error {
	// เข้ารหัสรหัสผ่านโดยใช้ bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// ตั้งค่ารหัสผ่านที่เข้ารหัสแล้ว
	m.Password = string(hashedPassword)
	return nil
}
