package main

import (

 "github.com/pborman/uuid"
 "fmt"
 "github.com/priyanshu-space/golang/Modules_and_Packages/cryptit"

)

func main() {

	uuid := uuid.NewRandom()
	fmt.Println(uuid)

	encryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println(encryptedStr)

}