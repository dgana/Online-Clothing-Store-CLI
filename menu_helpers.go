package main

import (
	"fmt"
	"strings"
)

func printRoleMenu() {
	fmt.Println("\n=== ZYGAR ONLINE CLOTHING STORE CLI ===")
	fmt.Println("1. Admin")
	fmt.Println("2. Client")
	fmt.Println("0. Keluar")
}

func isValidRole(role int) bool {
	return role == 1 || role == 2
}

func printMenu(role int) {
	fmt.Println("\n=== MENU ===")
	if role == 1 {
		fmt.Println("1. Tambah User")
		fmt.Println("2. Tambah Produk")
		fmt.Println("3. Buat Order")
		fmt.Println("4. Tambah Item Order")
		fmt.Println("5. Laporan User")
		fmt.Println("6. Laporan Order")
		fmt.Println("7. Laporan Stok")
		fmt.Println("8. Delete User")
		fmt.Println("9. Delete Product")
		fmt.Println("10. Update User")
		fmt.Println("11. Kembali")
	}
	if role == 2 {
		fmt.Println("1. Buat Order")
		fmt.Println("2. Tambah Item Order")
		fmt.Println("3. Laporan User")
		fmt.Println("4. Laporan Order")
		fmt.Println("5. Laporan Stok")
		fmt.Println("6. Kembali")
	}
	fmt.Println("0. Keluar")
}

func shouldBackToRoleMenu(role, pilihan int) bool {
	return (role == 1 && pilihan == 11) || (role == 2 && pilihan == 6)
}

func mapClientChoiceToAdminChoice(pilihan int) (int, bool) {
	switch pilihan {
	case 1:
		return 3, true
	case 2:
		return 4, true
	case 3:
		return 5, true
	case 4:
		return 6, true
	case 5:
		return 7, true
	default:
		return 0, false
	}
}

func waitEnterToContinue() {
	_ = readLine("Enter untuk kembali...")
}

func getUserIDByUsernameOrPrint(username string) int {
	username = strings.TrimSpace(username)
	userID, err := getUserIDByUsername(username)
	if err != nil {
		fmt.Println("User tidak ditemukan.")
		return 0
	}
	return userID
}

func getProductIDByCodeOrPrint(code string) int {
	code = strings.TrimSpace(code)
	productID, err := getProductIDByCode(code)
	if err != nil {
		fmt.Println("Produk tidak ditemukan.")
		return 0
	}
	return productID
}
