package pawnshop_encryption

import "errors"

var (

	// ErrPlaintextUnavailable 传入的明文不可用
	ErrPlaintextUnavailable = errors.New("plaintext unavailable, plaintext all character must number")

	// ErrCiphertextUnavailable 要解密的密文在字典中没有发现
	ErrCiphertextUnavailable = errors.New("ciphertext unavailable, ciphertext character not found in dictionary")
)
