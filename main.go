package main

import (
	"fmt"
	"log"

	"github.com/SzymekN/OpenX-Internship/model"
	"github.com/SzymekN/OpenX-Internship/pkg/requests"
)

func main() {

	d := requests.DefaultHttpGetter{}
	// task1
	users, err := requests.FetchUsers(d, requests.URLAllUsers)
	if err != nil {
		log.Fatal(err)
	}

	carts, err := requests.FetchCarts(d, requests.URLAllCarts)
	if err != nil {
		log.Fatal(err)
	}

	products, err := requests.FetchProducts(d, requests.URLAllProducts)
	if err != nil {
		log.Fatal(err)
	}

	// task2
	categories := model.Categories{}
	categories = categories.DiscoverAllCategories(products)
	fmt.Println("\nTASK2")
	for _, cat := range categories {
		fmt.Printf("%s: %.2f\n", cat.Name, cat.TotalValue)
	}

	// task3
	fmt.Println("\nTASK3")
	c := carts.FindHighestValueCart(products, users)
	fmt.Println("User with highest value cart: ", c.Firstname, c.Lastname, "\nTotal value:", c.Total)

	// task4
	fmt.Println("\nTASK4")
	u := users.FindFurthestAwayLiving()
	fmt.Println("Users living furthest away")
	fmt.Printf("User1: %s %s, city: %s, coordinates (%s, %s)\n", u[0].Firstname, u[0].Firstname, u[0].Address.City, u[0].Address.Geolocation.Lat, u[0].Address.Geolocation.Long)
	fmt.Printf("User2: %s %s, city: %s, coordinates (%s, %s)\n", u[1].Firstname, u[1].Firstname, u[1].Address.City, u[1].Address.Geolocation.Lat, u[1].Address.Geolocation.Long)

}
