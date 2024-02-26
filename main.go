package main

import (
	"github.com/albertojnk/calendar/environment"
)

func main() {
	environment.InitEnv(".env")
	// migrations.CreateDatabase(environment.PostgresURLConn)
}
