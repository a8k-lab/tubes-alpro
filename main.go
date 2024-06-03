package main

import (
	"fmt"
	"os"
	"os/exec"
)

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

func menuBarang() {
	for {
		listMenuBarang()
		pilihan := inputMenu(6)

		switch pilihan {
		case 1:

		case 2:

		case 3:

		case 4:

		case 5:

		case 6:
			clearScreen()
			return
		}
	}
}

func menuTransaksi() {
	for {
		listMenuTransaksi()
		pilihan := inputMenu(6)

		switch pilihan {
		case 1:

		case 2:

		case 3:

		case 4:

		case 5:
			clearScreen()
			return
		}
	}
}

func menuUtama() {
	for {
		listMenuUtama()
		pilihan := inputMenu(3)

		switch pilihan {
		case 1:
			clearScreen()
			menuBarang()
		case 2:
			clearScreen()
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
