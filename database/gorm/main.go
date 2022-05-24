package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type Person struct {
	gorm.Model
	FirstName  string     `gorm:"not null" db:"first_name" json:"first_name"`
	LastName   string     `gorm:"not null" db:"last_name" json:"last_name"`
	Age        int        `gorm:"not null" db:"age" json:"age"`
	City       string     `gorm:"not null" db:"city" json:"city"`
	DeptID     int        `gorm:"not null" db:"dept_id" json:"dept_id"`
	Department Department `gorm:"foreignkey:DeptID"`
}

type Department struct {
	gorm.Model
	Name string `gorm:"not null" db:"name" json:"name"`
}

func (p Person) TableName() string {
	return "person"
}

func main() {
	db, err = gorm.Open(sqlite.Open("./person.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// migrate schema
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Department{})

	department := []Department{
		{Name: "IT"},
		{Name: "HR"},
		{Name: "Sales"},
	}
	db.Create(&department)

	// gin server

	r := gin.Default()
	r.GET("/person", GetPerson)
	r.GET("/person/:id", GetPersonById)
	r.POST("/person", CreatePerson)
	r.PUT("/person/:id", UpdatePerson)
	r.DELETE("/person/:id", DeletePerson)
	r.Run(":8080")
}

func GetPerson(c *gin.Context) {
	var person []Person
	err := db.Find(&person).Error
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, person)
}

func GetPersonById(c *gin.Context) {
	id := c.Param("id")
	var person Person
	err := db.Where("id = ?", id).First(&person).Error
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, person)
}

func CreatePerson(c *gin.Context) {
	var person Person
	err := c.BindJSON(&person)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = db.Create(&person).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func UpdatePerson(c *gin.Context) {
	var updatePerson Person
	err := c.BindJSON(&updatePerson)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var person Person
	err = db.Where("id = ?", c.Param("id")).First(&person).Error
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	person.FirstName = updatePerson.FirstName
	person.LastName = updatePerson.LastName
	person.Age = updatePerson.Age
	person.City = updatePerson.City

	err = db.Save(&person).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func DeletePerson(c *gin.Context) {
	var person Person
	err := db.Where("id = ?", c.Param("id")).First(&person).Error
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	err = db.Delete(&person).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}
