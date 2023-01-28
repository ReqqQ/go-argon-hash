package main

import (
	GeneratePassword "hashPassword/GeneratePassword/Structs"
	"hashPassword/Input"
)

func main() {
	option := Input.GetUserOption()
	hashStruct := Input.GetUserText()
	hashStruct = Input.GetUserHash(option, hashStruct)
	customOptionsStruct := Input.GetSpecialSettings(option)
	chooseStrategy(option, GeneratePassword.BuildHashStruct(hashStruct, customOptionsStruct))
}
