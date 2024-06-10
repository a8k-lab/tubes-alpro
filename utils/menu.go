package utils

import "fmt"

func InputMenu(selectedNumber *int, max int) {
	for {
		fmt.Print("Pilih (1")
		for i := 2; i <= max; i++ {
			fmt.Printf("/%d", i)
		}
		fmt.Println(")")
		fmt.Print("> ")
		fmt.Scanln(selectedNumber)

		if *selectedNumber > max || *selectedNumber < 1 {
			fmt.Println("❌ Input tidak valid")
		} else {
			return
		}
	}
}

func PrintIntro() {
	ClearScreen()
	fmt.Println("👋 Selamat datang di aplikasi Jual Beli")
	fmt.Println("🎨 Crafted by Abdul Malik & Putri Galuh Mandarizka")
	fmt.Println("😼 GitHub: https://github.com/a8k-lab/tubes-alpro")
	fmt.Println("---------------------------------------------------")
	fmt.Println("▶️ Klik Enter untuk lanjut ...")
	fmt.Scanln()
}

func PrintMainMenu() {
	PrintBreadcrumb("Menu")
	fmt.Println("1. Barang 📦              ")
	fmt.Println("2. Transaksi 💵           ")
	fmt.Println("3. Keluar ⛔              ")
	fmt.Println("--------------------------")
}

func PrintItemMenu() {
	PrintBreadcrumb("Menu", "Barang")
	fmt.Println("1. Tambah ➕")
	fmt.Println("2. Edit ✏️")
	fmt.Println("3. Hapus 🗑️")
	fmt.Println("4. Lihat 📊")
	fmt.Println("5. Cari 🔍")
	fmt.Println("6. Kembali 🔙")
	fmt.Println("--------------------------")
}

func PrintTransactionMenu() {
	PrintBreadcrumb("Menu", "Transaksi")
	fmt.Println("1. Tambah ➕")
	fmt.Println("2. Edit ✏️")
	fmt.Println("3. Hapus 🗑️")
	fmt.Println("4. Lihat 📊")
	fmt.Println("5. Kembali 🔙")
	fmt.Println("--------------------------")
}

func PrintBreadcrumb(menus ...string) {
	for i := 0; i < len(menus)-1; i++ {
		fmt.Printf("%s > ", menus[i])
	}
	fmt.Println(menus[len(menus)-1])
	fmt.Println("--------------------------")
}

func ShowEmptyItemList() {
	fmt.Println("Belum ada barang yang tersimpan")
	fmt.Println("Klik Enter untuk kembali ...")
	fmt.Scanln()
}

func ShowEmptyTransactionList() {
	fmt.Println("Belum ada transaksi yang tersimpan")
	fmt.Println("Klik Enter untuk kembali ...")
	fmt.Scanln()
}

func ConfirmInput(isConfirm *bool, additionalMessage string) {
	var input string

	fmt.Println("--------------------------")
	fmt.Printf("Yakin ingin %s? (y/n):\n", additionalMessage)

	for {
		fmt.Print("> ")
		fmt.Scan(&input)

		if input == "y" || input == "Y" {
			*isConfirm = true
			return
		} else if input == "n" || input == "N" {
			*isConfirm = false
			return
		} else {
			fmt.Println("❌ Input tidak valid")
		}
	}
}
