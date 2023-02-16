package user

import (
	"database/sql"
	"fmt"
)

func ReadAccount(db *sql.DB) {
	var PasswordUser string
	var NoHpUser string
	fmt.Println("Masukkan nomor Telepon anda :")
	fmt.Scanln(&NoHpUser)
	fmt.Println("Masukkan Password anda :")
	fmt.Scanln(&PasswordUser)

	var user User
	err := db.
		QueryRow("SELECT name, no_hp, saldo FROM user WHERE no_hp = ? and password = ?",
			NoHpUser, PasswordUser).
		Scan(&user.Nama, &user.NoHp, &user.Saldo)
	if err != nil {
		fmt.Println("Akun tidak di temukan")
		return
	}
	fmt.Println("Nama :", user.Nama)
	fmt.Println("No HP :", user.NoHp)
	fmt.Println("Isi saldo :", user.Saldo)
}
