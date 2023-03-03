package konfig_test

import (
	"testing"

	"github.com/jarkata/konfig"
)

func TestRead(t *testing.T) {
	konfig.ReadConfig("test.properties")
}
