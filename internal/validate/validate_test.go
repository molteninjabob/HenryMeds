package validate

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/molteninjabob/HenryMeds/config"
	"github.com/molteninjabob/HenryMeds/internal/access"
	"github.com/molteninjabob/HenryMeds/internal/types"
	"github.com/molteninjabob/HenryMeds/internal/util"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateSchedule(t *testing.T) {
	ctx := context.Background()
	db, mock, err := access.NewMockDbConn(ctx)
	if err != nil {
		t.Fatal("error", err.Error())
	}
	defer db.Close()

	tests := []struct {
		name    string
		input   *types.SetScheduleInput
		wantErr bool
		msg     string
	}{
		{
			name: "valid schedule",
			input: &types.SetScheduleInput{
				ProviderId: util.NewUUID(),
				StartTime:  time.Now().Add(time.Hour * 1),
				EndTime:    time.Now().Add(time.Hour * 2),
			},
			wantErr: false,
			msg:     "",
		},
		{
			name: "missing provider id",
			input: &types.SetScheduleInput{
				ProviderId: nil,
				StartTime:  time.Now().Add(time.Hour * 1),
				EndTime:    time.Now().Add(time.Hour * 2),
			},
			wantErr: true,
			msg:     "error validating schedule request: ProviderId cannot be nil",
		},
		{
			name: "start and end time equal",
			input: &types.SetScheduleInput{
				ProviderId: util.NewUUID(),
				StartTime:  time.Now().Add(time.Hour * 1),
				EndTime:    time.Now().Add(time.Hour * 1),
			},
			wantErr: true,
			msg:     "error validating schedule request: appointment start and end times must be at least 15 minutes apart",
		},
		{
			name: "end time after start",
			input: &types.SetScheduleInput{
				ProviderId: util.NewUUID(),
				StartTime:  time.Now().Add(time.Hour * 2),
				EndTime:    time.Now().Add(time.Hour * 1),
			},
			wantErr: true,
			msg:     "error validating schedule request: appointment start time must be before the end time",
		},
		{
			name: "start time in the past",
			input: &types.SetScheduleInput{
				ProviderId: util.NewUUID(),
				StartTime:  time.Now().Add(time.Hour * -1),
				EndTime:    time.Now().Add(time.Hour * 1),
			},
			wantErr: true,
			msg:     "error validating schedule request: appointment start time must be in the future",
		},
		{
			name: "duration is 14 minutes",
			input: &types.SetScheduleInput{
				ProviderId: util.NewUUID(),
				StartTime:  time.Now().Add(time.Hour * 1),
				EndTime:    time.Now().Add(1*time.Hour + 14*time.Minute),
			},
			wantErr: true,
			msg:     "error validating schedule request: appointment start and end times must be at least 15 minutes apart",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the expected query and result
			rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
				AddRow(tt.input.ProviderId, "", "", "")

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM Provider WHERE id=$1`)).
				WithArgs(tt.input.ProviderId).
				WillReturnRows(rows)

			err := ValidateSetSchedule(ctx, db, tt.input)
			if err != nil {
				assert.True(t, tt.wantErr)
				assert.Equal(t, tt.msg, err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.msg, "")
			}
		})
	}
}
