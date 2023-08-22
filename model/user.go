// harus sama seperti nama folder dia berada
package model

type UserModel interface {
	IsRole() bool
}

// kalo mau digunakan secara luas atau beda package: main, model, repo itu harus di awali dengan huruf besar (awalnya aja)
type User struct {
	Id       string
	Username string
	Password string
	Role     string
}

// method
// tidak bisa langsung di panggil harus melalui struct nya
// atau bisa lewat interface
func (u User) IsRole() bool {
	return u.Role == "admin" || u.Role == "staff"
}

func IsRole() {}
