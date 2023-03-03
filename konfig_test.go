package konfig_test

import (
	"fmt"
	"testing"

	"github.com/jarkata/konfig"
)

func TestRead(t *testing.T) {
	m, err := konfig.ReadConfig("test.properties")
	fmt.Printf("err: %v\n", err)
	fmt.Printf("m: %v\n", m)
}
