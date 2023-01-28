package DecodePassword

import (
	"fmt"

	"github.com/matthewhartstonge/argon2"
	"hashPassword/Consts"
	HashStruct "hashPassword/GeneratePassword/Structs"
	"hashPassword/Validation"
)

func Verify(hashStruct HashStruct.Hash) {
	decoded, err := argon2.Decode([]byte(hashStruct.ArgonText))
	Validation.Validate(err, Consts.ErrorOnDecodePass)
	verified, err := decoded.Verify([]byte(hashStruct.UserText))
	Validation.Validate(err, Consts.ErrorOnVerify)
	resultVerifiedPass(verified)
}
func resultVerifiedPass(verified bool) {
	if verified {
		fmt.Println(Consts.PasswordOKInfo)
	} else {
		fmt.Println(Consts.PasswordBADInfo)
	}
}
