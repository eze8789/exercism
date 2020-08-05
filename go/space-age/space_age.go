package space

// Planet is a custom type to create the map
type Planet string

// Age function calculate an age starting from it in seconds against different planets
func Age(n float64, planet Planet) float64 {

	// Initializing a variable for seconds in a year within Earth
	var secsInYear = 31557600.0

	// Map matching planets against orbital periods
	var planets = map[Planet]float64{
		"Earth":   1.0,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return n / secsInYear / planets[planet]
}
