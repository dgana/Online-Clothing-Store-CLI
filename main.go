package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		printRoleMenu()
		role := readInt("Pilih role: ")
		if role == 0 {
			return
		}
		if !isValidRole(role) {
			fmt.Println("Pilihan tidak valid.")
			continue
		}

		for {
			printMenu(role)

			pilihan := readInt("Pilih menu: ")
			if pilihan == 0 {
				return
			}
			if shouldBackToRoleMenu(role, pilihan) {
				break
			}

			if role == 2 {
				mapped, ok := mapClientChoiceToAdminChoice(pilihan)
				if !ok {
					fmt.Println("Pilihan tidak valid.")
					continue
				}
				pilihan = mapped
			}

			switch pilihan {
			case 1:
				username := readLine("Username: ")
				name := readLine("Name: ")
				email := readLine("Email: ")
				password := readLine("Password: ")
				tambahUser(username, email, password, name)

			case 2:
				productCode := readLine("Kode Produk: ")
				name := readLine("Nama Produk: ")
				price := readFloat("Harga: ")
				stock := readInt("Stok: ")
				tambahProduk(productCode, name, price, stock)

			case 3:
				userName := strings.ToLower(readLine("Username: "))
				userID := getUserIDByUsernameOrPrint(userName)
				if userID == 0 {
					continue
				}
				buatOrder(userID)

			case 4:
				userName := strings.ToLower(readLine("Username: "))

				userID := getUserIDByUsernameOrPrint(userName)
				if userID == 0 {
					continue
				}

				productCode := strings.ToLower(readLine("Kode Produk: "))

				productID := getProductIDByCodeOrPrint(productCode)
				if productID == 0 {
					continue
				}
				
				qty := readInt("Qty: ")

				orderID, err := getLatestOrderIDByUserID(userID)
				if err != nil {
					orderID = buatOrder(userID)
					if orderID == 0 {
						fmt.Println("Gagal membuat order.")
						continue
					}
					fmt.Println("Tidak ada order sebelumnya. Membuat order baru:", orderID)
				}
				tambahItem(orderID, productID, qty)

			case 5:
				laporanUser()
				waitEnterToContinue()

			case 6:
				laporanOrder()
				waitEnterToContinue()

			case 7:
				laporanStok()
				waitEnterToContinue()

			case 8:
				username := strings.ToLower(readLine("Username: "))
				deleteUserByUsername(username)

			case 9:
				code := strings.ToLower(readLine("Kode Produk: "))
				deleteProductByCode(code)

			case 10:
				username := strings.ToLower(readLine("Username: "))
				address := readLine("Address: ")
				phone := readLine("Phone: ")
				updateUserContactByUsername(username, address, phone)

			default:
				fmt.Println("Pilihan tidak valid.")
			}
		}
	}
}