package types

type Pizza struct {
	Toppings  []Topping
	CrustType CrustType
	Size      Size
	Quantity  int
}

type Topping string

const (
	TOPPING_UNSPECIFIED Topping = "TOPPING_UNSPECIFIED"
	PEPPERONI           Topping = "PEPPERONI"
	MUSHROOMS           Topping = "MUSHROOMS"
	ONIONS              Topping = "ONIONS"
	SAUSAGE             Topping = "SAUSAGE"
	BACON               Topping = "BACON"
	BLACK_OLIVES        Topping = "BLACK_OLIVES"
	GREEN_PEPPERS       Topping = "GREEN_PEPPERS"
	PINEAPPLE           Topping = "PINEAPPLE"
	ANCHOVIES           Topping = "ANCHOVIES"
)

type CrustType string

const (
	CRUST_TYPE_UNSPECIFIED CrustType = "CRUST_TYPE_UNSPECIFIED"
	THIN                   CrustType = "THIN"
	REGULAR                CrustType = "REGULAR"
	STUFFED                CrustType = "STUFFED"
	NEW_YORK               CrustType = "NEW_YORK"
	DEEP_DISH              CrustType = "DEEP_DISH"
	GLUTEN_FREE            CrustType = "GLUTEN_FREE"
)

type Size string

const (
	SIZE_UNSPECIFIED Size = "SIZE_UNSPECIFIED"
	SMALL            Size = "SMALL"
	MEDIUM           Size = "MEDIUM"
	LARGE            Size = "LARGE"
	EXTRA_LARGE      Size = "EXTRA_LARGE"
)
