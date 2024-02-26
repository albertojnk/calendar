package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	HttpPort        string
	PostgresURLConn string
	Machine         string

	TelemetryGrpcPort  string
	PermGrpcPort       string
	GeoCodGrpcPort     string
	IntegrationSvcPort string
	CoreGrpcServerPort string
)

func InitEnv(pathEnv string) {

	if err := godotenv.Load(pathEnv); err != nil {
		logrus.Fatal("Err:", err)
	}

	PostgresURLConn = getEnvOrPanic("DATABASE_URL")

}

func getEnv(env string) string {
	return os.Getenv(env)
}

func getEnvOrPanic(env string) string {

	if temp := os.Getenv(env); temp != "" {
		return temp
	}

	panic(fmt.Errorf(`not possible to recover %s from environment`, env))
}
