package validations

import (
	"errors"
	"healthclub/entity"
	"time"
)

// validating class details
func ValidateClass(newClass entity.Class, allClasses []entity.Class) error {
	if newClass.Name == "" {
		return errors.New("class name cannot be empty")
	}

	if newClass.Capacity <= 0 {
		return errors.New("capacity must be greater than zero")
	}

	// Check if start_date is before or equal to end_date
	if newClass.StartDate.After(newClass.EndDate) {
		return errors.New("start_date cannot be after end_date")
	}

	// checking no overlapping class exists for any date between start_date and end_date
	for _, existingClass := range allClasses {
		// Check for any overlap with the current class
		if newClass.StartDate.Before(existingClass.EndDate.Add(24*time.Hour)) &&
			newClass.EndDate.After(existingClass.StartDate.Add(-24*time.Hour)) {
			return errors.New("a class already exists in the specified date range")
		}
	}

	return nil
}
