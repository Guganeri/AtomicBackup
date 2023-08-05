package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/lib/pq"
)

func main() {
	// Configurações do banco de dados
	dsn := "user=username dbname=database sslmode=disable"

	// Caminho para salvar o arquivo de dump
	dumpFilePath := "/path/to/backups/dump.sql"

	// Comando para executar o pg_dump
	cmd := exec.Command("pg_dump", dsn, "--file="+dumpFilePath)

	// Executa o comando e verifica por erros
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao executar o pg_dump:", err)
		os.Exit(1)
	}

	fmt.Println("Dump do banco de dados PostgreSQL concluído com sucesso!")
}
