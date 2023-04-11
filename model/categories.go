package model

type Category struct {
	Name       string
	TotalValue float64
}

type Categories []Category

func (categories Categories) DiscoverAllCategories(products []Product) Categories {
	categories = Categories{}
	for _, prod := range products {
		newCategory := true
		for i, category := range categories {
			if prod.Category == category.Name {
				categories[i].TotalValue += prod.Price
				newCategory = false
				break
			}
		}

		if newCategory {
			categories = append(categories, Category{Name: prod.Category, TotalValue: prod.Price})
		}

	}

	return categories
}
