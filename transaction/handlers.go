package transaction

import (
	"fmt"
	"jual-beli/item"
	"jual-beli/utils"
)

type Transaction struct {
	Item      item.Item
	Quantity  int
	BuyerName string
}

type TabTransaction []Transaction

var TransactionList TabTransaction

func IsTransactionExist() bool {
	return len(TransactionList) > 0
}

func DeleteTransactionByIndex(index int) {
	var newTransactionList TabTransaction

	for i := 0; i < len(TransactionList); i++ {
		if i != index {
			newTransactionList = append(newTransactionList, TransactionList[i])
		}
	}

	TransactionList = newTransactionList
}

func ShowTransactionList() {
	for index, item := range TransactionList {
		number := index + 1
		fmt.Printf("%d. %s seharga %f dengan banyak %d, dibeli oleh %s\n", number, item.Item.Name, item.Item.Price, item.Quantity, item.BuyerName)
	}
	fmt.Println("--------------------------")
}

func AddTransactionMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Transaksi", "Tambah")

	if item.IsItemExist() {
		var newTransaction Transaction

		item.ShowItemList()
		fmt.Println("📦 Masukkan nomor barang")

		itemListLength := len(item.ItemList)
		selectedItemNumber := utils.InputMenu(itemListLength)
		selectedItem := item.ItemList[selectedItemNumber-1]

		utils.ClearScreen()
		fmt.Printf("%s dipilih\n", selectedItem.Name)
		fmt.Println("--------------------------")
		fmt.Println("🔢 Masukkan jumlah barang:")
		fmt.Print("> ")
		fmt.Scan(&newTransaction.Quantity)

		fmt.Println("📒 Masukkan nama pembeli:")
		fmt.Print("> ")
		fmt.Scan(&newTransaction.BuyerName)

		newTransaction.Item = selectedItem
		TransactionList = append(TransactionList, newTransaction)
		utils.ClearScreen()
		fmt.Println("✅ Transaksi berhasil ditambahkan")
		fmt.Println("Klik enter untuk kembali ...")
		fmt.Scanln()
	} else {
		fmt.Println("Tidak dapat menambahkan transaksi")
		utils.ShowEmptyItemList()
	}
}

func EditTransactionMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Transaksi", "Edit")

	if IsTransactionExist() {
		var newTransaction Transaction
		var isConfirm string

		ShowTransactionList()
		fmt.Println("Masukkan nomor transaksi yang ingin diedit")
		selectedNumber := utils.InputMenu(len(TransactionList))
		selectedTransaction := TransactionList[selectedNumber-1]

		fmt.Println("🔢 Masukkan jumlah baru barang:")
		fmt.Print("> ")
		fmt.Scan(&newTransaction.Quantity)

		fmt.Println("📒 Masukkan nama baru pembeli:")
		fmt.Print("> ")
		fmt.Scan(&newTransaction.BuyerName)

		utils.ClearScreen()
		fmt.Println("Konfirmasi perubahan transaksi:")
		fmt.Printf("🔢 Jumlah %d diubah menjadi %d\n", selectedTransaction.Quantity, newTransaction.Quantity)
		fmt.Printf("📒 Nama pembeli \"%s\" diubah menjadi \"%s\"\n", selectedTransaction.BuyerName, newTransaction.BuyerName)
		fmt.Println("Apa sudah benar? (y/n):")
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			TransactionList[selectedNumber-1] = newTransaction
			return
		} else if isConfirm == "n" || isConfirm == "N" {
			return
		} else {
			fmt.Print("> ")
			fmt.Scan(&isConfirm)
		}
	} else {
		utils.ShowEmptyTransactionList()
	}
}

func DeleteTransactionMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Transaksi", "Hapus")

	if IsTransactionExist() {
		var isConfirm string

		ShowTransactionList()
		fmt.Println("Masukkan nomor transaksi yang ingin dihapus")
		selectedNumber := utils.InputMenu(len(TransactionList))
		selectedTransaction := TransactionList[selectedNumber-1]

		fmt.Printf("Transaksi nomor %d\n", selectedNumber)
		fmt.Printf("Nama: %s\n", selectedTransaction.Item.Name)
		fmt.Printf("Jumlah: %d\n", selectedTransaction.Quantity)
		fmt.Printf("Pembeli: %s\n", selectedTransaction.BuyerName)
		fmt.Println("--------------------------")
		fmt.Println("Yakin ingin menghapus transaksi? (y/n):")
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			DeleteTransactionByIndex(selectedNumber - 1)
			return
		} else if isConfirm == "n" || isConfirm == "N" {
			return
		} else {
			fmt.Print("> ")
			fmt.Scan(&isConfirm)
		}
	} else {
		utils.ShowEmptyTransactionList()
	}
}

func ShowTransactionMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Transaksi", "Lihat")

	if IsTransactionExist() {
		ShowTransactionList()
		fmt.Println("Klik enter untuk kembali ...")
		fmt.Scanln()
	} else {
		utils.ShowEmptyTransactionList()
	}
}
