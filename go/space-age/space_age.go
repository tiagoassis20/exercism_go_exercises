// Given an age in seconds, calculate how old someone would be on:
// Mercury: orbital period 0.2408467 Earth years
// Venus: orbital period 0.61519726 Earth years
// Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
// Mars: orbital period 1.8808158 Earth years
// Jupiter: orbital period 11.862615 Earth years
// Saturn: orbital period 29.447498 Earth years
// Uranus: orbital period 84.016846 Earth years
// Neptune: orbital period 164.79132 Earth years
// So if you were told someone were 1,000,000,000 seconds old, you should be able to say that they're 31.69 Earth-years old.

package space

type Planet string

// Given age in seconds and witch planet the function
// returns the age on years to the relative planet
func Age(seconds float64, planet Planet) float64 {
	const earthYear float64 = 31557600
	switch planet {
	case "Mercury":
		return seconds / (0.2408467 * earthYear)
	case "Venus":
		return seconds / (0.61519726 * earthYear)
	case "Earth":
		return seconds / (1 * earthYear)
	case "Mars":
		return seconds / (1.8808158 * earthYear)
	case "Jupiter":
		return seconds / (11.862615 * earthYear)
	case "Saturn":
		return seconds / (29.447498 * earthYear)
	case "Uranus":
		return seconds / (84.016846 * earthYear)
	case "Neptune":
		return seconds / (164.79132 * earthYear)
	default:
		return -1
	}
}
