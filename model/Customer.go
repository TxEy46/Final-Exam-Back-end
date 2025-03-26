package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	CustomerID  int       `gorm:"column:customer_id;primary_key;AUTO_INCREMENT"`
	FirstName   string    `gorm:"column:first_name;NOT NULL"`
	LastName    string    `gorm:"column:last_name;NOT NULL"`
	Email       string    `gorm:"column:email;NOT NULL"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Address     string    `gorm:"column:address"`
	Password    string    `gorm:"column:password;NOT NULL"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
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
