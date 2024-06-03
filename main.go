package main

import "fmt"

func menuIntro() {
	fmt.Println("#------------------------#")
	fmt.Println("👋 Selamat Datang")
	fmt.Println("Aplikasi Jual Beli Barang")
	fmt.Println("#------------------------#")
}

func menuUtama(pilihan *int) {
	fmt.Println("#------------------------#")
	fmt.Println("#         M E N U        #")
	fmt.Println("#------------------------#")
	fmt.Println("1. Barang 📦              ")
	fmt.Println("2. Transaksi 💵           ")
	fmt.Println("3. Keluar ⛔              ")
	fmt.Println("--------------------------")
	fmt.Print("Pilih (1/2/3): ")
	fmt.Scan(pilihan)
}

func menuBarang() {
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
	fmt.Print("Pilih (1/2/3/4/5/6): ")
}

func menuTransaksi() {
	fmt.Println("#------------------------#")
	fmt.Println("#   Menu Transaksi 💵   #")
	fmt.Println("#------------------------#")
	fmt.Println("1. Tambah ➕")
	fmt.Println("2. Edit ✏️")
	fmt.Println("3. Hapus 🗑️")
	fmt.Println("4. Lihat 📊")
	fmt.Println("5. Kembali 🔙")
	fmt.Println("--------------------------")
	fmt.Print("Pilih (1/2/3/4/5): ")
}

func main() {
	var pilihanUtama int

	menuIntro()

	// 1. Prosedur untuk menampilkan menu utama (Barang, Transaksi), dan pilihan exit program
	menuUtama(&pilihanUtama)

	switch pilihanUtama {
	case 1:
		menuBarang()
	case 2:
		menuTransaksi()
	case 3:
		fmt.Println("Menu Berakhir")
	}
}
