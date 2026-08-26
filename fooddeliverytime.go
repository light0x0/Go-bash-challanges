package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	if order == "burger" {
		price := food{preptime: 15}
		return price.preptime
	} else if order == "chips" {
		price := food{preptime: 10}
		return price.preptime
	} else if order == "nuggets" {
		price := food{preptime: 12}
		return price.preptime

	} else {
		return 404
	}
}
