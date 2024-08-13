package actions

import (
	"github.com/mike_jacks/pizza_co/inventory_service/domain/types"
)

func CheckInventory(pizzas []types.Pizza) (map[types.Topping]int, map[types.CrustType]int, map[types.Size]int) {
	toppingsCount := make(map[types.Topping]int)
	crustTypesCount := make(map[types.CrustType]int)
	crustSizesCount := make(map[types.Size]int)

	for _, pizza := range pizzas {
		for i := 0; i < pizza.Quantity; i++ {
			for _, topping := range pizza.Toppings {
				toppingsCount[topping]++
			}
			crustTypesCount[pizza.CrustType]++
			crustSizesCount[pizza.Size]++
		}
	}

	return toppingsCount, crustTypesCount, crustSizesCount
}
