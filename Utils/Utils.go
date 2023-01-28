package Utils

import (
	"log"
	"strconv"

	"hashPassword/Consts"
)

func InArray(textToSearch string, currentList map[int]string) bool {
	exists := false
	for index, value := range currentList {
		if index != Consts.FirstQuestionId && value == textToSearch {
			exists = true
		}
	}
	return exists
}

func ParseOptionToInt(userOption string) int {
	userOptionConverted, err := strconv.Atoi(userOption)
	if err != nil {
		log.Fatal(err)
	}

	return userOptionConverted
}
