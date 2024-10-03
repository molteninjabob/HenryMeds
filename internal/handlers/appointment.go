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

func GetAppointmentsForProvider(ctx context.Context, db *access.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content/Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}

func MakeReservation(ctx context.Context, db *access.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reservationInput := &types.ReservationInput{}
		if err := json.NewDecoder(r.Body).Decode(reservationInput); err != nil {
			http.Error(w, fmt.Sprintf("error parsing input values: %v", err.Error()), http.StatusBadRequest)
			return
		}

		err := validate.ValidateReservation(ctx, db, reservationInput)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := usecase.RequestAppointment(ctx, db, reservationInput); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func ConfirmReservation(ctx context.Context, db *access.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		confirmationInput := &types.ConfirmationInput{}
		if err := json.NewDecoder(r.Body).Decode(confirmationInput); err != nil {
			http.Error(w, fmt.Sprintf("error parsing input values: %v", err.Error()), http.StatusBadRequest)
			return
		}

		err := validate.ValidateConfirmation(ctx, db, confirmationInput)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := usecase.ConfirmAppointment(ctx, db, confirmationInput); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
