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

func createTime(str string) time.Time {
	t, _ := time.Parse(time.RFC822Z, str)
	return t
}

func TestCommuteLabel(t *testing.T) {
	morning := Commute{
		Start: createTime("22 Jun 19 09:18 -0700"),
	}

	assert.Equal(t, morning.Label(), "Morning")

	evening := Commute{
		Start: createTime("22 Jun 19 18:18 -0700"),
	}

	assert.Equal(t, evening.Label(), "Evening")
}
