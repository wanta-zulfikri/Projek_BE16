package main

import (
	"Projek_BE16/config"
	"Projek_BE16/user"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn := config.InitSQL()
	mdl := user.UserModel{}
	mdl.SetSQLConnection(conn)
	user_login := user.User{}
	var menu int
	for menu != 99 {
		fmt.Println("TODOLITS AYO BUAT CATATAN UNTUK KEGIATAN MU")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Update")
		fmt.Println("4. Deactive Akun ")
		fmt.Print("Masukkan Pilihan : ")
		fmt.Scanln(&menu)
		switch menu {
		case 1:
			if user_login == (user.User{}) {
				fmt.Println("Register")
				fmt.Print("Nama \t\t:  ")
				fmt.Scanln(&user_login.Nama)
				fmt.Print("No. Telp\t:  ")
				fmt.Scanln(&user_login.No_HP)
				fmt.Print("Password \t: ")
				fmt.Scanln(&user_login.Password)
				if user_login == (user.User{}) {
					log.Fatal("Gagal mendaftar. Mohon isi semua data yang dibutuhkan.")
				} else {
					err := mdl.Create(&user_login)
					if err != nil {
						log.Fatal(err)
					} else {
						fmt.Println("Registrasi berhasil!")
					}
				}
			} else {
				fmt.Println("Login gagal. Anda sudah terdaftar. Silakan login atau deaktifkan akun jika ingin membuat akun baru.")
			}
		case 2:
			if user_login == (user.User{}) {
				fmt.Println("Login")
				fmt.Print("No. Telp\t:  ")
				fmt.Scanln(&user_login.No_HP)
				fmt.Print("Password \t: ")
				fmt.Scanln(&user_login.Password)
				usr, err := mdl.GetByNoHP(user_login.No_HP)
				if err != nil {
					log.Fatal(err)
				}
				if usr.Password != user_login.Password {
					fmt.Println("Login gagal. Password yang anda masukkan salah.")
				} else {
					fmt.Println("Login berhasil!")
				}
			} else {
				fmt.Println("Login gagal. Anda sudah login dengan akun", user_login.Nama)
			}
		}
	}
}
