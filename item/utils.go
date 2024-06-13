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

func GetSortedItemsBySold() TabItem {
	var sortedItemList TabItem
	var pass, i, n int
	var temp Item

	sortedItemList = ItemList
	pass = 1
	n = len(ItemList)

	for pass <= n-1 {
		i = pass
		temp = sortedItemList[pass]
		for i > 0 && temp.TotalSold > sortedItemList[i-1].TotalSold {
			sortedItemList[i] = sortedItemList[i-1]
			i--
		}

		sortedItemList[i] = temp
		pass++
	}

	return sortedItemList
}

func GetSortedItemsByCategory() TabItem {
	var sortedItemList TabItem
	var pass, i, n int
	var temp Item

	sortedItemList = ItemList
	pass = 1
	n = len(ItemList)

	for pass <= n-1 {
		i = pass
		temp = sortedItemList[pass]
		for i > 0 && temp.Category < sortedItemList[i-1].Category {
			sortedItemList[i] = sortedItemList[i-1]
			i--
		}

		sortedItemList[i] = temp
		pass++
	}

	return sortedItemList
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

func ShowSortedItemListByCategory() {
	sortedItemList := GetSortedItemsByCategory()
	for index, item := range sortedItemList {
		number := index + 1
		fmt.Printf("%d. %s dengan harga modal %.f, seharga %.f ada di kategori %s\n", number, item.Name, item.CapitalPrice, item.SalePrice, item.Category)
	}
	fmt.Println("--------------------------")
}
