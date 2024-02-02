package pawnshop_encryption

func Encrypt(numbers string) (string, error) {
	return DefaultDictionary.Encrypt(numbers)
}

func Decrypt(chineseString string) (string, error) {
	return DefaultDictionary.Decrypt(chineseString)
}
