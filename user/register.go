package user

import (
	"PROJECT1/database"
	"fmt"
)

func Register() {

	var NamaUser string
	var PasswordUser string
	var NoHpUser string
	var SaldoUser int
	fmt.Println("Masukkan nama :")
	fmt.Scanln(&NamaUser)
	fmt.Println("Masukkan password :")
	fmt.Scanln(&PasswordUser)
	fmt.Println("Masukkan No Hp :")
	fmt.Scanln(&NoHpUser)
	fmt.Println("Saldo anda :")
	fmt.Scanln(&SaldoUser)

	_, err := database.DB.
		Exec("insert into user(name,password,no_hp,saldo) values (?, ?, ?, ?)",
			NamaUser, PasswordUser, NoHpUser, SaldoUser)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")
}
