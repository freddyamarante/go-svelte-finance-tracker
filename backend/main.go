package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type transaction struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
}

var db *gorm.DB

func initSampleData() {
	var count int64
	db.Model(&transaction{}).Count(&count)

	if count == 0 {
		sampleTransactions := []transaction{
			{ID: "1", Description: "Initial Savings", Amount: 1000.0, Type: "income"},
			{ID: "2", Description: "Groceries", Amount: 50.0, Type: "expense"},
			{ID: "3", Description: "Salary", Amount: 2500.0, Type: "income"},
			{ID: "4", Description: "Gas", Amount: 45.0, Type: "expense"},
			{ID: "5", Description: "Freelance Work", Amount: 300.0, Type: "income"},
		}

		for _, t := range sampleTransactions {
			db.Create(&t)
		}
	}
}

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("finance.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&transaction{})
	initSampleData()

	router := gin.Default()

	// Add CORS middleware with more permissive settings
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * 60 * 60, // 12 hours
	}))

	router.GET("/transactions", getTransactions)
	router.POST("/transactions", postTransactions)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Backend is running"})
	})

	router.Run("localhost:8080")
}

func getTransactions(c *gin.Context) {
	var transactions []transaction
	db.Find(&transactions)
	c.IndentedJSON(http.StatusOK, transactions)
}

func postTransactions(c *gin.Context) {
	var newTransaction transaction
	if err := c.BindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	db.Create(&newTransaction)
	c.IndentedJSON(http.StatusCreated, newTransaction)
}
