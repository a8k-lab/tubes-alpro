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
			AddTransactionMenu()
		case 2:
			EditTransactionMenu()
		case 3:
			DeleteTransactionMenu()
		case 4:
			ShowTransactionMenu()
		case 5:
			return
		}
	}
}
