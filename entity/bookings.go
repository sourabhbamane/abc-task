package entity

import "time"

type Bookings struct {
	MemberName string    `json:"name"`
	Date       time.Time `json:"date"`
}
