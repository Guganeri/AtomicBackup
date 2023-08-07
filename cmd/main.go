package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	_ "github.com/lib/pq"
)

func dockerFunc() {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	containerName := "postgres"
	user := "pgApp"
	password := "pgApp"
	dbname := "postgres"
	backupFilename := formattedTime + " " + dbname + "bkp" + ".sql"

	// Comando de backup usando docker exec
	cmd := exec.Command("docker", "exec", containerName, "pg_dump",
		"--username="+user,
		"--dbname="+dbname,
		"--file=/var/lib/postgres/"+backupFilename)

	// Configurar a senha usando variáveis de ambiente
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", password))

	// Executar o comando
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Erro ao executar o comando: %s\n%s", err, output)
	}

	fmt.Println("Backup concluído e salvo dentro do container com sucesso!")
}

func main() {
	dockerFunc()
}
