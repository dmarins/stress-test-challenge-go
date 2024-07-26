package main

import (
	"flag"
	"log"

	"github.com/dmarins/stress-test-challenge-go/internal/infrastructure"
	"github.com/dmarins/stress-test-challenge-go/internal/infrastructure/cli"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser estressado")
	requests := flag.Int("requests", 0, "Número total de solicitações HTTP")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")

	flag.Parse()

	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		log.Fatal("Todos os parâmetros são obrigatórios e devem ser maiores que zero.")
	}

	stressTest := infrastructure.InitializeStressTest()

	cli := cli.NewCLI(*stressTest)
	cli.Run(*url, *requests, *concurrency)
}
