package user

import (
	"database/sql"
	"fmt"
)

func Login(db *sql.DB) {
	var PasswordUser string
	var NoHpUser string
	fmt.Println("Masukkan nomor Telepon anda :")
	fmt.Scanln(&NoHpUser)
	fmt.Println("Masukkan Password anda :")
	fmt.Scanln(&PasswordUser)

	var user User
	err := db.
		QueryRow("SELECT name FROM user WHERE no_hp = ? and password = ?",
			NoHpUser, PasswordUser).
		Scan(&user.Nama)
	if err != nil {
		fmt.Println("Maaf, kata sandi Anda salah. Harap periksa kembali kata sandi Anda.")
		return
	}
	fmt.Println("Berhasil masuk")
}
