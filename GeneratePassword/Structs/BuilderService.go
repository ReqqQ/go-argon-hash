package GeneratePassword

import (
	"hashPassword/Consts"
)

func BuildHashStruct(hash Hash, options CustomOptions) Hash {
	hash.CustomSetting = options

	return hash
}
func AddArgonTextToHash(hash Hash, hashText string) Hash {
	hash.ArgonText = hashText

	return hash
}
func prepareSpecialSettingsData(option int) (bool, bool) {
	salt, specialKey := false, false
	if option == Consts.SpecialSettingsNothing {
		return salt, specialKey
	}
	switch option {
	case Consts.SpecialSettingsSalt:
		salt = true
	case Consts.SpecialSettingsDecode:
		specialKey = true
	}

	return salt, specialKey
}
func CreateHashStruct(userText string) Hash {
	return Hash{UserText: userText}
}
func CreateCustomOptionStruct(option int) CustomOptions {
	salt, specialKey := prepareSpecialSettingsData(option)

	return CustomOptions{
		Salt:       salt,
		SpecialKey: specialKey,
	}
}
