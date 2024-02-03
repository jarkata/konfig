package konfig

import (
	"fmt"
	"log/slog"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	os.Setenv("CONFIG_PATH", "./test.properties")
	s := os.Getenv("CONFIG_PATH")
	slog.Info("", "File", s)
	GetCache("")
	m, err := ReadConfig("test.properties")
	fmt.Printf("err: %v\n", err)
	fmt.Printf("m: %v\n", m)
}

///Users/code/gospace/konfig

///Users/code/gospace/konfig/konfig.properties
