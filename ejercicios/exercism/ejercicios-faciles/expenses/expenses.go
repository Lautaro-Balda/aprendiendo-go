package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var newIn []Record
	for _, v := range in {
		if predicate(v) {
			newIn = append(newIn, v)
		}
	}
	return newIn
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return c == r.Category
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var sum float64
	bydays := ByDaysPeriod(p)
	for _, value := range in {
		if bydays(value) {
			sum += value.Amount
		}
	}
	return sum
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var sum float64
	byDays := ByDaysPeriod(p)
	byCat := ByCategory(c)
	var isACat bool
	for _, value := range in {
		if byDays(value) && byCat(value) {
			sum += value.Amount
		}
		if byCat(value) {
			isACat = true
		}
	}
	if !isACat {
		return 0, errors.New("unknown category entertainment")
	} else {
		return sum, nil
	}

}
