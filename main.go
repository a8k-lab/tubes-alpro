package main

import (
	"fmt"
	"jual-beli/item"
	"jual-beli/transaction"
	"jual-beli/utils"
	"os"
)

func main() {
	utils.ClearScreen()
	mainMenu()
}

func mainMenu() {
	utils.PrintIntro()

	for {
		var selectMenu int
		utils.ClearScreen()
		utils.PrintMainMenu()
		utils.InputMenu(&selectMenu, 3)

		switch selectMenu {
		case 1:
			item.ItemMenu()
		case 2:
			transaction.TransactionMenu()
		case 3:
			end()
		}
	}
}

func end() {
	utils.ClearScreen()
	fmt.Println("🛑 Program Berakhir")
	os.Exit(0)
}
