package store_test

import (
	"testing"
)

func Test_String(t *testing.T) {
	t.Paralell()

	act, err := String()
	if err != nil {
		t.Fatalf("expect no error, got %s", err)
	}
}
