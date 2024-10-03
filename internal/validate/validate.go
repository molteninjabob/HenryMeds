package validate

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/molteninjabob/HenryMeds/internal/access"
	"github.com/molteninjabob/HenryMeds/internal/access/model"
	"github.com/molteninjabob/HenryMeds/internal/types"
)

func ValidateGetClient(ctx context.Context, db *access.DB, clientId *uuid.UUID) error {
	// TODO
	return nil
}
func ValidateAddClient(ctx context.Context, db *access.DB, client *model.Client) error {
	// TODO
	return nil
}
func ValidateGetProvider(ctx context.Context, db *access.DB, providerId *uuid.UUID) error {
	// TODO
	return nil
}
func ValidateAddProvider(ctx context.Context, db *access.DB, client *model.Provider) error {
	// TODO
	return nil
}
func ValidateGetSchedule(ctx context.Context, db *access.DB, input *types.GetScheduleInput) error {
	// TODO
	return nil
}
func ValidateReservation(ctx context.Context, db *access.DB, input *types.ReservationInput) error {
	// TODO
	return nil
}
func ValidateConfirmation(ctx context.Context, db *access.DB, input *types.ConfirmationInput) error {
	// TODO
	return nil
}

func ValidateSetSchedule(ctx context.Context, db *access.DB, sched *types.SetScheduleInput) error {
	if sched.ProviderId == nil {
		return errors.New("error validating schedule request: ProviderId cannot be nil")
	}

	if !sched.StartTime.Before(sched.EndTime) {
		return errors.New("error validating schedule request: appointment start time must be before the end time")
	}
	if sched.StartTime.Before(time.Now()) {
		return errors.New("error validating schedule request: appointment start time must be in the future")
	}
	if sched.EndTime.Sub(sched.StartTime).Minutes() < 15 {
		return errors.New("error validating schedule request: appointment start and end times must be at least 15 minutes apart")
	}

	_, err := db.GetProviderById(ctx, sched.ProviderId)
	if err != nil {
		return fmt.Errorf("error validating provider id: %w", err)
	}

	return nil
}
