package calculations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSI(t *testing.T) {
	t1 := calculateSI(42000.00, 2.5, 5.00)
	assert.Equal(t, "5250.00", t1, "they should be equal")

	t2 := calculateSI(36000.00, 3.75, 5)
	assert.Equal(t, "6750.00", t2, "they should be equal")
}

func TestCI(t *testing.T) {
	// compounded annually
	t1 := calculateCI(42000.00, 2.5, 5.00)
	assert.Equal(t, "5519.14", t1, "they should be equal")

	t2 := calculateCI(36000.00, 3.75, 5)
	assert.Equal(t, "7275.59", t2, "they should be equal")
}
