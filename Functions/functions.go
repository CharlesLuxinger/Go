package main

import (
	"fmt"
	"net/http"
	"os"
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

	sites := []string{"https://random-status-code.herokuapp.com/", "https://random-status-code.herokuapp.com/"}
	sites = append(sites, "https://www.google.com.br")

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			testaSite(site, i)
		}
		time.Sleep(delay * time.Millisecond)
	}

}

func testaSite(site string, i int) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Server", site, "Up, position: ", i)
	} else {
		fmt.Println("Server", site, "Down, Status Code:", response.StatusCode)
	}

}
