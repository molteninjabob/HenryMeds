# Ben Bolton - Reservation - Backend v3

## Pre-requisites
- go 1.23.1 installed
- postgresql installed

## Setup
- From a PostgreSQL client or the command line, create and set up a new database schema:

```
CREATE DATABASE my_database;
CREATE USER my_user WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE my_database TO my_user;

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
);
```

- Create a .env file in the root directory populating the values in .env.example
- Run `go mod download`
- Run `go run main.go`

## Assumptions
- Front-end will constrain time values selectable by the user on the hour, quarter-hour and half-hour
- If given a start and end time of a schedule that is not in 15 minute increments, api will discard extra time that is not enough to create a slot
- To retrieve appointments, the front-end must provide multiple datetimes

## TO-DO
- CRUD operations for both Client and Provider
- Appointment Confirmation use case
- Request Appointment use case
- Comprehensive unit tests

## Later Enhancements
- Implement logging
- Implement robust authentication
- Implement authorization & permissions management
- Github Actions workflow to execute unit tests
- Dockerize the service