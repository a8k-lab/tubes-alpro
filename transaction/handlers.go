package transaction

import (
	"fmt"
	"jual-beli/item"
	"jual-beli/utils"
)

func AddTransactionMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Transaksi", "Tambah")

	if item.IsItemExist() {
		var newTransaction Transaction
		var selectedItem item.Item
		var selectedNumber, itemListLength int
		var isConfirm bool

		ShowItemList()
		fmt.Println("📦 Masukkan nomor barang")

		itemListLength = len(item.ItemList)
		utils.InputMenu(&selectedNumber, itemListLength)
		selectedItem = item.ItemList[selectedNumber-1]

		utils.ClearScreen()
		fmt.Printf("%s dipilih\n", selectedItem.Name)
		fmt.Println("--------------------------")
		fmt.Println("🔢 Masukkan jumlah barang:")
		fmt.Print("> ")
		fmt.Scan(&newTransaction.Quantity)

		fmt.Println("📒 Masukkan nama pembeli (tidak boleh ada spasi):")
		fmt.Print("> ")
		fmt.Scan(&newTransaction.BuyerName)

		item.ItemList[selectedNumber-1].TotalSold += newTransaction.Quantity

		utils.ClearScreen()
		fmt.Println("Konfirmasi penambahan transaksi:")
		fmt.Printf("%s sebanyak %d, untuk pembeli %s\n", selectedItem.Name, newTransaction.Quantity, newTransaction.BuyerName)
		utils.ConfirmInput(&isConfirm, "menambahkan transaksi")

		if isConfirm {
			newTransaction.Item = selectedItem
			TransactionList = append(TransactionList, newTransaction)
			utils.PrintSuccessMessage("Transaksi berhasil ditambahkan")
			return
		} else {
			return
		}
	} else {
		fmt.Println("Tidak dapat menambahkan transaksi")
		utils.ShowEmptyItemList()
	}
}

func EditTransactionMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Transaksi", "Edit")

	if IsTransactionExist() {
		var newTransaction, selectedTransaction Transaction
		var selectedNumber int
		var isConfirm bool

		ShowTransactionList()
		fmt.Println("Masukkan nomor transaksi yang ingin diedit")
		utils.InputMenu(&selectedNumber, len(TransactionList))
		newTransaction = TransactionList[selectedNumber-1]
		selectedTransaction = TransactionList[selectedNumber-1]

		utils.ClearScreen()
		fmt.Printf("Edit data transaksi \"%d\"\n", selectedNumber)
		fmt.Println("Kosongkan input jika tidak ingin diubah")
		fmt.Println("--------------------------")

		fmt.Println("🔢 Masukkan jumlah baru barang:")
		fmt.Print("> ")
		fmt.Scanln(&newTransaction.Quantity)

		fmt.Println("📒 Masukkan nama baru pembeli (tidak boleh ada spasi):")
		fmt.Print("> ")
		fmt.Scanln(&newTransaction.BuyerName)

		utils.ClearScreen()
		fmt.Println("Konfirmasi perubahan transaksi:")
		fmt.Printf("🔢 Jumlah %d diubah menjadi %d\n", selectedTransaction.Quantity, newTransaction.Quantity)
		fmt.Printf("📒 Nama pembeli \"%s\" diubah menjadi \"%s\"\n", selectedTransaction.BuyerName, newTransaction.BuyerName)
		utils.ConfirmInput(&isConfirm, "mengedit transaksi")

		if isConfirm {
			itemName := selectedTransaction.Item.Name
			itemIndex := item.GetItemIndexByName(itemName)
			if newTransaction.Quantity > selectedTransaction.Quantity {
				item.ItemList[itemIndex].TotalSold += newTransaction.Quantity - selectedTransaction.Quantity
			} else {
				item.ItemList[itemIndex].TotalSold -= selectedTransaction.Quantity - newTransaction.Quantity
			}

			newTransaction.Item = selectedTransaction.Item
			TransactionList[selectedNumber-1] = newTransaction
			utils.PrintSuccessMessage("Transaksi berhasil diubah")
		}
		return
	} else {
		utils.ShowEmptyTransactionList()
	}
}

func DeleteTransactionMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Transaksi", "Hapus")

	if IsTransactionExist() {
		var selectedTransaction Transaction
		var selectedNumber int
		var isConfirm bool

		ShowTransactionList()
		fmt.Println("Masukkan nomor transaksi yang ingin dihapus")
		utils.InputMenu(&selectedNumber, len(TransactionList))
		selectedTransaction = TransactionList[selectedNumber-1]

		utils.ClearScreen()
		fmt.Println("Konfirmasi penghapusan transaksi:")
		fmt.Printf("Transaksi nomor %d\n", selectedNumber)
		fmt.Printf("Nama: %s\n", selectedTransaction.Item.Name)
		fmt.Printf("Jumlah: %d\n", selectedTransaction.Quantity)
		fmt.Printf("Pembeli: %s\n", selectedTransaction.BuyerName)
		utils.ConfirmInput(&isConfirm, "menghapus transaksi")

		if isConfirm {
			DeleteTransactionByIndex(selectedNumber - 1)
			utils.PrintSuccessMessage("Transaksi berhasil dihapus")
		}
		return
	} else {
		utils.ShowEmptyTransactionList()
	}
}

func ShowTransactionMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Transaksi", "Lihat")

	if IsTransactionExist() {
		ShowTransactionList()
		fmt.Println("Klik Enter untuk kembali ...")
		fmt.Scanln()
	} else {
		utils.ShowEmptyTransactionList()
	}
}

func ShowCapitalMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Transaksi", "Data modal")

	sumCapital := GetSumCapital()
	fmt.Printf("Data modal saat ini adalah %.f\n", sumCapital)
	fmt.Println("Klik Enter untuk kembali ...")
	fmt.Scanln()
}

func ShowIncomeMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Transaksi", "Data pendapatan")

	sumIncome := GetSumIncome()
	fmt.Printf("Data pendapatan saat ini adalah %.f\n", sumIncome)
	fmt.Println("Klik Enter untuk kembali ...")
	fmt.Scanln()
}

func ShowMostSoldsMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Barang", "Paling banyak terjual")

	if IsTransactionExist() {
		fmt.Println("5 barang yang paling banyak terjual")
		ShowMostSoldItemList()
		fmt.Println("Klik Enter untuk kembali ...")
		fmt.Scanln()
	} else {
		utils.ShowEmptyTransactionList()
	}
}
