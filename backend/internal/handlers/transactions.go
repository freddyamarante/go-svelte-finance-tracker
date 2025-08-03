package handlers

import (
	"net/http"

	"github.com/freddyamarante/go-svelte-finance-tracker/internal/database"
	"github.com/freddyamarante/go-svelte-finance-tracker/internal/models"
	"github.com/gin-gonic/gin"
)

// GetTransactions returns all transactions for the authenticated user
func GetTransactions(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var transactions []models.Transaction
	if err := database.DB.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.IndentedJSON(http.StatusOK, transactions)
}

// CreateTransaction creates a new transaction for the authenticated user
func CreateTransaction(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var newTransaction models.Transaction
	if err := c.BindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Set the user ID for the transaction
	newTransaction.UserID = userID.(string)

	if err := database.DB.Create(&newTransaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newTransaction)
}
