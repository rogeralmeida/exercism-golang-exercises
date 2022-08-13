package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, exists := units[unit]
	if exists {
		bill[item] += quantity
	}
	return exists
}

func cannotRemoveItem(doesItemExistInBill, doesUnitExistInUnits bool, newQuantity int) bool {
	return !doesItemExistInBill || !doesUnitExistInUnits || newQuantity < 0
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemQuantity, doesItemExistInBill := bill[item]
	unitQuantity, doesUnitExistInUnits := units[unit]
	newQuantity := itemQuantity - unitQuantity
	if cannotRemoveItem(doesItemExistInBill, doesUnitExistInUnits, newQuantity) {
		return false
	}
	if newQuantity == 0 {
		delete(bill, item)
	} else {
		bill[item] = newQuantity
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qtd, exists := bill[item]
	return qtd, exists
}
