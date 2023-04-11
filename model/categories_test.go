package model

import (
	"reflect"
	"testing"
)

func TestCategories_DiscoverAllCategories(t *testing.T) {
	type args struct {
		products []Product
	}
	tests := []struct {
		name       string
		categories Categories
		args       args
		want       Categories
	}{{
		name:       "Empty slice",
		categories: Categories{},
		args:       args{[]Product{}},
		want:       Categories{},
	}, {
		name:       "Positive case",
		categories: Categories{},
		args: args{[]Product{
			{Price: 10.50, Category: "Category A"},
			{Price: 20.75, Category: "Category B"},
			{Price: 30.25, Category: "Category A"},
			{Price: 15.00, Category: "Category C"},
			{Price: 12.50, Category: "Category B"},
			{Price: 22.00, Category: "Category C"},
		}},
		want: Categories{
			{Name: "Category A", TotalValue: 40.75},
			{Name: "Category B", TotalValue: 33.25},
			{Name: "Category C", TotalValue: 37.00},
		},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.categories.DiscoverAllCategories(tt.args.products); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Categories.DiscoverAllCategories() = %v, want %v", got, tt.want)
			}
		})
	}
}
