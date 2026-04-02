package main

import "strings"

func getUserIDByUsername(username string) (int, error) {
	username = strings.TrimSpace(username)
	var userID int
	err := db.QueryRow(
		"SELECT id FROM users WHERE LOWER(username) = LOWER(?) LIMIT 1",
		username,
	).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func getProductIDByCode(code string) (int, error) {
	code = strings.TrimSpace(code)
	var productID int
	err := db.QueryRow(
		"SELECT id FROM products WHERE LOWER(product_code) = LOWER(?) LIMIT 1",
		code,
	).Scan(&productID)
	if err != nil {
		return 0, err
	}
	return productID, nil
}

func getLatestOrderIDByUserID(userID int) (int, error) {
	var orderID int
	err := db.QueryRow("SELECT id FROM orders WHERE user_id = ? ORDER BY id DESC LIMIT 1", userID).Scan(&orderID)
	if err != nil {
		return 0, err
	}
	return orderID, nil
}
