package pawnshop_encryption

import (
	"math/rand"
	"unicode/utf8"
)

// Encrypt 对数字字符串进行当铺加密
func Encrypt(plaintext string) (string, error) {
	runeSlice := make([]rune, utf8.RuneCountInString(plaintext))
	for index, numberCharacter := range plaintext {
		// 要加密的明文的每一个字符都必须是数字
		if numberCharacter < '0' || numberCharacter > '9' {
			return "", ErrPlaintextUnavailable
		}
		number := int(numberCharacter-'0') + 1
		runes := encryptDictionary[number]
		// 将对应笔画的汉字随机选择一个
		runeSlice[index] = runes[rand.Intn(len(runes))]
	}
	return string(runeSlice), nil
}

// Decrypt 对当铺密码加密的密文进行解密
func Decrypt(ciphertext string) (string, error) {
	runeSlice := make([]rune, utf8.RuneCountInString(ciphertext))
	for index, character := range []rune(ciphertext) {
		number, exists := decryptDictionary[character]
		if !exists {
			return "", ErrCiphertextUnavailable
		}
		runeSlice[index] = rune(number - 1 + '0')
	}
	return string(runeSlice), nil
}
