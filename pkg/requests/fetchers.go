package requests

import (
	"encoding/json"

	"github.com/SzymekN/OpenX-Internship/model"
)

// url for getting all carts had to be changed to provide start and end date
var URLAllCarts = "https://fakestoreapi.com/carts?startdate=2019-12-10&enddate=2022-10-10"
var URLAllUsers = "https://fakestoreapi.com/users"
var URLAllProducts = "https://fakestoreapi.com/products"

// generic fetch function for any datatype
func fetchData(d HttpGetter, url string, data interface{}) (interface{}, error) {
	body, err := MakeApiRequest(d, url)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// fetch users function with type assertion
func FetchUsers(d HttpGetter, url string) (model.Users, error) {

	result, err := fetchData(d, url, &model.Users{})
	if err != nil {
		return model.Users{}, err
	}

	if users, ok := result.(*model.Users); ok {
		return *users, nil
	}

	return model.Users{}, nil
}

// fetch products function with type assertion
func FetchProducts(d HttpGetter, url string) (model.Products, error) {

	result, err := fetchData(d, url, &model.Products{})
	if err != nil {
		return model.Products{}, err
	}

	if products, ok := result.(*model.Products); ok {
		return *products, nil
	}

	return model.Products{}, nil
}

// fetch carts function with type assertion
func FetchCarts(d HttpGetter, url string) (model.Carts, error) {

	result, err := fetchData(d, url, &model.Carts{})
	if err != nil {
		return model.Carts{}, err
	}

	if carts, ok := result.(*model.Carts); ok {
		return *carts, nil
	}

	return model.Carts{}, nil
}
