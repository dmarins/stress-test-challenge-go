package usecase_test

import (
	"testing"

	mock "github.com/dmarins/stress-test-challenge-go/internal/usecase/mocks"
	"go.uber.org/mock/gomock"
)

type TestVars struct {
	Ctrl *gomock.Controller

	MockHttpClient *mock.MockHttpClientInterface
}

func BuildTestVars(t *testing.T) TestVars {
	testVars := TestVars{}

	ctrl := gomock.NewController(t)
	testVars.Ctrl = ctrl

	testVars.MockHttpClient = mock.NewMockHttpClientInterface(ctrl)

	return testVars
}
