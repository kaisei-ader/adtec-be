package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminUser struct {
	gorm.Model
	name  string
	email string
	pass  string
}

func main() {
	// db := dbInit()
	// db.AutoMigrate(&AdminUser{})
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("./template/index.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.Run(":8080")
}

// func dbInit() *gorm.DB {
//	dsn := "root:password@tcp(127.0.0.1:3306)/sample_db?charset=utf8mb4&parseTime=true"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//	return db
// }
