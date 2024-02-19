package postgres

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

const nTries = 3
const timeBetweenVerifies = time.Second * 30
const waitTimeBetweenFailures = time.Second * 10

var isPgConnActive bool = false

func SelfMonitor() {

	try := 0
	for {

		switch isConnActive() {

		case true:
			time.Sleep(timeBetweenVerifies)

		case false:
			//Recovering
			logrus.Error("invalid connection detected, retrying again in ", (nTries-try)*10, " seconds")

			try = ((try + 1) % nTries)
			if (try % nTries) == 0 {
				logrus.Error("now recovering postgres connection after pool failure")
				ClosePgXConn()
				GetDBX()
			}

			time.Sleep(waitTimeBetweenFailures)
		}
	}
}

func isConnActive() bool {

	var err error

	if _, err = dbPool.Exec(context.Background(), "SELECT 1"); err != nil {
		logrus.Info("isConnActive :: error : ", err)
	} else {
		tn := time.Now()
		if ((tn.Minute() % 5) == 0) && tn.Second() < 30 {
			logrus.Info("isConnActive :: success : pg connection is alive!")
		}
	}
	isPgConnActive = err == nil

	return isPgConnActive
}

func isPgValid() bool {

	if dbPool == nil || !isPgConnActive {
		return false
	}

	return true
}
