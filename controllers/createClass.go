package controllers

import (
	"encoding/json"
	"healthclub/entity"
	"healthclub/error"
	"healthclub/repository"
	"healthclub/validations"
	"net/http"
	"strconv"
	"time"
)

// Date format (YYYY-MM-DD)
const format = "2006-01-02"

func CreateClass(w http.ResponseWriter, r *http.Request) {

	//set content type
	w.Header().Set("Content-Type", "application/json")

	var input struct {
		Name      string `json:"name"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		Capacity  int    `json:"capacity"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.Error{Message: "invalid input"})
		return
	}

	// Parse dates
	startDate, err := time.Parse(format, input.StartDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.Error{Message: "invalid start_date format. use YYYY-MM-DD."})
		return
	}
	endDate, err := time.Parse(format, input.EndDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.Error{Message: "invalid end_date format. use YYYY-MM-DD."})
		return
	}

	newClass := entity.Class{
		Name:      input.Name,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  input.Capacity,
	}
	// Read existing classes from the CSV file
	existingClasses, err := repository.ReadClassesFromCSV("classes.csv")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.Error{Message: "unable to read classes data"})
		return
	}

	// Validate the new class details
	if err := validations.ValidateClass(newClass, existingClasses); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.Error{Message: err.Error()})
		return
	}

	record := []string{newClass.Name, newClass.StartDate.Format(format), newClass.EndDate.Format(format), strconv.Itoa(newClass.Capacity)}
	if err := repository.AppendToCSV("classes.csv", record); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.Error{Message: "unable to write to file"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(error.Success{Message: "class created successfully of name " + input.Name})
}
