// package / folder utama untuk menjalankan apps
// jika file tidak berada dalam sebuah folder
package main

import (
	"fmt"

	"github.com/jutionck/golang-upskilling-agt/model"
)

// sebuah function utama untuk menjalankan apps
// ini harus ada, dan dalam 1 project/folder biasanya hanya ada 1 buah func main
func main() {

	var u model.User

	u.Id = "1"
	u.Username = "admin"
	u.Password = "password"
	u.Role = "aaaa"

	u.IsRole()

	fmt.Println("u:", u)

	model.IsRole()

	uu := model.User{
		Id:       "2",
		Username: "staff",
		Password: "password",
	}

	fmt.Println("uu:", uu)

	// println("Hello World")

	// var name string = "Jution"
	// fmt.Println(name)

	// namaSaya := "Jution"
	// println(namaSaya)

	// users := []string{"budi", "tono"}
	// fmt.Println(users)

	// greeting()
	// fmt.Println(greeting())

	// greet := greeting()
	// fmt.Println(greet)

	// _, b, _ := multipleReturn()
	// // fmt.Println("a:", a, "b:", b, "c:", c)

	// fmt.Println("b:", b)

	// // fmt.Println(multipleReturn())

	// // & => ampersan (bisanya itu dia mengambil dari variabel lain atau mencetak nilai memorinya)

	// name2 := &name
	// // var name3 *string

	// fmt.Printf("Alamat memori name: %p \n", &name)
	// fmt.Printf("Alamat memori name2: %p \n", name2) // ini mencetak alamat memori dari alamat memori

	// // pass by value & pass by reference
	// // tapi kalo mau melihat isi (value) dari sebuah pointer => *name2

}

// func greeting() string {
// 	return "Greeting"
// }

// func multipleReturn() (string, int, bool) {
// 	return "1", 0, true
// }

// 1. Slice -> kumpulan data (array) => [] => ditambahkan secara flexible (Go) "append" => (JS) => .push => (Java) => .add
//		var students []string
// 2. Pointer -> Tidak semua harus menggunakan POINTER, walaupun POINTER ini sangat PENTING
// 		=> Referensi alamat memori (*) => tipe data, struct *int, *string, diterapkan pada sebuah METHOD (hubungan sama struct)
// 3. Function -> function => func namaFunction(a struct) {}
// 4. Struct -> mirip seperti class tetapi bukan OOP (type ... struct)
// 5. Method => sebuah function di GO tetapi dia ada sebuah receiver -> func (a *receiver) namaFunction() {}
// 6. Interface => sebuah kontrak => interface biasanya mendifinisakan method (type ... interface)
