package main

import (
	"ec2/controller"
	"ec2/model"
	"ec2/service"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	connectionString := os.Getenv("APP_DB_CONNECTION_STRING")
	if connectionString == "" {
		connectionString = "root:gromizk123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}
	db.AutoMigrate(&model.Person{})

	ps := service.NewDBPersonService(db)
	pc := controller.NewPersonController(ps)
	e.GET("/persons", pc.Get)
	e.POST("/persons", pc.Add)
	e.Logger.Fatal(e.Start(":" + port))
}
