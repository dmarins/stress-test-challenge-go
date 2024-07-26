package usecase

import (
	"sync"
	"time"

	"github.com/dmarins/stress-test-challenge-go/internal/domain/entity"
	"github.com/dmarins/stress-test-challenge-go/internal/infrastructure/http"
)

type StressTest struct {
	HttpClient http.HttpClient
}

func NewStressTest(httpClient http.HttpClient) *StressTest {
	return &StressTest{
		HttpClient: httpClient,
	}
}

func (st *StressTest) Execute(url string, totalRequests, concurrency int) entity.Result {
	var wg sync.WaitGroup

	resultChan := make(chan entity.Result, totalRequests)

	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < totalRequests/concurrency; j++ {
				statusCode := st.HttpClient.DoRequest(url)
				resultChan <- entity.Result{StatusCodeDist: map[int]int{statusCode: 1}}
			}
		}()
	}

	wg.Wait()
	close(resultChan)

	result := entity.Result{
		TotalTime:      time.Since(startTime).Seconds(),
		TotalRequests:  totalRequests,
		StatusCodeDist: make(map[int]int),
	}

	for r := range resultChan {

		result.SuccessfulRequests += r.StatusCodeDist[200]

		for code, count := range r.StatusCodeDist {
			result.StatusCodeDist[code] += count
		}
	}

	return result
}
