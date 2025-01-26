package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"net/http"
	"os"
	"os/exec"

	"io"
	"sitevila/db"        // Pacote para banco de dados
	"sitevila/financial" // Pacote para rotas financeiras

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const repoPath = "/var/www/sitevila"

func main() {
	// Carregue as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
		return
	}

	// Obtenha a chave secreta do ambiente
	secret := os.Getenv("WEBHOOK_SECRET")
	if secret == "" {
		fmt.Println("Chave secreta não configurada")
		return
	}

	// Inicializar o banco de dados
	db.InitDatabase()
	defer db.Database.Close()

	// Popular o banco de dados com dados iniciais
	db.SeedDatabase()

	r := gin.Default()

	// Webhook para deploy automático
	r.POST("/webhook", func(c *gin.Context) {
		// Verifique a assinatura do webhook
		signature := c.GetHeader("X-Hub-Signature-256")
		body, _ := io.ReadAll(c.Request.Body)

		if !verifySignature(signature, body, secret) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Assinatura inválida"})
			return
		}

		// Execute o comando git pull
		cmd := exec.Command("git", "-C", repoPath, "pull")
		output, err := cmd.CombinedOutput()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":  err.Error(),
				"output": string(output),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Deploy realizado com sucesso",
			"output":  string(output),
		})
	})

	// Rotas do sistema financeiro
	financial.RegisterRoutes(r)

	// Inicie o servidor na porta 8080
	r.Run(":8080")
}

// Função para verificar a assinatura do webhook
func verifySignature(signature string, body []byte, secret string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expectedSignature := "sha256=" + hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}
