package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRound(t *testing.T) {
	var (
		values = []struct {
			old float64
			num int
			new float64
		}{
			{288.934592, 2, 288.93},
			{1.945123, 3, 1.945},
			{1.944956231, 3, 1.945},
		}
	)

	for _, v := range values {
		assert.Equal(t, v.new, Round(v.old, v.num), "no eq")
	}

}
