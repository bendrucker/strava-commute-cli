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

func noon(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 12, 0, 0, 0, t.Location())
}

// Label returns a Morning or Evening label
func (c Commute) Label() string {
	if c.Start.Before(noon(c.Start)) {
		return "Morning"
	}

	return "Evening"
}
