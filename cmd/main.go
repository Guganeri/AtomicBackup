package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	_ "github.com/lib/pq"
	"github.com/shirou/gopsutil/process"
)

func dockerProcessRunning(processDocker string) bool {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("Erro ao criar cliente Docker:", err)
		return false
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		fmt.Println("Erro ao listar contêineres:", err)
		return false
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if name == processDocker {
				return true
			}
		}
	}

	return false
}

func dockerFunc() {
	containerName := "postgres"
	user := "pgApp"
	password := "pgApp"
	dbname := "postgres"
	backupFilename := generateBackupFilename(dbname)

	cmd := exec.Command("docker", "exec", containerName, "pg_dump",
		"--username="+user,
		"--dbname="+dbname,
		"--file=/var/lib/postgres/"+backupFilename)

	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", password))

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Erro ao executar o comando: %s\n%s", err, output)
	}

	fmt.Println("Backup concluído e salvo dentro do container com sucesso!")
}

func generateBackupFilename(dbname string) string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s_%s_bkp.sql", formattedTime, dbname)
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

	nameProcess := "postgres"

	fmt.Println("Iniciando...")

	if processRunning(nameProcess) {
		fmt.Println("Existe um PG na máquina, iniciando processo de DUMP")
		//onPremFunc()
	} else if dockerProcessRunning("/" + nameProcess) {
		fmt.Println("Existe um PG utilizando Docker, iniciando processo e DUMP")
		dockerFunc()
	}

}
