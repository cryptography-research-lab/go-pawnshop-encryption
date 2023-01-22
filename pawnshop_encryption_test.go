package pawnshop_encryption

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plaintext := "012345678987654321"
	encrypt, err := Encrypt(plaintext)
	assert.Nil(t, err)
	decryptResult, err := Decrypt(encrypt)
	assert.Nil(t, err)
	assert.Equal(t, plaintext, decryptResult)
}
