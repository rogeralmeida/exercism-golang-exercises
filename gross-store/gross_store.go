package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	myMap := map[string]int{}
	myMap["quarter_of_a_dozen"] = 3
	myMap["half_of_a_dozen"] = 6
	myMap["dozen"] = 12
	myMap["small_gross"] = 120
	myMap["gross"] = 144
	myMap["great_gross"] = 1728
	return myMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	return 0, false
}
