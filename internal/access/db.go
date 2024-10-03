package access

import (
	"context"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	cfg "github.com/molteninjabob/HenryMeds/config"
)

/*
DB SCHEMA for visibility

CREATE TABLE provider (
  id text,
  name text,
  email text,
	phone text
);

CREATE TABLE client (
  id text,
  name text,
  email text,
	phone text
);

CREATE TABLE appointment (
  id text,
  provider_id text,
	start_time timestamp without time zone,
	client_id text,
  reserved_at timestamp without time zone,
  confirmed_at timestamp without time zone,
	confirmed boolean
);
*/

type DB struct {
	*sqlx.DB
}

type MockDB struct {
	DB
}

func NewDbConn(ctx context.Context) (*DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable", cfg.DBUser, cfg.DBName)) // "postgres://postgres:postgres@localhost/mydb?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &DB{
		DB: db,
	}, nil
}

func NewMockDbConn(ctx context.Context) (*DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	return &DB{DB: sqlx.NewDb(db, "sqlmock")}, mock, nil
}
