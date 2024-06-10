package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearScreen() {
	switch runtime.GOOS {
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
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
