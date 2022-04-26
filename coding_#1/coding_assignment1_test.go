package hau

import "testing"

func TestAssignment1(t *testing.T) {
	t.Run("invalid number of array", func(t *testing.T) {
		t.Parallel()
		Assignment1(-1)
		Assignment1(0)
	})
	t.Run("success n=3", func(t *testing.T) {
		t.Parallel()
		Assignment1(3)
	})
	t.Run("success n=4", func(t *testing.T) {
		t.Parallel()
		Assignment1(4)
	})

}
