package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	WorkingCarsPerHour := (float64(productionRate) * successRate) / 100
	return WorkingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {

	WorkingCarsPerMinute := (float64(productionRate/60) * successRate) / 100
	return int(WorkingCarsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var decenas int
	for {
		if carsCount-10 < 0 {
			break
		}
		carsCount -= 10
		decenas++

	}
	coste := decenas*95000 + carsCount*10000
	return uint(coste)
}
