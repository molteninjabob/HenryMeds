package usecase

import (
	"context"
	"time"

	"github.com/molteninjabob/HenryMeds/internal/access"
	m "github.com/molteninjabob/HenryMeds/internal/access/model"
	"github.com/molteninjabob/HenryMeds/internal/types"
)

const SlotLength = 15

// Allow providers to submit times they are available for appointments
// e.g. On Friday the 13th of August, Dr. Jekyll wants to work between 8am and 3pm
func SubmitSchedule(ctx context.Context, db *access.DB, schedule *types.SetScheduleInput) error {
	var appointments []*m.Appointment
	slots := int(schedule.EndTime.Sub(schedule.StartTime).Minutes() / SlotLength)
	start := schedule.StartTime

	for slot := 0; slot < slots; slot++ {
		appointments = append(appointments, &m.Appointment{
			ProviderId: schedule.ProviderId,
			StartTime:  start,
		})
		start = start.Add(time.Minute * 15)
	}

	if err := db.UpsertAppointments(ctx, appointments); err != nil {
		return err
	}
	return nil
}
