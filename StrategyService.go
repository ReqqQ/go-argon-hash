package main

import (
	"hashPassword/Consts"
	"hashPassword/DecodePassword"
	"hashPassword/GeneratePassword"
	HashStruct "hashPassword/GeneratePassword/Structs"
)

func chooseStrategy(userOption int, hashStruct HashStruct.Hash) {
	switch userOption {
	case Consts.OptionDecode:
		GeneratePassword.Generate(hashStruct)
	case Consts.OptionVerify:
		DecodePassword.Verify(hashStruct)
	}
}
