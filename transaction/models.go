package transaction

import "jual-beli/item"

type Transaction struct {
	Item      item.Item
	Quantity  int
	BuyerName string
}

type TabTransaction []Transaction

var TransactionList TabTransaction
