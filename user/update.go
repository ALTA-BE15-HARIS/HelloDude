package user

import (
	"database/sql"
	"fmt"
)

func Update(db *sql.DB) {
	var NamaUser string
	var PasswordUser string
	var NoHpUser string
	fmt.Println("Nama yang akan di ubah :")
	fmt.Scanln(&NamaUser)
	fmt.Println("Password yang akan di ubah :")
	fmt.Scanln(&PasswordUser)
	fmt.Println("Nomor yang akan di ubah :")
	fmt.Scanln(&NoHpUser)
	_, err := db.
		Exec("update user set(name,password,no_hp) Where id (?, ?, ?)",
			NamaUser, PasswordUser, NoHpUser)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")
}
