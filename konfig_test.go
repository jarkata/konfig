package konfig

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	m, err := ReadConfig("test.properties")
	fmt.Printf("err: %v\n", err)
	fmt.Printf("m: %v\n", m)
}
