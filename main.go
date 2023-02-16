package main

import (
	"PROJECT1/database"
	"PROJECT1/user"
	"fmt"
)

func main() {

	database.Connect()

	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Read Account")
	fmt.Println("4. Updated Account")
	fmt.Println("5. Delete Account")

	var pilihan int
	fmt.Println("Masukkan pilihan")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		user.Register()
	} else if pilihan == 2 {
		user.Login(database.DB)
	} else if pilihan == 3 {
		user.ReadAccount(database.DB)
	} else if pilihan == 4 {
		user.Update(database.DB)
	}
}
