package usecase_test

import (
	"testing"

	"github.com/dmarins/stress-test-challenge-go/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

type MockHttpClient struct{}

func (m *MockHttpClient) DoRequest(url string) int {
	return 200
}

func TestLoadTest(t *testing.T) {
	result := &entity.Result{} //usecase.Execute("http://example.com", 100, 10)

	assert.Equal(t, 100, result.TotalRequests)
	assert.Equal(t, 100, result.SuccessfulRequests)
	assert.Equal(t, 100, result.StatusCodeDist[200])
}
