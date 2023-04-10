package appstats

import (
	"io"
	"io/ioutil"
	"net/http"
)

/**
 * @brief Interface for API Caller
 */
type ApiCallerInterface interface {
	HttpGet(url string) (*http.Response, error)
	IoRead(buffer io.Reader) ([]byte, error)
}

/**
 * Helper to perform HTTP Operations
 */
type ApiCaller struct {
	// Empty
}

/**
 * @brief Constructor for API Caller
 * @return Instance of ApiCaller
 */
func MakeApiCaller() ApiCaller {
	return ApiCaller{}
}

/**
 * @brief Make HTTP GET call to given URL
 * @param url URL to which HTTPS GET call to be made
 * @return API Reponse if succeeded
 * @return Instance of error if failed
 */
func (apiCaller *ApiCaller) HttpGet(url string) (*http.Response, error) {
	return http.Get(url)
}

/**
 * @brief Read API response content from HTTP Response buffer
 * @param buffer API response buffer
 * @return API Reponse as byte array if succeeded
 * @return Instance of error if failed
 */
func (apiCaller *ApiCaller) IoRead(buffer io.Reader) ([]byte, error) {
	return ioutil.ReadAll(buffer)
}
