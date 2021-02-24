package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databeseURL string
)

func TestMain(m *testing.M) {
	databeseURL = os.Getenv("DATABESE_URL")
	if databeseURL == "" {
		databeseURL = "host=localhost dbname=restapi_test sslmode=disable user=postgres password=12"
	}

	os.Exit(m.Run())
}
