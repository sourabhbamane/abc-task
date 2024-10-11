package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClass_Success(t *testing.T) {
	// Prepare valid input
	input := map[string]interface{}{
		"name":       "Yoga Class",
		"start_date": "2024-10-01",
		"end_date":   "2024-10-15",
		"capacity":   20,
	}
	inputBody, _ := json.Marshal(input)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodPost, "/classes", bytes.NewBuffer(inputBody))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()

	CreateClass(w, req)

	// Get the response
	res := w.Result()
	defer res.Body.Close()

	// check  status code
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	// Decode the response body
	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	// check the response message
	assert.Equal(t, "class created successfully of name Yoga Class", response["message"])
}

func TestCreateClass_InvalidInput(t *testing.T) {
	// Prepare invalid input (missing name)
	input := map[string]interface{}{
		"start_date": "2024-10-01",
		"end_date":   "2024-10-15",
		"capacity":   20,
	}
	inputBody, _ := json.Marshal(input)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodPost, "/classes", bytes.NewBuffer(inputBody))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()

	CreateClass(w, req)

	// Get the response
	res := w.Result()
	defer res.Body.Close()

	// check  status code
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	// Decode the response body
	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	// check the response message
	assert.Equal(t, "class name cannot be empty", response["message"])
}

func TestCreateClass_InvalidDateFormat(t *testing.T) {
	// Prepare input with an invalid date format
	input := map[string]interface{}{
		"name":       "Yoga Class",
		"start_date": "10-01-2024", // Invalid format
		"end_date":   "2024-10-15",
		"capacity":   20,
	}
	inputBody, _ := json.Marshal(input)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodPost, "/classes", bytes.NewBuffer(inputBody))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()

	CreateClass(w, req)

	// Get the response
	res := w.Result()
	defer res.Body.Close()

	// check  status code
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	// Decode the response body
	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	// check the response message
	assert.Equal(t, "invalid start_date format. use YYYY-MM-DD.", response["message"])
}
