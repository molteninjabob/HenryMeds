package access

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	m "github.com/molteninjabob/HenryMeds/internal/access/model"
	"github.com/molteninjabob/HenryMeds/internal/util"
)

func (db *DB) Get(ctx context.Context, clientId uuid.UUID) (*m.Client, error) {
	client := &m.Client{}
	sql := fmt.Sprintf("SELECT * FROM Client WHERE Id='%s'", clientId.String())
	result := db.QueryRow(sql)
	if err := result.Scan(client); err != nil {
		return nil, err
	}
	return client, nil
}

func (db *DB) Create(ctx context.Context, client *m.Client) error {
	if client.Id == nil {
		client.Id = util.NewUUID()
	}

	sql := fmt.Sprintf("INSERT INTO Client (Id, Name, Email, Phone) VALUES (%s, %s, %s, %s)", client.Id, client.Name, client.Email, client.Phone)
	result, err := db.Exec(sql)
	if err != nil {
		return err
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsInserted < 1 {
		return errors.New("error inserting client into the database")
	}

	return nil
}
