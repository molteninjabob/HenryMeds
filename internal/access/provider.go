package access

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	m "github.com/molteninjabob/HenryMeds/internal/access/model"
	"github.com/molteninjabob/HenryMeds/internal/util"
)

func (db *DB) GetProviderById(ctx context.Context, providerId *uuid.UUID) (*m.Provider, error) {
	provider := &m.Provider{}
	result := db.QueryRow("SELECT * FROM Provider WHERE id=$1", providerId.String())
	if err := result.Scan(&provider.Id, &provider.Name, &provider.Email, &provider.Phone); err != nil {
		return nil, err
	}
	return provider, nil
}

func (db *DB) CreateProvider(ctx context.Context, prov *m.Provider) error {
	if prov.Id == nil {
		prov.Id = util.NewUUID()
	}

	sql := fmt.Sprintf("INSERT INTO Provider (Id, Name, Email, Phone) VALUES (%s, %s, %s, %s)", prov.Id, prov.Name, prov.Email, prov.Phone)
	result, err := db.Exec(sql)
	if err != nil {
		return err
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsInserted < 1 {
		return errors.New("error inserting provider into the database")
	}

	return nil
}
