// Package weather provee herramientas para pronosticar el clima en varias ciudades de "Goblinocus".
package weather

var (
	// CurrentCondition reprecenta la condicion actual.
	CurrentCondition string
	// CurrentLocation reprecenta la ubicacion actual.
	CurrentLocation string
)

// Forecast retorna un string que resume los datos de CurrentLocation y CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
