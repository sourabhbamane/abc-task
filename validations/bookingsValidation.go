package validations

import (
	"errors"
	"healthclub/entity"
)

// validations for class booking
func ValidateBooking(booking entity.Bookings, allClasses []entity.Class) error {
	if booking.MemberName == "" {
		return errors.New("member name cannot be empty")
	}

	var validClass *entity.Class

	// Check if booking date falls within an existing class schedule
	for _, class := range allClasses {
		if !booking.Date.Before(class.StartDate) &&
			!booking.Date.After(class.EndDate) {
			validClass = &class
			break
		}
	}

	if validClass == nil {
		return errors.New("no class is available on the selected date")
	}

	return nil
}
