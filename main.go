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
		fmt.Print("Pilih (1")
		for i := 2; i <= maks; i++ {
			fmt.Printf("/%d", i)
		}
		fmt.Println(")")
		fmt.Print("> ")
		fmt.Scanln(&pilihan)

		if pilihan > maks || pilihan < 1 {
			fmt.Println("❌ Pilihan tidak valid")
		} else {
			return pilihan
		}
	}
}

func hapusBarangDenganIndex(index int) {
	var listBarangBaru TabBarang

	for i := 0; i < len(listBarang); i++ {
		if i != index {
			listBarangBaru = append(listBarangBaru, listBarang[i])
		}
	}

	listBarang = listBarangBaru
}

func cekAdaBarang() bool {
	return len(listBarang) > 0
}

func tampilkanListBarang() {
	for index, barang := range listBarang {
		nomor := index + 1
		fmt.Printf("%d. %s seharga %.f dengan kategori %s\n", nomor, barang.nama, barang.harga, barang.kategori)
	}
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
		fmt.Println("🔠 Masukkan nama Barang:")
		fmt.Print("> ")
		fmt.Scan(&barang.nama)

		fmt.Println("💰 Masukkan harga Barang:")
		fmt.Print("> ")
		fmt.Scan(&barang.harga)

		fmt.Println("🏷️  Masukkan kategori Barang:")
		fmt.Print("> ")
		fmt.Scan(&barang.kategori)

		fmt.Println("Konfirmasi Barang:")
		fmt.Printf("%s seharga %.f dengan kategori %s\n", barang.nama, barang.harga, barang.kategori)
		fmt.Println("Apa sudah benar? (y/n):")
		fmt.Print("> ")
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
		fmt.Println("Pilih nomor Barang yang ingin diedit:")
		fmt.Print("> ")
		fmt.Scan(&pilihan)

		barangBaru = listBarang[pilihan-1]
		barangLama = listBarang[pilihan-1]
		fmt.Println("🔠 Masukkan nama baru Barang:")
		fmt.Print("> ")
		fmt.Scan(&barangBaru.nama)

		fmt.Println("💰 Masukkan harga baru Barang:")
		fmt.Print("> ")
		fmt.Scan(&barangBaru.harga)

		fmt.Println("🏷️  Masukkan kategori baru Barang:")
		fmt.Print("> ")
		fmt.Scan(&barangBaru.kategori)

		fmt.Println("Konfirmasi perubahan:")
		fmt.Printf("\"%s\" diubah menjadi \"%s\"\n", barangLama.nama, barangBaru.nama)
		fmt.Printf("%.f diubah menjadi %.f\n", barangLama.harga, barangBaru.harga)
		fmt.Printf("\"%s\" diubah menjadi \"%s\"\n", barangLama.kategori, barangBaru.kategori)
		fmt.Println("Apa sudah benar? (y/n):")
		fmt.Print("> ")
		fmt.Scan(&benar)

		if benar == "y" || benar == "Y" {
			listBarang[pilihan-1] = barangBaru
			return
		}
	} else {
		tampilkanBarangKosong()
	}
}

func hapusBarang() {
	clearScreen()
	var barang Barang
	var pilihan int
	var benar string

	if cekAdaBarang() {
		tampilkanListBarang()
		fmt.Println("Pilih nomor Barang yang ingin dihapus:")
		fmt.Print("> ")
		fmt.Scan(&pilihan)
		barang = listBarang[pilihan-1]

		fmt.Printf("Yakin ingin menghapus \"%s\"? (y/n):\n", barang.nama)
		fmt.Print("> ")
		fmt.Scan(&benar)

		if benar == "y" || benar == "Y" {
			hapusBarangDenganIndex(pilihan - 1)
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
		fmt.Println("Klik Enter untuk kembali")
		fmt.Scanln()
	} else {
		tampilkanBarangKosong()
	}
}

func cariBarang() {
	clearScreen()

	if cekAdaBarang() {
		var namaCari string
		var ketemu bool

		fmt.Println("Masukkan nama barang (case sensitive):")
		fmt.Print("> ")
		fmt.Scan(&namaCari)

		clearScreen()
		for index, barang := range listBarang {
			if namaCari == barang.nama {
				ketemu = true
				fmt.Println("Hasil pencarian barang:")
				fmt.Printf("Nomor: %d\n", index+1)
				fmt.Printf("Nama: %s\n", barang.nama)
				fmt.Printf("Harga: %.f\n", barang.harga)
				fmt.Printf("Kategori: %s\n", barang.kategori)
				fmt.Println("--------------------------")
			}
		}

		if !ketemu {
			fmt.Println("Tidak ada barang yang ditemukan")
		}
		fmt.Println("Klik Enter untuk kembali")
		fmt.Scanln()
	} else {
		tampilkanBarangKosong()
	}
}

func menuBarang() {
	for {
		clearScreen()
		listMenuBarang()
		pilihan := inputMenu(6)

		switch pilihan {
		case 1:
			tambahBarang()
		case 2:
			editBarang()
		case 3:
			hapusBarang()
		case 4:
			lihatBarang()
		case 5:
			cariBarang()
		case 6:
			return
		}
	}
}

func menuTransaksi() {
	for {
		clearScreen()
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
	for {
		clearScreen()
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
