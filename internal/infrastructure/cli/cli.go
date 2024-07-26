package cli

import (
	"fmt"

	"github.com/dmarins/stress-test-challenge-go/internal/domain/entity"
	"github.com/dmarins/stress-test-challenge-go/internal/usecase"
)

type Cli struct {
	StressTestUseCase usecase.StressTest
}

func NewCLI(stressTestUseCase usecase.StressTest) *Cli {
	return &Cli{
		StressTestUseCase: stressTestUseCase,
	}
}

func (cli *Cli) Run(url string, requests, concurrency int) {
	result := cli.StressTestUseCase.Execute(url, requests, concurrency)

	cli.PrintReport(result)
}

func (cli *Cli) PrintReport(result entity.Result) {
	fmt.Printf("Tempo total gasto na execução: %.2f segundos\n", result.TotalTime)
	fmt.Printf("Quantidade total de requests realizados: %d\n", result.TotalRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", result.SuccessfulRequests)
	fmt.Println("Distribuição de outros códigos de status HTTP:")

	for code, count := range result.StatusCodeDist {
		fmt.Printf("%d: %d\n", code, count)
	}
}
