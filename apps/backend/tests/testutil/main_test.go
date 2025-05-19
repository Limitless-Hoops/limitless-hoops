package testutil

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ConnectTestDB()
	code := m.Run()
	TearDownTestDB()
	os.Exit(code)
}
