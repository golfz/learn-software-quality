package ticket

import "testing"

func TestTicketPrice(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want float64
	}{
		{"Free Ticket when age under 3", 3, 0.0},
		{"Ticket $15 when age at 4 year old", 4, 15.0},
		{"Ticket $15 when age is 15", 15, 15.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Price(tt.age)
			if got != tt.want {
				t.Errorf("Price(%d) = %f; want %f", tt.age, got,
					tt.want)
			}
		})
	}
}
