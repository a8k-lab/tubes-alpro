package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Item struct {
	name     string
	price    float64
	category string
}

type TabItem []Item

var itemList TabItem

func inputMenu(max int) int {
	var selectedNumber int

	for {
		fmt.Print("Pilih (1")
		for i := 2; i <= max; i++ {
			fmt.Printf("/%d", i)
		}
		fmt.Println(")")
		fmt.Print("> ")
		fmt.Scanln(&selectedNumber)

		if selectedNumber > max || selectedNumber < 1 {
			fmt.Println("❌ Pilihan tidak valid")
		} else {
			return selectedNumber
		}
	}
}

func deleteItemByIndex(index int) {
	var newItemList TabItem

	for i := 0; i < len(itemList); i++ {
		if i != index {
			newItemList = append(newItemList, itemList[i])
		}
	}

	itemList = newItemList
}

func isItemExist() bool {
	return len(itemList) > 0
}

func showItemList() {
	for index, item := range itemList {
		number := index + 1
		fmt.Printf("%d. %s seharga %.f dengan kategori %s\n", number, item.name, item.price, item.category)
	}
}

func showEmptyItemList() {
	fmt.Println("Belum ada barang yang tersimpan")
	fmt.Println("Klik Enter untuk kembali")
	fmt.Scanln()
}

func printItemMenu() {
	fmt.Println("#------------------------#")
	fmt.Println("#     Menu Barang 📦     #")
	fmt.Println("#------------------------#")
	fmt.Println("1. Tambah ➕")
	fmt.Println("2. Edit ✏️")
	fmt.Println("3. Hapus 🗑️")
	fmt.Println("4. Lihat 📊")
	fmt.Println("5. Cari 🔍")
	fmt.Println("6. Kembali 🔙")
	fmt.Println("--------------------------")
}

func printTransactionMenu() {
	fmt.Println("#------------------------#")
	fmt.Println("#    Menu Transaksi 💵   #")
	fmt.Println("#------------------------#")
	fmt.Println("1. Tambah ➕")
	fmt.Println("2. Edit ✏️")
	fmt.Println("3. Hapus 🗑️")
	fmt.Println("4. Lihat 📊")
	fmt.Println("5. Kembali 🔙")
	fmt.Println("--------------------------")
}

func printMainMenu() {
	fmt.Println("#------------------------#")
	fmt.Println("#         M E N U        #")
	fmt.Println("#------------------------#")
	fmt.Println("1. Barang 📦              ")
	fmt.Println("2. Transaksi 💵           ")
	fmt.Println("3. Keluar ⛔              ")
	fmt.Println("--------------------------")
}

func addItem() {
	clearScreen()

	for {
		var newItem Item
		var isConfirm string

		fmt.Println("🔠 Masukkan nama Barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.name)

		fmt.Println("💰 Masukkan harga Barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.price)

		fmt.Println("🏷️  Masukkan kategori Barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.category)

		fmt.Println("Konfirmasi Barang:")
		fmt.Printf("%s seharga %.f dengan kategori %s\n", newItem.name, newItem.price, newItem.category)
		fmt.Println("Apa sudah benar? (y/n):")
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			itemList = append(itemList, newItem)
			return
		}
	}
}

func editItem() {
	clearScreen()

	if isItemExist() {
		var newItem, selectedItem Item
		var selectedNumber int
		var isConfirm string

		showItemList()
		fmt.Println("Pilih nomor Barang yang ingin diedit:")
		fmt.Print("> ")
		fmt.Scan(&selectedNumber)

		newItem = itemList[selectedNumber-1]
		selectedItem = itemList[selectedNumber-1]
		fmt.Println("🔠 Masukkan nama baru Barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.name)

		fmt.Println("💰 Masukkan harga baru Barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.price)

		fmt.Println("🏷️  Masukkan kategori baru Barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.category)

		fmt.Println("Konfirmasi perubahan:")
		fmt.Printf("\"%s\" diubah menjadi \"%s\"\n", selectedItem.name, newItem.name)
		fmt.Printf("%.f diubah menjadi %.f\n", selectedItem.price, newItem.price)
		fmt.Printf("\"%s\" diubah menjadi \"%s\"\n", selectedItem.category, newItem.category)
		fmt.Println("Apa sudah benar? (y/n):")
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			itemList[selectedNumber-1] = newItem
			return
		}
	} else {
		showEmptyItemList()
	}
}

func deleteItem() {
	clearScreen()

	if isItemExist() {
		var selectedItem Item
		var selectedNumber int
		var isConfirm string

		showItemList()
		fmt.Println("Pilih nomor Barang yang ingin dihapus:")
		fmt.Print("> ")
		fmt.Scan(&selectedNumber)
		selectedItem = itemList[selectedNumber-1]

		fmt.Printf("Yakin ingin menghapus \"%s\"? (y/n):\n", selectedItem.name)
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			deleteItemByIndex(selectedNumber - 1)
			return
		}
	} else {
		showEmptyItemList()
	}
}

func showItem() {
	clearScreen()

	if isItemExist() {
		showItemList()
		fmt.Println("Klik Enter untuk kembali")
		fmt.Scanln()
	} else {
		showEmptyItemList()
	}
}

func searchItem() {
	clearScreen()

	if isItemExist() {
		var keyword string
		var isFound bool

		fmt.Println("Masukkan nama barang (case sensitive):")
		fmt.Print("> ")
		fmt.Scan(&keyword)

		clearScreen()
		for index, item := range itemList {
			if keyword == item.name {
				isFound = true
				fmt.Println("Hasil pencarian barang:")
				fmt.Printf("Nomor: %d\n", index+1)
				fmt.Printf("Nama: %s\n", item.name)
				fmt.Printf("Harga: %.f\n", item.price)
				fmt.Printf("Kategori: %s\n", item.category)
				fmt.Println("--------------------------")
			}
		}

		if !isFound {
			fmt.Println("Tidak ada barang yang ditemukan")
		}
		fmt.Println("Klik Enter untuk kembali")
		fmt.Scanln()
	} else {
		showEmptyItemList()
	}
}

func itemMenu() {
	for {
		clearScreen()
		printItemMenu()
		selectMenu := inputMenu(6)

		switch selectMenu {
		case 1:
			addItem()
		case 2:
			editItem()
		case 3:
			deleteItem()
		case 4:
			showItem()
		case 5:
			searchItem()
		case 6:
			return
		}
	}
}

func transactionMenu() {
	for {
		clearScreen()
		printTransactionMenu()
		selectMenu := inputMenu(5)

		switch selectMenu {
		case 1:

		case 2:

		case 3:

		case 4:

		case 5:
			return
		}
	}
}

func mainMenu() {
	for {
		clearScreen()
		printMainMenu()
		selectMenu := inputMenu(3)

		switch selectMenu {
		case 1:
			itemMenu()
		case 2:
			transactionMenu()
		case 3:
			end()
		}
	}
}

func main() {
	clearScreen()
	mainMenu()
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func end() {
	clearScreen()
	fmt.Println("Program Berakhir")
	os.Exit(0)
}
