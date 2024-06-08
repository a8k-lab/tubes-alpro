package item

import (
	"jual-beli/utils"
)

func ItemMenu() {
	for {
		utils.ClearScreen()
		utils.PrintItemMenu()
		selectMenu := utils.InputMenu(6)

		switch selectMenu {
		case 1:
			AddItem()
		case 2:
			EditItem()
		case 3:
			DeleteItem()
		case 4:
			ShowItem()
		case 5:
			SearchItem()
		case 6:
			return
		}
	}
}
