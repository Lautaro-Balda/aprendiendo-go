package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{"quarter_of_a_dozen": 3, "half_of_a_dozen": 6, "dozen": 12, "small_gross": 120, "gross": 144, "great_gross": 1728}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unidadExtraida, unidadExiste := units[unit]

	if !unidadExiste {
		return false
	} else {
		bill[item] += unidadExtraida
		return true
	}

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	cantItemExtraido, itemExist := bill[item]
	unidadExtraida, unitExist := units[unit]
	if !itemExist || !unitExist {
		return false
	} else if cantItemExtraido-unidadExtraida < 0 {
		return false
	} else if cantItemExtraido-unidadExtraida > 0 {
		bill[item] = cantItemExtraido - unidadExtraida
		return true
	} else {
		delete(bill, item)
		return true
	}

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	CantItemExtraido, itemExist := bill[item]

	return CantItemExtraido, itemExist
}
