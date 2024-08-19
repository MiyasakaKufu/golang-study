package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Tasks struct {
	*gorm.Model
	Content string `json:"content"`
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
}

func getDBConfig() DBConfig {
	port, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	return DBConfig{
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     port,
		DB:       os.Getenv("DATABASE_NAME"),
	}
}

func connectionDB() (*gorm.DB, error) {
	config := getDBConfig()
	fmt.Println(config)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", config.User, config.Password, config.Host, config.Port, config.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func errorDB(result *gorm.DB, c *gin.Context) bool {
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
		return true
	}
	return false
}

func main() {
	r := gin.Default()
	db, err := connectionDB()

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Tasks{})

	r.POST("/todos/create", func(c *gin.Context) {
		var tasks Tasks

		if err := c.BindJSON(&tasks); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		result := db.Create(&tasks) // これなんで参照渡しなんだろう
		if errorDB(result, c) {
			return
		}
	})

	r.GET("/todos", func(c *gin.Context) {
		var todos []Tasks
		db.Find(&todos)
		c.JSON(200, todos)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
