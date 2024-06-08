package transaction

import (
	"jual-beli/utils"
)

func TransactionMenu() {
	for {
		utils.ClearScreen()
		utils.PrintTransactionMenu()
		selectMenu := utils.InputMenu(5)

		switch selectMenu {
		case 1:
			// AddTransaction()
		case 2:
			// EditTransaction()
		case 3:
			// DeleteTransaction()
		case 4:
			// ShowTransactions()
		case 5:
			return
		}
	}
}
