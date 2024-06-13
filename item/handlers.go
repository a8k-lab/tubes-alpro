package item

import (
	"fmt"
	"jual-beli/utils"
)

type Item struct {
	Name         string
	CapitalPrice float64
	SalePrice    float64
	Category     string
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
		fmt.Printf("%d. %s dengan harga modal %.f, seharga %.f ada di kategori %s\n", number, item.Name, item.CapitalPrice, item.SalePrice, item.Category)
	}
	fmt.Println("--------------------------")
}

func AddItemMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Barang", "Tambah")

	for {
		var newItem Item
		var isConfirm bool

		fmt.Println("🔠 Masukkan nama barang (tidak boleh ada spasi):")
		fmt.Print("> ")
		fmt.Scan(&newItem.Name)

		fmt.Println("💵 Masukkan harga modal barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.CapitalPrice)

		fmt.Println("💶 Masukkan harga jual barang:")
		fmt.Print("> ")
		fmt.Scan(&newItem.SalePrice)

		fmt.Println("🏷️  Masukkan kategori barang (tidak boleh ada spasi):")
		fmt.Print("> ")
		fmt.Scan(&newItem.Category)

		utils.ClearScreen()
		fmt.Println("Konfirmasi penambahan barang:")
		fmt.Printf("%s dengan harga modal %.f, seharga %.f ada di kategori %s\n", newItem.Name, newItem.CapitalPrice, newItem.SalePrice, newItem.Category)
		utils.ConfirmInput(&isConfirm, "menambahkan barang")

		if isConfirm {
			ItemList = append(ItemList, newItem)
			utils.PrintSuccessMessage("Barang berhasil ditambahkan")
			return
		} else {
			return
		}
	}
}

func EditItemMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Barang", "Edit")

	if IsItemExist() {
		var newItem Item
		var selectedNumber int
		var isConfirm bool

		ShowItemList()
		fmt.Println("Masukkan nomor barang yang ingin diedit")
		utils.InputMenu(&selectedNumber, len(ItemList))
		newItem = ItemList[selectedNumber-1]
		selectedItem := ItemList[selectedNumber-1]

		utils.ClearScreen()
		fmt.Printf("Edit data barang \"%s\"\n", selectedItem.Name)
		fmt.Println("Kosongkan input jika tidak ingin diubah")
		fmt.Println("--------------------------")

		fmt.Println("🔠 Masukkan nama baru barang (tidak boleh ada spasi):")
		fmt.Print("> ")
		fmt.Scanln(&newItem.Name)

		fmt.Println("💵 Masukkan harga baru modal barang:")
		fmt.Print("> ")
		fmt.Scanln(&newItem.CapitalPrice)

		fmt.Println("💶 Masukkan harga baru jual barang:")
		fmt.Print("> ")
		fmt.Scanln(&newItem.SalePrice)

		fmt.Println("🏷️ Masukkan kategori baru barang (tidak boleh ada spasi):")
		fmt.Print("> ")
		fmt.Scanln(&newItem.Category)

		utils.ClearScreen()
		fmt.Println("Konfirmasi perubahan barang:")
		fmt.Printf("🔠 Nama \"%s\" diubah menjadi \"%s\"\n", selectedItem.Name, newItem.Name)
		fmt.Printf("💵 Harga modal %.f diubah menjadi %.f\n", selectedItem.CapitalPrice, newItem.CapitalPrice)
		fmt.Printf("💶 Harga jual %.f diubah menjadi %.f\n", selectedItem.SalePrice, newItem.SalePrice)
		fmt.Printf("🏷️ Kategori \"%s\" diubah menjadi \"%s\"\n", selectedItem.Category, newItem.Category)
		utils.ConfirmInput(&isConfirm, "mengedit barang")

		if isConfirm {
			ItemList[selectedNumber-1] = newItem
			utils.PrintSuccessMessage("Barang berhasil diubah")
		}
		return
	} else {
		utils.ShowEmptyItemList()
	}
}

func DeleteItemMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Barang", "Hapus")

	if IsItemExist() {
		var selectedNumber int
		var isConfirm bool

		ShowItemList()
		fmt.Println("Masukkan nomor barang yang ingin dihapus")
		utils.InputMenu(&selectedNumber, len(ItemList))
		selectedItem := ItemList[selectedNumber-1]

		utils.ClearScreen()
		fmt.Println("Konfirmasi penghapusan barang:")
		fmt.Printf("Nomor: %d\n", selectedNumber)
		fmt.Printf("Nama: %s\n", selectedItem.Name)
		fmt.Printf("Harga modal: %.f\n", selectedItem.CapitalPrice)
		fmt.Printf("Harga jual: %.f\n", selectedItem.SalePrice)
		fmt.Printf("Kategori: %s\n", selectedItem.Category)
		utils.ConfirmInput(&isConfirm, "menghapus barang")

		if isConfirm {
			DeleteItemByIndex(selectedNumber - 1)
			utils.PrintSuccessMessage("Barang berhasil dihapus")
		}
		return
	} else {
		utils.ShowEmptyItemList()
	}
}

func ShowItemMenu() {
	utils.ClearScreen()
	utils.PrintBreadcrumb("Menu", "Barang", "Lihat")

	if IsItemExist() {
		ShowItemList()
		fmt.Println("Klik Enter untuk kembali ...")
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
				fmt.Printf("Harga modal: %.f\n", item.CapitalPrice)
				fmt.Printf("Harga jual: %.f\n", item.SalePrice)
				fmt.Printf("Kategori: %s\n", item.Category)
				fmt.Println("--------------------------")
			}
		}

		if !isFound {
			fmt.Println("Tidak ada barang yang ditemukan")
		}
		fmt.Println("Klik Enter untuk kembali ...")
		fmt.Scanln()
	} else {
		utils.ShowEmptyItemList()
	}
}
