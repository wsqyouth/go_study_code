package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generatePassword(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#%^&*()_+{}:<>?|"
	var password string
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		password += string(charset[n.Int64()])
	}
	return password
}
func main() {
	password1 := generatePassword(12) // 生成长度为12的密码
	password2 := generatePassword(16) // 生成长度为16的密码
	fmt.Println("Password 1:", password1)
	fmt.Println("Password 2:", password2)
}
