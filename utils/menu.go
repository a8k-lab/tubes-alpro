package utils

import "fmt"

func InputMenu(max int) int {
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

func PrintIntro() {
	ClearScreen()
	fmt.Println("👋 Selamat datang di aplikasi Jual Beli")
	fmt.Println("🎨 Crafted by Abdul Malik & Putri Galuh Mandarizka")
	fmt.Println("😼 GitHub: https://github.com/a8k-lab/tubes-alpro")
	fmt.Println("---------------------------------------------------")
	fmt.Println("📦 Klik Enter untuk lanjut ...")
	fmt.Scanln()
}

func PrintMainMenu() {
	fmt.Println("#------------------------#")
	fmt.Println("#         M E N U        #")
	fmt.Println("#------------------------#")
	fmt.Println("1. Barang 📦              ")
	fmt.Println("2. Transaksi 💵           ")
	fmt.Println("3. Keluar ⛔              ")
	fmt.Println("--------------------------")
}

func PrintItemMenu() {
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

func PrintTransactionMenu() {
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
