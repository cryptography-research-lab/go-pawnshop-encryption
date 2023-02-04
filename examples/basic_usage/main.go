package main

import (
	"fmt"
	pawnshop_encryption "github.com/cryptography-research-lab/go-pawnshop-encryption"
)

func main() {

	// 对笔者的QQ号进行加密
	plaintext := "1451546085"
	encrypt, err := pawnshop_encryption.Encrypt(plaintext)
	if err != nil {
		fmt.Println("加密失败： " + err.Error())
		return
	}
	fmt.Println("加密结果： " + encrypt) // Output: 加密结果： 力训朴八向汀钉一姣亦

	decrypt, err := pawnshop_encryption.Decrypt(encrypt)
	if err != nil {
		fmt.Println("解密失败： " + err.Error())
		return
	}
	fmt.Println("解密结果： " + decrypt) // Output: 解密结果： 1451546085

}
