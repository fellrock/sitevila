package financial

import (
	"net/http"

	"sitevila/db" // Importa o pacote de banco de dados

	"github.com/gin-gonic/gin"
)

// GetTransactions retorna todas as transações
func GetTransactions(c *gin.Context) {
	rows, err := db.Database.Query("SELECT id, type, amount, date FROM transactions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transactions []map[string]interface{}

	for rows.Next() {
		var id int
		var transactionType string
		var amount float64
		var date string
		err := rows.Scan(&id, &transactionType, &amount, &date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar os dados"})
			return
		}

		transactions = append(transactions, map[string]interface{}{
			"id":     id,
			"type":   transactionType,
			"amount": amount,
			"date":   date,
		})
	}

	c.JSON(http.StatusOK, transactions)
}

// RegisterRoutes registra as rotas do sistema financeiro
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/financial")
	{
		api.GET("/transactions", GetTransactions)
	}
}
