package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("Test 1+2=3", func(t *testing.T) {
		want := 3
		got := Sum(1, 2)
		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}
