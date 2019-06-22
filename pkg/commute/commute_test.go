package commute

import (
	"testing"
	"time"

	"github.com/martinlindhe/unit"
	"github.com/stretchr/testify/assert"
)

func TestCommuteDuration(t *testing.T) {
	commute := Commute{
		Distance: 8 * unit.Mile,
		Speed:    12 * unit.MilesPerHour,
	}
	duration := commute.Duration()

	assert.Equal(t, time.Duration(40)*time.Minute, duration)
}
