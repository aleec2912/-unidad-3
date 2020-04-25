package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error
var Libro int

type Person struct {
	ID        uint   `json:"id"`
	Nombre string `json:"nombre"`
	Descripcion  string `json:"descripcion"`
	Autor      string `json:"autor"`
	Editorial	string `json:"editorial"`
	FechaPublicacion	string `json:"fechapublicacion"` 
}


func main() () {
	usuario := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "biblioteca"

	db, err = gorm.Open("mysql", "./gorm.db", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Libro)

	r := gin.Default()
	r.GET("/libro/", GetLibro)
	r.GET("/libro/:id", GetLibro)
	r.POST("/libro", CreateLibro)
	r.PUT("/libro/:id", UpdateLibro)
	r.DELETE("/libro/:id", DeleteLibro)

	r.Run(":8080")
}

func DeleteLibro(c *gin.Context) {
	id := c.Params.ByName("id")
	var libro string
	d := db.Where("id = ?", id).Delete(&libro)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateLibro(c *gin.Context) {

	var libro string
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&libro).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&libro)

	db.Save(&libro)
	c.JSON(200, libro)

}

func CreateLibro(c *gin.Context) {

	var libro string
	c.BindJSON(&libro)

	db.Create(&libro)
	c.JSON(200, libro)
}

func GetLibro(c *gin.Context) {
	id := c.Params.ByName("id")
	var libro string
	if err := db.Where("id = ?", id).First(&libro).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, libro)
	}
}
/*func GetLibro(c *gin.Context) {
	var libro []Libro
	if err := db.Find(&libro).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, libro)
	}

}*/
