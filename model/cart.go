package model

import (
	"sort"
	"time"
)

type Cart struct {
	ID       int           `json:"id"`
	UserID   int           `json:"userId"`
	Date     time.Time     `json:"date"`
	Products []cartProduct `json:"products"`
	V        int           `json:"__v"`
}

type Carts []Cart

type CartSummary struct {
	Name
	Total float64
	Cart
}

func (carts Carts) FindHighestValueCart(products Products, users Users) CartSummary {
	//sort by id to ensure order
	sort.Slice(products, func(i, j int) bool {
		return products[i].ID < products[j].ID
	})

	resultCart := CartSummary{}

	for _, cart := range carts {
		cartValue := 0.0
		for _, prod := range cart.Products {
			cartValue += products[prod.ProductID-1].Price * float64(prod.Quantity)
		}

		if cartValue > resultCart.Total {
			resultCart.Total = cartValue
			resultCart.Cart = cart

		}
	}

	user := users.GetByID(resultCart.UserID)
	resultCart.Name = user.Name

	return resultCart
}
