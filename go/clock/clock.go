package clock

import "fmt"

type Clock struct {
	hour, min int
}

func New(hour, minute int) Clock {
	return normalizeClock(hour, minute)

}
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}

func (c Clock) Add(min int) Clock {
	return normalizeClock(c.hour, c.min+min)
}

func (c Clock) Subtract(min int) Clock {
	return normalizeClock(c.hour, c.min-min)
}

func normalizeClock(hour, min int) Clock {
	totalMin := ((hour*60+min)%(24*60) + 24*60) % (24 * 60)
	return Clock{totalMin / 60, totalMin % 60}
}
