package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	text, _ := in.ReadString('\n')
	return strings.TrimSpace(text)
}

func readInt(prompt string) int {
	for {
		s := readLine(prompt)
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err == nil {
			return n
		}
		fmt.Println("Input tidak valid. Masukkan angka.")
	}
}

func readFloat(prompt string) float64 {
	for {
		s := readLine(prompt)
		n, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err == nil {
			return n
		}
		fmt.Println("Input tidak valid. Masukkan angka.")
	}
}
