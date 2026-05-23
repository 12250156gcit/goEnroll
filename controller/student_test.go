package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test case for admin controller

func TestAddStudent(t *testing.T) {
	// api endpoint for test
	url := "http://localhost:8080/student/add"

	// user data to send to server
	var jsonStr = []byte(`{"stdid":"10","fname":"Rin","lname":"Wa","email":"r@gmail.com"}`)

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

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// verifying the response body

	assert.JSONEq(t, `{"Status": "Student added successfully"}`, string(data))

}
