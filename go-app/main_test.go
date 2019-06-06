package main

import (
	"os"
	"testing"

	_ "github.com/favclip/testerator/datastore"
	_ "github.com/favclip/testerator/memcache"
)

func TestMain(m *testing.M) {
	Initialize()

	status := m.Run()

	os.Exit(status)
}
