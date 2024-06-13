package item

import "fmt"

func IsItemExist() bool {
	return len(ItemList) > 0
}

func GetItemIndexByName(name string) int {
	for index, item := range ItemList {
		if item.Name == name {
			return index
		}
	}
	return -1
}

func DeleteItemByIndex(index int) {
	var newItemList TabItem

	for i := 0; i < len(ItemList); i++ {
		if i != index {
			newItemList = append(newItemList, ItemList[i])
		}
	}

	ItemList = newItemList
}

func ShowItemList() {
	for index, item := range ItemList {
		number := index + 1
		fmt.Printf("%d. %s dengan harga modal %.f, seharga %.f ada di kategori %s\n", number, item.Name, item.CapitalPrice, item.SalePrice, item.Category)
	}
	fmt.Println("--------------------------")
}
