package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/molteninjabob/HenryMeds/internal/access"
	"github.com/molteninjabob/HenryMeds/internal/types"
	"github.com/molteninjabob/HenryMeds/internal/usecase"
	"github.com/molteninjabob/HenryMeds/internal/validate"
)

func GetSchedule(ctx context.Context, db *access.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		getScheduleInput := &types.GetScheduleInput{}
		err := validate.ValidateGetSchedule(ctx, db, getScheduleInput)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		schedule, err := usecase.GetAppointmentsForProvider(ctx, db, getScheduleInput)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content/Type", "application/json")
		json.NewEncoder(w).Encode(schedule)
		w.WriteHeader(http.StatusOK)
	})
}

func SetSchedule(ctx context.Context, db *access.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		schedule := &types.SetScheduleInput{}
		if err := json.NewDecoder(r.Body).Decode(schedule); err != nil {
			http.Error(w, fmt.Sprintf("error parsing input values: %v", err.Error()), http.StatusBadRequest)
			return
		}

		err := validate.ValidateSetSchedule(ctx, db, schedule)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := usecase.SubmitSchedule(ctx, db, schedule); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
