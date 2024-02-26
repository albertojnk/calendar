package migrations

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func CreateDatabase(connString string) error {

	ctx := context.Background()
	// Create a new connection pool
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		logrus.Errorf("error connecting to db, %v", err)
		return err
	}
	defer pool.Close()

	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logrus.Errorf("error beginTX, %v", err)
		return err
	}

	// Create tables within the database
	_, err = tx.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS calendars (
			id SERIAL PRIMARY KEY,
			date timestamptz NOT NULL,
			email text UNIQUE NOT NULL
		)
	 `)
	if err != nil {
		logrus.Errorf("error creating table calendars, %v", err)
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, `
		CREATE TYPE timestamp_schedule AS (
			start timestamptz,
			finish timestamptz
		);
	 `)
	if err != nil {
		logrus.Errorf("error creating type timestamp_schedule, %v", err)
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS calendar_schedule (
			id SERIAL PRIMARY KEY,
			week_day integer NOT NULL,
			schedule timestamp_schedule[],
			calendar_id integer REFERENCES calendars(id)
		)
	 `)
	if err != nil {
		logrus.Errorf("error creating table calendar_schedule, %v", err)
		tx.Rollback(ctx)
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
	}

	logrus.Info("Database and tables created successfully!")
	return nil
}
