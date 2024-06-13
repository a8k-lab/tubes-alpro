package item

type Item struct {
	Name         string
	CapitalPrice float64
	SalePrice    float64
	Category     string
	TotalSold    int
}

type TabItem []Item

var ItemList TabItem
