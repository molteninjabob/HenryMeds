package model

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	Id          *uuid.UUID `db:"id"`
	ProviderId  *uuid.UUID `db:"provider_id"`
	StartTime   time.Time  `db:"start_time"`
	ClientId    *uuid.UUID `db:"client_id"`
	ReservedAt  time.Time  `db:"reserved_at"`
	ConfirmedAt time.Time  `db:"confirmed_at"`
}
