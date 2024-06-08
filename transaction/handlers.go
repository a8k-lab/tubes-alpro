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
		fmt.Printf("%d. %s x %d\n", number, item.Item.Name, item.Quantity)
	}
	fmt.Println("--------------------------")
}

func AddTransactionMenu() {
	utils.ClearScreen()

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
		fmt.Println("Klik Enter untuk kembali")
		fmt.Scanln()
	} else {
		fmt.Println("Tidak dapat menambahkan transaksi")
		utils.ShowEmptyItemList()
	}
}

func EditTransactionMenu() {
	utils.ClearScreen()

	if IsTransactionExist() {
		var newTransaction, selectedTransaction Transaction
		var isConfirm string

		ShowTransactionList()
		fmt.Println("Masukkan nomor transaksi yang ingin diedit")
		selectedNumber := utils.InputMenu(len(TransactionList))
		selectedTransaction = TransactionList[selectedNumber-1]

		fmt.Println("🔢 Masukkan jumlah baru barang:")
		fmt.Print("> ")
		fmt.Scan(&newTransaction.Quantity)

		fmt.Println("📒 Masukkan nama baru pembeli:")
		fmt.Print("> ")
		fmt.Scan(&newTransaction.BuyerName)

		fmt.Println("Konfirmasi perubahan:")
		fmt.Printf("- %d diubah menjadi %d\n", selectedTransaction.Quantity, newTransaction.Quantity)
		fmt.Printf("- \"%s\" diubah menjadi \"%s\"\n", selectedTransaction.BuyerName, newTransaction.BuyerName)
		fmt.Println("Apa sudah benar? (y/n):")
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			TransactionList[selectedNumber-1] = newTransaction
			return
		}
	} else {
		utils.ShowEmptyTransactionList()
	}
}

func DeleteTransactionMenu() {
	utils.ClearScreen()

	if IsTransactionExist() {
		var selectedTransaction Transaction
		var isConfirm string

		ShowTransactionList()
		fmt.Println("Masukkan nomor transaksi yang ingin dihapus")
		selectedNumber := utils.InputMenu(len(TransactionList))
		selectedTransaction = TransactionList[selectedNumber-1]

		fmt.Printf("Transaksi nomor %d\n", selectedNumber)
		fmt.Printf("Nama: %s\n", selectedTransaction.Item.Name)
		fmt.Printf("Jumlah: %d\n", selectedTransaction.Quantity)
		fmt.Printf("Pembeli: %s\n", selectedTransaction.BuyerName)
		fmt.Println("--------------------------")
		fmt.Println("Yakin ingin menghapus transaksi?")
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			DeleteTransactionByIndex(selectedNumber - 1)
			return
		}
	} else {
		utils.ShowEmptyTransactionList()
	}
}

func ShowTransactionMenu() {
	utils.ClearScreen()

	if IsTransactionExist() {
		ShowTransactionList()
		fmt.Println("Klik Enter untuk kembali")
		fmt.Scanln()
	} else {
		utils.ShowEmptyTransactionList()
	}
}
