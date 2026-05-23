package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test case for admin controller

func TestAdmLogin(t *testing.T) {
	// api endpoint for test
	url := "http://localhost:8080/login"

	// user data to send to server
	var jsonStr = []byte(`{"email":"wangdi@gmail.com","password":"pass"}`)

	// create an http request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// create a http client
	client := &http.Client{}

	// send api
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// check the data in response body
	data, _ := io.ReadAll(resp.Body)

	// verifying the status code and response body

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// verifying the response body

	assert.JSONEq(t, `{"Status": "Login Success"}`, string(data))

}

// admin not exist test case
func TestAdmLoginNotFound(t *testing.T) {
	// api endpoint for test
	url := "http://localhost:8080/login"

	// user data to send to server
	var jsonStr = []byte(`{"email":"rg@gmail.com","password":"pass"}`)

	// create an http request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// create a http client
	client := &http.Client{}

	// send api
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// check the data in response body
	data, _ := io.ReadAll(resp.Body)

	// verifying the status code and response body

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	// verifying the response body

	assert.JSONEq(t, `{"error": "sql: no rows in result set"}`, string(data))

}
