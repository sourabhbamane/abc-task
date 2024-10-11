package controllers

import (
	"encoding/json"
	"healthclub/entity"
	"healthclub/error"
	"healthclub/repository"
	"healthclub/validations"
	"net/http"
	"time"
)

func BookClass(w http.ResponseWriter, r *http.Request) {

	//will set headers
	w.Header().Set("Content-Type", "application/json")

	var input struct {
		MemberName string `json:"name"`
		Date       string `json:"date"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.Error{Message: "invalid input"})
		return
	}

	// Parse date
	bookingDate, err := time.Parse(format, input.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.Error{Message: "invalid date format. Use YYYY-MM-DD"})
		return
	}

	booking := entity.Bookings{
		MemberName: input.MemberName,
		Date:       bookingDate,
	}

	// Read existing classes from CSV to validate booking
	existingClasses, err := repository.ReadClassesFromCSV("classes.csv")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.Error{Message: "unable to read classes data from csv file"})
		return
	}

	// Validate the booking
	if err := validations.ValidateBooking(booking, existingClasses); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.Error{Message: err.Error()})
		return
	}

	//write data in .csv file
	record := []string{booking.MemberName, booking.Date.Format(format)}
	if err := repository.AppendToCSV("bookings.csv", record); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.Error{Message: "unable to write data to file"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(error.Success{Message: "class booked successfully on date " + input.Date})
}
