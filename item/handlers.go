package item

import (
	"fmt"
	"jual-beli/utils"
)

type Item struct {
	Name     string
	Price    float64
	Category string
}

type TabItem []Item

var itemList TabItem

func IsItemExist() bool {
	return len(itemList) > 0
}

func DeleteItemByIndex(index int) {
	var newItemList TabItem

	for i := 0; i < len(itemList); i++ {
		if i != index {
			newItemList = append(newItemList, itemList[i])
		}
	}

	itemList = newItemList
}

func AddItem() {
	utils.ClearScreen()

	for {
		var newItem Item
		var isConfirm string

		fmt.Println("🔠 Masukkan nama Barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.Name)

		fmt.Println("💰 Masukkan harga Barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.Price)

		fmt.Println("🏷️  Masukkan kategori Barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.Category)

		fmt.Println("Konfirmasi Barang:")
		fmt.Printf("%s seharga %.f dengan kategori %s\n", newItem.Name, newItem.Price, newItem.Category)
		fmt.Println("Apa sudah benar? (y/n):")
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			itemList = append(itemList, newItem)
			return
		}
	}
}

func EditItem() {
	utils.ClearScreen()

	if IsItemExist() {
		var newItem, selectedItem Item
		var selectedNumber int
		var isConfirm string

		ShowItemList()
		fmt.Println("Pilih nomor Barang yang ingin diedit:")
		fmt.Print("> ")
		fmt.Scan(&selectedNumber)

		newItem = itemList[selectedNumber-1]
		selectedItem = itemList[selectedNumber-1]
		fmt.Println("🔠 Masukkan nama baru Barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.Name)

		fmt.Println("💰 Masukkan harga baru Barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.Price)

		fmt.Println("🏷️  Masukkan kategori baru Barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.Category)

		fmt.Println("Konfirmasi perubahan:")
		fmt.Printf("\"%s\" diubah menjadi \"%s\"\n", selectedItem.Name, newItem.Name)
		fmt.Printf("%.f diubah menjadi %.f\n", selectedItem.Price, newItem.Price)
		fmt.Printf("\"%s\" diubah menjadi \"%s\"\n", selectedItem.Category, newItem.Category)
		fmt.Println("Apa sudah benar? (y/n):")
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			itemList[selectedNumber-1] = newItem
			return
		}
	} else {
		utils.ShowEmptyItemList()
	}
}

func DeleteItem() {
	utils.ClearScreen()

	if IsItemExist() {
		var selectedItem Item
		var selectedNumber int
		var isConfirm string

		ShowItemList()
		fmt.Println("Pilih nomor Barang yang ingin dihapus:")
		fmt.Print("> ")
		fmt.Scan(&selectedNumber)
		selectedItem = itemList[selectedNumber-1]

		fmt.Printf("Yakin ingin menghapus \"%s\"? (y/n):\n", selectedItem.Name)
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			DeleteItemByIndex(selectedNumber - 1)
			return
		}
	} else {
		utils.ShowEmptyItemList()
	}
}

func ShowItemList() {
	for index, item := range itemList {
		number := index + 1
		fmt.Printf("%d. %s seharga %.f dengan kategori %s\n", number, item.Name, item.Price, item.Category)
	}
}

func ShowItem() {
	utils.ClearScreen()

	if IsItemExist() {
		ShowItemList()
		fmt.Println("Klik Enter untuk kembali")
		fmt.Scanln()
	} else {
		utils.ShowEmptyItemList()
	}
}

func SearchItem() {
	utils.ClearScreen()

	if IsItemExist() {
		var keyword string
		var isFound bool

		fmt.Println("Masukkan nama barang (case sensitive):")
		fmt.Print("> ")
		fmt.Scan(&keyword)

		utils.ClearScreen()
		for index, item := range itemList {
			if keyword == item.Name {
				isFound = true
				fmt.Println("Hasil pencarian barang:")
				fmt.Printf("Nomor: %d\n", index+1)
				fmt.Printf("Nama: %s\n", item.Name)
				fmt.Printf("Harga: %.f\n", item.Price)
				fmt.Printf("Kategori: %s\n", item.Category)
				fmt.Println("--------------------------")
			}
		}

		if !isFound {
			fmt.Println("Tidak ada barang yang ditemukan")
		}
		fmt.Println("Klik Enter untuk kembali")
		fmt.Scanln()
	} else {
		utils.ShowEmptyItemList()
	}
}
