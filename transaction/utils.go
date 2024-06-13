package transaction

import (
	"fmt"
	"jual-beli/item"
)

func IsTransactionExist() bool {
	return len(TransactionList) > 0
}

func GetSumCapital() float64 {
	var sumCapital float64
	for _, item := range item.ItemList {
		sumCapital += item.CapitalPrice
	}
	return sumCapital
}

func GetSumIncome() float64 {
	var sumIncome float64
	for _, transaction := range TransactionList {
		sumIncome += float64(transaction.Quantity) * transaction.Item.SalePrice
	}
	return sumIncome
}

func DeleteTransactionByIndex(index int) {
	var newTransactionList TabTransaction

	for i := 0; i < len(TransactionList); i++ {
		if i != index {
			newTransactionList = append(newTransactionList, TransactionList[i])
		} else {
			itemName := TransactionList[i].Item.Name
			itemIndex := item.GetItemIndexByName(itemName)
			item.ItemList[itemIndex].TotalSold -= TransactionList[i].Quantity
		}
	}

	TransactionList = newTransactionList
}

func ShowItemList() {
	for index, item := range item.ItemList {
		number := index + 1
		fmt.Printf("%d. %s seharga %.f ada di kategori %s\n", number, item.Name, item.SalePrice, item.Category)
	}
	fmt.Println("--------------------------")
}

func ShowTransactionList() {
	for index, transaction := range TransactionList {
		number := index + 1
		fmt.Printf("%d. %s seharga %.f dengan banyak %d, dibeli oleh %s\n", number, transaction.Item.Name, transaction.Item.SalePrice, transaction.Quantity, transaction.BuyerName)
	}
	fmt.Println("--------------------------")
}

func ShowMostSoldItemList() {
	sortedItemsBySold := item.GetSortedItemsBySold()
	for index, item := range sortedItemsBySold {
		if index == 5 {
			return
		}
		number := index + 1
		fmt.Printf("%d. %s terjual sebanyak %d\n", number, item.Name, item.TotalSold)
	}
	fmt.Println("--------------------------")
}
