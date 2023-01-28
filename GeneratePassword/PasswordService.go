package GeneratePassword

import (
	"fmt"
	"log"

	"github.com/matthewhartstonge/argon2"
	"github.com/thanhpk/randstr"
	HashStruct "hashPassword/GeneratePassword/Structs"
)

const PassSalt = "CookIeSaltBefore"
const RandomStringLength = 100

func Generate(hashStruct HashStruct.Hash) {
	argon := argon2.DefaultConfig()
	AddSalt(&hashStruct)
	AddRandomText(&hashStruct)
	encoded, err := argon.Hash([]byte(hashStruct.UserText), nil)

	if err != nil {
		log.Fatal("Error on encoding pass")
	}

	fmt.Println("Your hashed text is: " + string(encoded.Encode()))
}
func AddSalt(hashStruct *HashStruct.Hash) *HashStruct.Hash {
	if hashStruct.CustomSetting.Salt {
		hashStruct.UserText += PassSalt
	}
	return hashStruct
}
func AddRandomText(hashStruct *HashStruct.Hash) *HashStruct.Hash {
	if hashStruct.CustomSetting.SpecialKey {
		randomString := randstr.Hex(RandomStringLength)
		fmt.Println("Key to decode: " + randomString)
		hashStruct.UserText += randomString
	}

	return hashStruct
}
