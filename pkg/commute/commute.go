package commute

import (
	"time"

	"github.com/martinlindhe/unit"
)

// Commute represents general attributes of a regular commute
type Commute struct {
	Type     string
	Start    time.Time
	Distance unit.Length
	Speed    unit.Speed
}

// Duration computes the commute time given the distance and speed
func (c Commute) Duration() time.Duration {
	return time.Duration(c.Distance.Nanometers() / c.Speed.MetersPerSecond())
}
