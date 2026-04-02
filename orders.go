package main

import "fmt"

func buatOrder(userID int) int {
	result, err := db.Exec("INSERT INTO orders (user_id) VALUES (?)", userID)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	id, _ := result.LastInsertId()
	fmt.Println("Order dibuat dengan ID:", id)
	return int(id)
}

func tambahItem(orderID, productID, qty int) {
	query := "INSERT INTO order_items (order_id, product_id, quantity) VALUES (?, ?, ?)"
	_, err := db.Exec(query, orderID, productID, qty)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Item ditambahkan ke order")
}
