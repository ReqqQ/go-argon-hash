package Validation

import (
	"log"

	"hashPassword/Consts"
	"hashPassword/Utils"
)

func ValidateOption(option int, questions map[int]string) {
	if option == Consts.FirstQuestionId && !Utils.InArray(string(rune(option)), questions) {
		log.Fatal(Consts.OptionNotExist)
	}
}
func Validate(err error, errorName string) {
	if err != nil {
		log.Fatal(Consts.ErrorOnVerify)
	}
}
