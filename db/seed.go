package db

// SeedDatabase popula o banco de dados com dados iniciais
func SeedDatabase() {
	// Inserir transações de exemplo
	_, err := Database.Exec(`
	INSERT INTO transactions (type, amount, date) VALUES
		('deposit', 5000, '2025-01-01'),
		('expense', 1200, '2025-01-10'),
		('deposit', 3000, '2025-01-15'),
		('expense', 800, '2025-01-20');
	`)
	if err != nil {
		panic("Erro ao popular a tabela transactions")
	}

	// Inserir residentes de exemplo
	_, err = Database.Exec(`
	INSERT INTO residents (name, contributed) VALUES
		('João Silva', 1),
		('Maria Oliveira', 1),
		('Carlos Souza', 0),
		('Ana Pereira', 1);
	`)
	if err != nil {
		panic("Erro ao popular a tabela residents")
	}
}
