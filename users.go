package main

import (
	"fmt"
	"strings"
)

func tambahUser(username, email, password, name string) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	result, err := db.Exec(query, username, email, password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = db.Exec("INSERT INTO profiles (user_id, name) VALUES (?, ?)", userID, name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User berhasil ditambahkan")
}

func deleteUserByUsername(username string) {
	username = strings.TrimSpace(username)

	result, err := db.Exec("DELETE FROM users WHERE LOWER(username) = LOWER(?)", username)
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
		fmt.Println("User tidak ditemukan.")
		return
	}

	fmt.Println("User berhasil dihapus")
}

func updateUserContactByUsername(username, address, phone string) {
	username = strings.TrimSpace(username)
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE LOWER(username) = LOWER(?) LIMIT 1", username).Scan(&userID)
	if err != nil {
		fmt.Println("User tidak ditemukan.")
		return
	}

	_, err = db.Exec(
		"INSERT INTO profiles (user_id, address, phone) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE address = VALUES(address), phone = VALUES(phone)",
		userID,
		address,
		phone,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User berhasil diupdate")
}
