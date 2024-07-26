//go:build wireinject
// +build wireinject

package infrastructure

import (
	"github.com/dmarins/stress-test-challenge-go/internal/infrastructure/http"
	"github.com/dmarins/stress-test-challenge-go/internal/usecase"
	"github.com/google/wire"
)

// InitializeStressTest sets up the dependencies for StressTest using Google Wire.
func InitializeStressTest() *usecase.StressTest {
	wire.Build(
		http.NewHttpClient,
		usecase.NewStressTest,
	)

	return &usecase.StressTest{}
}
