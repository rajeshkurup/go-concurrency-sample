package appstats

import (
	"io"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type ApiCallerMock struct {
	mock.Mock
}

func (apiCallerMock *ApiCallerMock) HttpGet(url string) (*http.Response, error) {
	args := apiCallerMock.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (apiCallerMock *ApiCallerMock) IoRead(buffer io.Reader) ([]byte, error) {
	args := apiCallerMock.Called(buffer)
	return args.Get(0).([]byte), args.Error(1)
}
