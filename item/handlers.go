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

var ItemList TabItem

func IsItemExist() bool {
	return len(ItemList) > 0
}

func DeleteItemByIndex(index int) {
	var newItemList TabItem

	for i := 0; i < len(ItemList); i++ {
		if i != index {
			newItemList = append(newItemList, ItemList[i])
		}
	}

	ItemList = newItemList
}

func ShowItemList() {
	for index, item := range ItemList {
		number := index + 1
		fmt.Printf("%d. %s seharga %.f dengan kategori %s\n", number, item.Name, item.Price, item.Category)
	}
	fmt.Println("--------------------------")
}

func AddItemMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Barang", "Tambah")

	for {
		var newItem Item
		var isConfirm string

		fmt.Println("🔠 Masukkan nama barang (tidak boleh ada spasi):")
		fmt.Print("> ")
		fmt.Scan(&newItem.Name)

		fmt.Println("💰 Masukkan harga barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.Price)

		fmt.Println("🏷️  Masukkan kategori barang (tidak boleh ada spasi):")
		fmt.Print("> ")
		fmt.Scan(&newItem.Category)

		fmt.Println("Konfirmasi barang:")
		fmt.Printf("%s seharga %.f dengan kategori %s\n", newItem.Name, newItem.Price, newItem.Category)
		fmt.Println("Apa sudah benar? (y/n):")
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			ItemList = append(ItemList, newItem)
			return
		} else if isConfirm == "n" || isConfirm == "N" {
			return
		} else {
			fmt.Print("> ")
			fmt.Scan(&isConfirm)
		}
	}
}

func EditItemMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Barang", "Edit")

	if IsItemExist() {
		var isConfirm string

		ShowItemList()
		fmt.Println("Masukkan nomor barang yang ingin diedit")
		selectedNumber := utils.InputMenu(len(ItemList))
		newItem := ItemList[selectedNumber-1]
		selectedItem := ItemList[selectedNumber-1]

		fmt.Println("🔠 Masukkan nama baru barang (tidak boleh ada spasi):")
		fmt.Print("> ")
		fmt.Scan(&newItem.Name)

		fmt.Println("💰 Masukkan harga baru barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.Price)

		fmt.Println("🏷️  Masukkan kategori baru barang (tidak boleh ada spasi):")
		fmt.Print("> ")
		fmt.Scan(&newItem.Category)

		utils.ClearScreen()
		fmt.Println("Konfirmasi perubahan barang:")
		fmt.Printf("🔠 Nama \"%s\" diubah menjadi \"%s\"\n", selectedItem.Name, newItem.Name)
		fmt.Printf("💰 Harga %.f diubah menjadi %.f\n", selectedItem.Price, newItem.Price)
		fmt.Printf("🏷️ Kategori \"%s\" diubah menjadi \"%s\"\n", selectedItem.Category, newItem.Category)
		fmt.Println("Apa sudah benar? (y/n):")
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			ItemList[selectedNumber-1] = newItem
			return
		} else if isConfirm == "n" || isConfirm == "N" {
			return
		} else {
			fmt.Print("> ")
			fmt.Scan(&isConfirm)
		}
	} else {
		utils.ShowEmptyItemList()
	}
}

func DeleteItemMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Barang", "Hapus")

	if IsItemExist() {
		var isConfirm string

		ShowItemList()
		fmt.Println("Masukkan nomor barang yang ingin dihapus")
		selectedNumber := utils.InputMenu(len(ItemList))
		selectedItem := ItemList[selectedNumber-1]

		fmt.Printf("Yakin ingin menghapus \"%s\"? (y/n):\n", selectedItem.Name)
		fmt.Print("> ")
		fmt.Scan(&isConfirm)

		if isConfirm == "y" || isConfirm == "Y" {
			DeleteItemByIndex(selectedNumber - 1)
			return
		} else if isConfirm == "n" || isConfirm == "N" {
			return
		} else {
			fmt.Print("> ")
			fmt.Scan(&isConfirm)
		}
	} else {
		utils.ShowEmptyItemList()
	}
}

func ShowItemMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Barang", "Lihat")

	if IsItemExist() {
		ShowItemList()
		fmt.Println("Klik enter untuk kembali ...")
		fmt.Scanln()
	} else {
		utils.ShowEmptyItemList()
	}
}

func SearchItemMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Barang", "Cari")

	if IsItemExist() {
		var keyword string
		var isFound bool

		fmt.Println("Masukkan nama barang (case sensitive):")
		fmt.Print("> ")
		fmt.Scan(&keyword)

		utils.ClearScreen()
		for index, item := range ItemList {
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
		fmt.Println("Klik enter untuk kembali ...")
		fmt.Scanln()
	} else {
		utils.ShowEmptyItemList()
	}
}
