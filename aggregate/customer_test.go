package aggregate

import (
	"testing"
)

func TestCusoumer_NewCustuoer(t *testing.T) {
	_, err := NewCustomer("", 0)
	if err != ErrInvalidCustomer {
		t.Errorf("NewCustomer should return ErrInvalidCustomer")
	}
}
