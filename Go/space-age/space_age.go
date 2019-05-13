// Calculate ages based on different planets in the solar system
package space

// Planet type/enum for cleaner code
type Planet string
const (
	Earth    Planet = "Earth"
	Mercury  Planet = "Mercury"
	Venus    Planet = "Venus"
	Mars 	 Planet = "Mars"
	Jupiter  Planet = "Jupiter"
	Saturn   Planet = "Saturn"
	Uranus   Planet = "Uranus"
	Neptune  Planet = "Neptune"
 )

// Translate age in seconds to given planet's years
func Age (seconds float64, planet Planet) float64 {
	switch planet {
		case Earth:
			return AgeInEarthYears(seconds)
		case Mercury:
			return AgeInMercuryYears(seconds)
		case Venus:
			return AgeInVenusYears(seconds)
		case Mars:
			return AgeInMarsYears(seconds)
		case Jupiter:
			return AgeInJupiterYears(seconds)
		case Saturn:
			return AgeInSaturnYears(seconds)
		case Uranus:
			return AgeInUranusYears(seconds)
		case Neptune:
			return AgeInNeptuneYears(seconds)
	}
	// Shouldn't get this far! I don't know how to handle 
	// errors yet so I'm just returning -1
	return -1
}

// Translate age in seconds to age in Earth years
func AgeInEarthYears(seconds float64) float64 {
	var orbital_period float64 = 31557600
	var age float64 = seconds / orbital_period
	return age
}

// Translate age in seconds to age in Mercury years
func AgeInMercuryYears(seconds float64) float64 {
	var age_in_earth_years float64 = AgeInEarthYears(seconds)
	var orbital_period float64 = 0.2408467 
	var age float64 = age_in_earth_years / orbital_period
	return age
}

// Translate age in seconds to age in Venus years
func AgeInVenusYears(seconds float64) float64 {
	var age_in_earth_years float64 = AgeInEarthYears(seconds)
	var orbital_period float64 = 0.61519726 
	var age float64 = age_in_earth_years / orbital_period
	return age
}

// Translate age in seconds to age in Mars years
func AgeInMarsYears(seconds float64) float64 {
	var age_in_earth_years float64 = AgeInEarthYears(seconds)
	var orbital_period float64 = 1.8808158  
	var age float64 = age_in_earth_years / orbital_period
	return age
}

// Translate age in seconds to age in Jupiter years
func AgeInJupiterYears(seconds float64) float64 {
	var age_in_earth_years float64 = AgeInEarthYears(seconds)
	var orbital_period float64 = 11.862615
	var age float64 = age_in_earth_years / orbital_period
	return age
}

// Translate age in seconds to age in Saturn years
func AgeInSaturnYears(seconds float64) float64 {
	var age_in_earth_years float64 = AgeInEarthYears(seconds)
	var orbital_period float64 = 29.447498 
	var age float64 = age_in_earth_years / orbital_period
	return age
}

// Translate age in seconds to age in Uranus years
func AgeInUranusYears(seconds float64) float64 {
	var age_in_earth_years float64 = AgeInEarthYears(seconds)
	var orbital_period float64 = 84.016846
	var age float64 = age_in_earth_years / orbital_period
	return age
}

// Translate age in seconds to age in Neptune years
func AgeInNeptuneYears(seconds float64) float64 {
	var age_in_earth_years float64 = AgeInEarthYears(seconds)
	var orbital_period float64 = 164.79132 
	var age float64 = age_in_earth_years / orbital_period
	return age
}
