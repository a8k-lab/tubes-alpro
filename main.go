package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Barang struct {
	nama     string
	harga    float64
	kategori string
}

type TabBarang []Barang

var listBarang TabBarang

func inputMenu(maks int) int {
	var pilihan int

	for {
		fmt.Print("Pilih (1/2/3): ")
		fmt.Scan(&pilihan)

		if pilihan > maks || pilihan < 1 {
			fmt.Println("❌ Pilihan tidak valid")
		} else {
			return pilihan
		}
	}
}

func cekAdaBarang() bool {
	return len(listBarang) > 0
}

func tampilkanListBarang() {
	for index, barang := range listBarang {
		nomor := index + 1
		fmt.Printf("%d. %s seharga %.f dengan kategori %s\n", nomor, barang.nama, barang.harga, barang.kategori)
	}
	fmt.Println("Klik Enter untuk kembali")
	fmt.Scanln()
}

func tampilkanBarangKosong() {
	fmt.Println("Belum ada barang yang tersimpan")
	fmt.Println("Klik Enter untuk kembali")
	fmt.Scanln()
}

func listMenuBarang() {
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

func listMenuTransaksi() {
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

func listMenuUtama() {
	fmt.Println("#------------------------#")
	fmt.Println("#         M E N U        #")
	fmt.Println("#------------------------#")
	fmt.Println("1. Barang 📦              ")
	fmt.Println("2. Transaksi 💵           ")
	fmt.Println("3. Keluar ⛔              ")
	fmt.Println("--------------------------")
}

func tambahBarang() {
	clearScreen()
	var barang Barang
	var benar string

	for {
		fmt.Print("🔠 Masukkan nama Barang: ")
		fmt.Scan(&barang.nama)

		fmt.Print("💰 Masukkan harga Barang: ")
		fmt.Scan(&barang.harga)

		fmt.Print("🏷️  Masukkan kategori Barang: ")
		fmt.Scan(&barang.kategori)

		fmt.Println("Konfirmasi Barang:")
		fmt.Printf("%s seharga %.f dengan kategori %s\n", barang.nama, barang.harga, barang.kategori)
		fmt.Print("Apa sudah benar? (y/n): ")

		fmt.Scan(&benar)
		if benar == "y" || benar == "Y" {
			listBarang = append(listBarang, barang)
			return
		}
	}
}

func editBarang() {
	clearScreen()
	var barangBaru, barangLama Barang
	var pilihan int
	var benar string

	if cekAdaBarang() {
		tampilkanListBarang()
		fmt.Print("Pilih Barang yang ingin diedit berdasarkan nomornya: ")
		fmt.Scan(&pilihan)

		barangBaru = listBarang[pilihan-1]
		barangLama = listBarang[pilihan-1]
		fmt.Print("🔠 Masukkan nama baru Barang: ")
		fmt.Scan(&barangBaru.nama)

		fmt.Print("💰 Masukkan harga baru Barang: ")
		fmt.Scan(&barangBaru.harga)

		fmt.Print("🏷️  Masukkan kategori baru Barang: ")
		fmt.Scan(&barangBaru.kategori)

		fmt.Println("Konfirmasi perubahan:")
		fmt.Printf("\"%s\" diubah menjadi \"%s\"\n", barangLama.nama, barangBaru.nama)
		fmt.Printf("%.f diubah menjadi %.f\n", barangLama.harga, barangBaru.harga)
		fmt.Printf("\"%s\" diubah menjadi \"%s\"\n", barangLama.kategori, barangBaru.kategori)
		fmt.Print("Apa sudah benar? (y/n): ")

		fmt.Scan(&benar)
		if benar == "y" || benar == "Y" {
			listBarang[pilihan-1] = barangBaru
			return
		}
	} else {
		tampilkanBarangKosong()
	}
}

func lihatBarang() {
	clearScreen()

	if cekAdaBarang() {
		tampilkanListBarang()
	} else {
		tampilkanBarangKosong()
	}
}

func menuBarang() {
	clearScreen()

	for {
		listMenuBarang()
		pilihan := inputMenu(6)

		switch pilihan {
		case 1:
			tambahBarang()
		case 2:
			editBarang()
		case 3:

		case 4:
			lihatBarang()
		case 5:

		case 6:
			return
		}
	}
}

func menuTransaksi() {
	clearScreen()

	for {
		listMenuTransaksi()
		pilihan := inputMenu(6)

		switch pilihan {
		case 1:

		case 2:

		case 3:

		case 4:

		case 5:
			return
		}
	}
}

func menuUtama() {
	clearScreen()

	for {
		listMenuUtama()
		pilihan := inputMenu(3)

		switch pilihan {
		case 1:
			menuBarang()
		case 2:
			menuTransaksi()
		case 3:
			end()
		}
	}
}

func main() {
	clearScreen()
	menuUtama()
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
