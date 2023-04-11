package model

import (
	"reflect"
	"testing"
)

func TestCarts_FindHighestValueCart(t *testing.T) {
	testCarts := Carts{
		{
			ID:     1,
			UserID: 1,
			// value - 83.25
			Products: []cartProduct{
				cartProduct{
					ProductID: 1,
					Quantity:  2,
				},
				cartProduct{
					ProductID: 2,
					Quantity:  3,
				},
			},
		},
		{
			ID:     2,
			UserID: 2,
			// value - 92.5
			Products: []cartProduct{
				cartProduct{
					ProductID: 2,
					Quantity:  3,
				},
				cartProduct{
					ProductID: 3,
					Quantity:  1,
				},
			},
		},
	}

	type args struct {
		products Products
		users    Users
	}
	tests := []struct {
		name  string
		carts Carts
		args  args
		want  CartSummary
	}{
		{
			name:  "Slice empty",
			carts: Carts{},
			args:  args{products: []Product{}, users: Users{}},
			want:  CartSummary{},
		},
		{
			name:  "Positive case",
			carts: testCarts,
			args: args{products: Products{
				{ID: 1, Price: 10.50},
				{ID: 2, Price: 20.75},
				{ID: 3, Price: 30.25},
			}, users: Users{
				{ID: 1, Name: Name{Firstname: "Firstname A", Lastname: "Lastname A"}},
				{ID: 2, Name: Name{Firstname: "Firstname B", Lastname: "Lastname B"}},
			}},
			want: CartSummary{
				Name:  Name{Firstname: "Firstname B", Lastname: "Lastname B"},
				Total: 92.5,
				Cart:  testCarts[1],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.carts.FindHighestValueCart(tt.args.products, tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Carts.FindHighestValueCart() = %v, want %v", got, tt.want)
			}
		})
	}
}
