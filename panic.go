package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error message: ", message)
	}
	fmt.Println("Aplikasi Selesai...")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Dika")
}
