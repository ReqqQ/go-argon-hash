package GeneratePassword

import (
	"fmt"

	"github.com/matthewhartstonge/argon2"
	"github.com/thanhpk/randstr"
	"hashPassword/Consts"
	HashStruct "hashPassword/GeneratePassword/Structs"
	"hashPassword/Validation"
)

func Generate(hashStruct HashStruct.Hash) {
	argon := argon2.DefaultConfig()
	AddSalt(&hashStruct)
	AddRandomText(&hashStruct)
	encoded, err := argon.Hash([]byte(hashStruct.UserText), nil)
	Validation.Validate(err, Consts.ErrorOnDecodePass)
	fmt.Println(Consts.PasswordResult + string(encoded.Encode()))
}

func AddSalt(hashStruct *HashStruct.Hash) *HashStruct.Hash {
	if hashStruct.CustomSetting.Salt {
		hashStruct.UserText += Consts.PassSalt
	}
	return hashStruct
}

func AddRandomText(hashStruct *HashStruct.Hash) *HashStruct.Hash {
	if hashStruct.CustomSetting.SpecialKey {
		randomString := randstr.Hex(Consts.RandomStringLength)
		fmt.Println(Consts.PasswordSpecialKeyToDecode + randomString)
		hashStruct.UserText += randomString
	}

	return hashStruct
}
