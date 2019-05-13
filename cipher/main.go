package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	plainText := []byte("Bob loves Alice.")

	// size of key (bits)
	size := 1024

	// Generate private and public key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	// Get public key from private key and encrypt
	publicKey := &privateKey.PublicKey

	// 第二引数がポインタ型なのは、この関数の第二引数がポインタを使って参照・操作しているので
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	fmt.Printf("Cipher text: %x\n", cipherText)

	// Decrypt with private key
	// 第二引数がポインタ型なのは(privateKeyもポインタ型で返ってくる)、この関数の第二引数がポインタを使って参照・操作しているので
	decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	fmt.Printf("Decrypted text: %s\n", decryptedText)
}
