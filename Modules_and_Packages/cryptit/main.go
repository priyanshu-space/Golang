package main

import (
	"fmt"
	"github.com/priyanshu-space/golang/Modules_and_Packages/cryptit/encrypt"
	"github.com/priyanshu-space/golang/Modules_and_Packages/cryptit/decrypt"
)

func main() {

	encryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println(encrypt.Nimbus("Kodekloud"))
	fmt.Println(decrypt.Nimbus(encryptedStr))
}