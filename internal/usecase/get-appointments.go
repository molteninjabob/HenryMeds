package usecase

import (
	"context"

	"github.com/molteninjabob/HenryMeds/internal/access"
	"github.com/molteninjabob/HenryMeds/internal/access/model"
	"github.com/molteninjabob/HenryMeds/internal/types"
)

func GetAppointmentsForProvider(ctx context.Context, db *access.DB, input *types.GetScheduleInput) ([]*model.Appointment, error) {
	var err error
	var appointments []*model.Appointment
	for _, day := range input.Days {
		appointments, err = db.GetAppointmentsByDay(ctx, input.ProviderId, day)
		if err != nil {
			return nil, err
		}
	}
	return appointments, nil
}
