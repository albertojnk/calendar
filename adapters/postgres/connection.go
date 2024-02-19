package postgres

import (
	"context"
	"log"
	"os"

	"github.com/albertojnk/calendar/environment"
	"github.com/go-pg/pg/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var dbPool *pgxpool.Pool

func InitDB() *pgxpool.Pool {
	urlOptions, err := pg.ParseURL(environment.PostgresURLConn)
	if err != nil {
		log.Fatal("InitPostgreDB :: error connecting Postgres DB", err)
	}

	if urlOptions.ApplicationName == "" {
		if urlOptions.ApplicationName, err = os.Hostname(); err != nil {
			logrus.Info("Error on get hostname generating : err", err)
		}
		if urlOptions.ApplicationName == "" {
			urlOptions.ApplicationName = "NoHostAvailable"
		}
	}

	dbPool, err = pgxpool.New(context.Background(), environment.PostgresURLConn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return dbPool
}

func GetDBX() *pgxpool.Pool {

	if !isPgValid() {
		dbPool = InitDB()
		if !isPgValid() {
			logrus.Info("GetPgConn :: isPgValid returned false")
		}
	}

	return dbPool
}

func ClosePgXConn() {
	if dbPool == nil {
		logrus.Info("ClosePgConn :: nil pointer of postgress connection")
		return
	}

	logrus.Info("ClosePgConn :: closing postgress connection")
	dbPool.Close()

	dbPool = nil
}
