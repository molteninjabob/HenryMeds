package types

import (
	"time"

	"github.com/google/uuid"
)

type GetScheduleInput struct {
	ProviderId *uuid.UUID
	Days       []time.Time
}

type SetScheduleInput struct {
	ProviderId *uuid.UUID
	StartTime  time.Time
	EndTime    time.Time
}

type ReservationInput struct {
	AppointmentId *uuid.UUID
	ClientId      *uuid.UUID
}

type ConfirmationInput struct {
	AppointmentId *uuid.UUID
}
