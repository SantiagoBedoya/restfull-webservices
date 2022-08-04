package main

import (
	"log"

	"github.com/SantiagoBedoya/restfull-webservices/chapter9/encryptString/utils"
)

func main() {
	key := "111023043350789514532147"
	message := "I'm a message"
	log.Println("original message: ", message)
	encryptedString := utils.EncryptString(key, message)
	log.Println("encrypted message: ", encryptedString)
	descryptedString := utils.DecryptString(key, encryptedString)
	log.Println("decrypted message: ", descryptedString)
}
