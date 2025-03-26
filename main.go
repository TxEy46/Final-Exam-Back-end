package main

import (
	"fmt"
	"go-backend/controller"
	"go-backend/model"

	"github.com/gin-gonic/gin"
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
	fmt.Println(viper.Get("mysql.dsn"))
	dsn := viper.GetString("mysql.dsn")
	dialactor := mysql.Open(dsn)
	_, err = gorm.Open(dialactor)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection success!")

	gin.SetMode(gin.ReleaseMode)
	controller.StartServer(db)

	// customer := model.Customer{CustomerID: 1}
	// result := db.Create(&customer)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// if result.RowsAffected > 0 {
	// 	fmt.Println("Insert complete")
	// }
	customer := []model.Customer{}
	db.Find(&customer)
	fmt.Printf("%v", customer)
}
