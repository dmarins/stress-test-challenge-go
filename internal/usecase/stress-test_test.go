package usecase_test

import (
	"testing"

	"github.com/dmarins/stress-test-challenge-go/internal/usecase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func buildTestContext(t *testing.T) (usecase.StressTest, TestVars) {
	testVars := BuildTestVars(t)
	sut := usecase.NewStressTest(testVars.MockHttpClient)

	return *sut, testVars
}

func TestStressTestExecute_ShouldBeReturnsHttpStatusCode200AtAllTimes(t *testing.T) {

	sut, testVars := buildTestContext(t)

	testVars.
		MockHttpClient.
		EXPECT().
		DoRequest(gomock.Any()).
		Return(200).
		Times(100)

	result := sut.Execute("http://example.com", 100, 10)

	assert.Equal(t, 100, result.TotalRequests)
	assert.Equal(t, 100, result.SuccessfulRequests)
	assert.Equal(t, 100, result.StatusCodeDist[200])
}

func TestStressTestExecute_ShouldBeReturnsHttpStatusCode50010Times(t *testing.T) {

	sut, testVars := buildTestContext(t)

	callCount := 0

	testVars.
		MockHttpClient.
		EXPECT().
		DoRequest(gomock.Any()).
		DoAndReturn(func(_ string) int {
			callCount++
			if callCount <= 90 {
				return 200
			}

			return 500
		}).
		Times(100)

	result := sut.Execute("http://example.com", 100, 10)

	assert.Equal(t, 100, result.TotalRequests)
	assert.Equal(t, 90, result.SuccessfulRequests)
	assert.Equal(t, 10, result.StatusCodeDist[500])
}
