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

	PostgresURLConn = getEnvOrPanic("DB_URL")

	HttpPort = ":" + getEnvOrPanic("HTTPLISTNER")

	TelemetryGrpcPort = getEnvOrPanic("TELEMETRYAPI_PORTRPC")
	TelemetryGrpcPort = ":" + TelemetryGrpcPort

	PermGrpcPort = getEnvOrPanic("PERMISSION_RPCPORT")
	PermGrpcPort = ":" + PermGrpcPort

	GeoCodGrpcPort = getEnvOrPanic("GEOCODING_RPCPORT")
	GeoCodGrpcPort = ":" + GeoCodGrpcPort

	IntegrationSvcPort = getEnvOrPanic("INTEGRATIONSERV_PORT")
	IntegrationSvcPort = ":" + IntegrationSvcPort

	CoreGrpcServerPort = getEnvOrPanic("CORESERVER_RPCPORT")
	CoreGrpcServerPort = ":" + CoreGrpcServerPort

	Machine = getEnv("ENVIRONMENT")

	fmt.Println(PostgresURLConn, HttpPort, Machine)
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
