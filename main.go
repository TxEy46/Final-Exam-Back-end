package main

import (
	"fmt"
	"go-backend/controller"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// ดึงค่า DSN สำหรับการเชื่อมต่อกับฐานข้อมูล
	dsn := viper.GetString("mysql.dsn")
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection success!")

	// ส่ง db ไปที่ controller
	controller.StartServer(db)
}
