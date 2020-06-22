package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 1000

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
		imprimeLogs()
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
	}
}

func exibeIntroducao() {
	nome := "Charles"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			testaSite(site, i)
		}
		time.Sleep(delay * time.Millisecond)
	}

}

func testaSite(site string, i int) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	if response.StatusCode == 200 {
		fmt.Println("Server", site, "Up, position: ", i)
	} else {
		fmt.Println("Server", site, "Down, Status Code:", response.StatusCode)
	}
	registraLog(site, response.StatusCode)
}

func lerSitesDoArquivo() []string {
	var lines []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		lines = append(lines, strings.TrimSpace(line))

		if err != io.EOF {
			break
		}
	}

	file.Close()

	return lines
}

func registraLog(site string, statusCode int) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	file.WriteString(time.Now().Format("2006/01/02 15:04:05") + " - " + site + " - status code: " + strconv.Itoa(statusCode) + "\n")

	file.Close()
}

func imprimeLogs(){
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	fmt.Println(string(arquivo))
}
