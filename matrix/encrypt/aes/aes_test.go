package aes

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	plainText := ""
	securityKey16 := "ca6590a271539cc89e2cc20bd6b58518"
	iv := "1234567890123456"
	aes := aesTool(securityKey16, iv)
	cipherText, _ := aes.encrypt(plainText)
	fmt.Println("加密后的密文：" + cipherText)
	outPlainText, _ := aes.decrypt(cipherText)
	fmt.Println("解密后明文：" + outPlainText)

}

//707a161834ea0e5a3f1cc419a37bc030
//cca690a271539cc89e2cc20bd6b58518