package item

import (
	"jual-beli/utils"
)

func ItemMenu() {
	var selectMenu int

	for {
		utils.ClearScreen()
		utils.PrintItemMenu()
		utils.InputMenu(&selectMenu, 6)

		switch selectMenu {
		case 1:
			AddItemMenu()
		case 2:
			EditItemMenu()
		case 3:
			DeleteItemMenu()
		case 4:
			ShowItemMenu()
		case 5:
			SearchItemMenu()
		case 6:
			return
		}
	}
}
