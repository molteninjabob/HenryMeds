package access

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	m "github.com/molteninjabob/HenryMeds/internal/access/model"
	"github.com/molteninjabob/HenryMeds/internal/util"
)

// Fetch appointments from the database for the given day
func (db *DB) GetAppointmentsByDay(ctx context.Context, providerId *uuid.UUID, date time.Time) ([]*m.Appointment, error) {
	dateString := date.Format("2024-01-01")

	sql := `SELECT * FROM Appointment 
					WHERE provider_id=':provider_id'
					AND DATE(start_time) = :start`

	appointments := []*m.Appointment{}
	err := db.Select(appointments, sql, map[string]interface{}{"provider_id": providerId, "start": dateString})
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (db *DB) UpsertAppointments(ctx context.Context, appts []*m.Appointment) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	for _, appt := range appts {
		if appt.Id == nil {
			appt.Id = util.NewUUID()
		}

		sql := `INSERT INTO Appointment (id, provider_id, start_time)
						VALUES (:id, :provider_id, :start_time)
						ON CONFLICT (provider_id, start_time) 
						DO NOTHING`

		_, err := tx.NamedExec(sql, appts)
		if err != nil {
			// Rollback transaction on error
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing appointments upsert transaction: %w", err)
	}

	return nil
}

func (db *DB) UpdateAppointment(ctx context.Context, appt *m.Appointment) (*m.Appointment, error) {
	// TODO

	// Obtain mutex lock on appointments

	sql := `INSERT INTO Appointment (id, provider_id, start_time)
					VALUES (:id, :provider_id, :start_time)
					ON CONFLICT (provider_id, start_time) 
					DO NOTHING`

	_, err := db.NamedExec(sql, appt)
	if err != nil {
		return nil, err
	}

	// Release mutex lock

	// scan result and return updated appointment details
	return &m.Appointment{}, nil
}
