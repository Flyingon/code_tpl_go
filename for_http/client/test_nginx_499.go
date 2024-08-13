package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.DefaultClient
	c.Timeout = 1 * time.Second

	jsonData := []byte(`{"key": "value"}`)

	// Send POST request
	resp, err := c.Post("http://127.0.0.1", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
