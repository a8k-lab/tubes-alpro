package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func ShowEmptyItemList() {
	fmt.Println("Belum ada barang yang tersimpan")
	fmt.Println("Klik Enter untuk kembali")
	fmt.Scanln()
}
