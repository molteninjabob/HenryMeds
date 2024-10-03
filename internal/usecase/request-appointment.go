package usecase

import (
	"context"

	"github.com/molteninjabob/HenryMeds/internal/access"
	"github.com/molteninjabob/HenryMeds/internal/types"
)

func RequestAppointment(ctx context.Context, db *access.DB, reservation *types.ReservationInput) error {
	// TODO

	// Verify that the start time of the appointment is at least 24 hours ahead of time.Now()

	// Add ClientID to appointment

	// Create goroutine that checks for a confirmation on the appointment.
	// Use a channel to receive a confirmation when added.
	// If 30 minutes go by without confirmation, remove the ClientId from the appointment
	// and return goroutine.
	return nil
}
