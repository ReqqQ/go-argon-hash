package Input

import (
	"bufio"
	_ "bufio"
	"fmt"
	"os"
	_ "os"

	"hashPassword/Consts"
	HashStruct "hashPassword/GeneratePassword/Structs"
	"hashPassword/Utils"
	"hashPassword/Validation"
)

func GetSpecialSettings(option int) HashStruct.CustomOptions {
	if option == Consts.OptionVerify {
		return HashStruct.CreateCustomOptionStruct(Consts.SpecialSettingsNothing)
	}
	for _, question := range Consts.CustomOptionsQuestions {
		fmt.Println(question)
	}
	inputOption := Utils.ParseOptionToInt(getInput())
	Validation.ValidateOption(inputOption, Consts.CustomOptionsQuestions)

	return HashStruct.CreateCustomOptionStruct(inputOption)
}
func GetUserHash(option int, hash HashStruct.Hash) HashStruct.Hash {
	if option == Consts.OptionVerify {
		fmt.Println(Consts.TextPrintDbHash)
		hash = HashStruct.AddArgonTextToHash(hash, getInput())
	}

	return hash
}

func GetUserText() HashStruct.Hash {
	fmt.Println(Consts.TextPrintToHash)
	userText := getInput()

	return HashStruct.CreateHashStruct(userText)
}
func getInput() string {
	userText := bufio.NewScanner(os.Stdin)
	userText.Scan()

	return userText.Text()
}
func GetUserOption() int {
	for _, question := range Consts.MainQuestions {
		fmt.Println(question)
	}
	option := Utils.ParseOptionToInt(getInput())
	Validation.ValidateOption(option, Consts.MainQuestions)

	return option
}
