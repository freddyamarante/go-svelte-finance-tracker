package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/freddyamarante/go-svelte-finance-tracker/internal/database"
	"github.com/freddyamarante/go-svelte-finance-tracker/internal/models"
)

// GetTransactions returns all transactions
func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	if err := database.DB.Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.IndentedJSON(http.StatusOK, transactions)
}

// CreateTransaction creates a new transaction
func CreateTransaction(c *gin.Context) {
	var newTransaction models.Transaction
	if err := c.BindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := database.DB.Create(&newTransaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newTransaction)
} 