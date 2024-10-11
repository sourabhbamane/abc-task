package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookClass_Failure(t *testing.T) {
	// Prepare the input JSON body
	input := map[string]interface{}{
		"name": "John Doe",
		"date": "2024-10-20",
	}
	inputBody, _ := json.Marshal(input)

	// Create a request and response recorder
	req := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(inputBody))
	w := httptest.NewRecorder()

	// Call the BookClass function
	BookClass(w, req)

	// Get the result and assert the response
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)
	assert.Equal(t, "no class is available on the selected date", response["message"])
}

func TestBookClass_InvalidInput(t *testing.T) {
	// Invalid input (missing name)
	input := map[string]interface{}{
		"date": "2024-10-20",
	}
	inputBody, _ := json.Marshal(input)

	// Create a request and response recorder
	req := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(inputBody))
	w := httptest.NewRecorder()

	// Calling function
	BookClass(w, req)

	// Get the result and assert the response
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)
	assert.Equal(t, "member name cannot be empty", response["message"])
}

func TestBookClass_InvalidDateFormat(t *testing.T) {
	// Invalid date format
	input := map[string]interface{}{
		"name": "John Doe",
		"date": "20-10-2024", // Wrong date format
	}
	inputBody, _ := json.Marshal(input)

	// Create a request and response recorder
	req := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(inputBody))
	w := httptest.NewRecorder()

	// Call the BookClass function
	BookClass(w, req)

	// Get the result and assert the response
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)
	assert.Equal(t, "invalid date format. Use YYYY-MM-DD", response["message"])
}
