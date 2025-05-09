package contolller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test case 1: for registered admin
func TestAddLogic(t *testing.T) {
	// add endpoint to call
	url := "http://localhost:8080/login"
	// data of type byte slice
	jsonStr := []byte(`{"email":"d@gmailcom", "password":"1122334455"}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// create a pointer variable client which points to Client struct
	client := &http.Client{}
	// client sends the request using Do() and gets http response
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// read the closing of response body until function terminates
	body, err := io.ReadAll(res.Body) // body is of type []byte
	if err != nil {
		panic(err)
	}

	// assert response code and message
	assert.Equal(t, 201, res.StatusCode)
	assert.Contains(t, string(body), "success")
	// for debug: fmt.Println("status:", res.Status)
	// for debug: fmt.Println("response:", string(body))
}

// test case 2: for admin which does not exist
func TestAddUserNotExist(t *testing.T) {
	url := "http://localhost:8080/login"
	data := []byte(`{"email":"none@gmail.com", "password":"pass"}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, _ := client.Do(req)
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	assert.Equal(t, 400, res.StatusCode)
	assert.Contains(t, string(body), "no rows in result set")
}
