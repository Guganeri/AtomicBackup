package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/shirou/gopsutil/process"
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

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Erro ao executar o comando: %s\n%s", err, output)
	}

	fmt.Println("Backup concluído e salvo dentro do container com sucesso!")
}

func onPremFunc() {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	user := "pgApp"
	password := "pgApp"
	dbname := "postgres"
	backupFilename := formattedTime + " " + dbname + "bkp" + ".sql"

	connectionString := "user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cmd := exec.Command()

	fmt.Println("Backup concluído e salvo dentro do container com sucesso!")

}

func processRunning(processName string) bool {
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Erro ao obter lista de processos:", err)
		return false
	}

	for _, proc := range processes {
		name, err := proc.Name()
		if err != nil {
			continue
		}

		if strings.Contains(name, processName) {
			return true
		}
	}

	return false
}

func main() {
	processName := "seu_processo" // Substitua pelo nome do processo que deseja verificar

	if processRunning(processName) {
		fmt.Printf("O processo %s está em execução.\n", processName)
	} else {
		fmt.Printf("O processo %s não está em execução.\n", processName)
	}
}
