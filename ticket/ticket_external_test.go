package ticket_test

import (
	"github.com/golfz/learn-software-quality/ticket"
	"testing"
)

func TestPrice(t *testing.T) {
	// Arrange
	want := 0.0

	// Act
	got := ticket.Price(1)

	// Assert
	if got != want {
		t.Errorf("ticket.Price(1) = %f, want %f", got, want)
	}
}
