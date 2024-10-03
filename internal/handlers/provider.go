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

func GetProvider(ctx context.Context, db *access.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO

		w.Header().Set("Content/Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}

func AddProvider(ctx context.Context, db *access.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		provider := &model.Provider{}
		if err := json.NewDecoder(r.Body).Decode(provider); err != nil {
			http.Error(w, fmt.Sprintf("error parsing input values: %v", err.Error()), http.StatusBadRequest)
			return
		}

		err := validate.ValidateAddProvider(ctx, db, provider)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		provider, err = usecase.NewProvider(ctx, db, provider)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(provider)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
