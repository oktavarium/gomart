package orders

import (
	"strconv"
	"strings"
)

func checkOrderNumber(order string) bool {
	if len(order) == 0 {
		return false
	}
	var sum int
	parity := len(order) % 2
	for i, v := range order {
		digit, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}

func compressOrderNumber(order string) string {
	return strings.ReplaceAll(order, " ", "")
}
