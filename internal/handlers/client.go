package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/molteninjabob/HenryMeds/internal/access"
	"github.com/molteninjabob/HenryMeds/internal/access/model"
	"github.com/molteninjabob/HenryMeds/internal/usecase"
	"github.com/molteninjabob/HenryMeds/internal/validate"
)

func GetClient(ctx context.Context, db *access.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO

		w.Header().Set("Content/Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}

func AddClient(ctx context.Context, db *access.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		client := &model.Client{}
		if err := json.NewDecoder(r.Body).Decode(client); err != nil {
			http.Error(w, fmt.Sprintf("error parsing input values: %v", err.Error()), http.StatusBadRequest)
			return
		}

		err := validate.ValidateAddClient(ctx, db, client)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		client, err = usecase.NewClient(ctx, db, client)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(client)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
