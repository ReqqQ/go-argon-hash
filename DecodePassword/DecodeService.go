package DecodePassword

import (
	"fmt"
	"log"

	"github.com/matthewhartstonge/argon2"
	HashStruct "hashPassword/GeneratePassword/Structs"
)

func Verify(hashStruct HashStruct.Hash) {
	decoded, err := argon2.Decode([]byte(hashStruct.ArgonText))
	if err != nil {
		log.Fatal("Error on encoded")
	}
	verified, err := decoded.Verify([]byte(hashStruct.UserText))
	if err != nil {
		log.Fatal("Error on verify")
	}
	if verified {
		fmt.Println("Password OK")
	} else {
		fmt.Println("Password BAD")
	}
}
