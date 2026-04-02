package main

import (
	"fmt"
	"strings"
)

func tambahProduk(productCode, name string, price float64, stock int) {
	query := "INSERT INTO products (product_code, name, price, stock) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, productCode, name, price, stock)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Produk berhasil ditambahkan")
}

func deleteProductByCode(code string) {
	code = strings.TrimSpace(code)

	result, err := db.Exec("DELETE FROM products WHERE LOWER(product_code) = LOWER(?)", code)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if affected == 0 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}

	fmt.Println("Produk berhasil dihapus")
}
