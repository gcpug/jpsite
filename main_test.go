package main

import (
	"os"
	"testing"

	_ "github.com/favclip/testerator/datastore"
	_ "github.com/favclip/testerator/memcache"
)

func TestMain(m *testing.M) {
	status := m.Run()

	os.Exit(status)
}
