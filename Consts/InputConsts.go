package Consts

const (
	SpecialSettingsSalt    = 1
	SpecialSettingsDecode  = 2
	SpecialSettingsNothing = 3
	OptionDecode           = 1
	OptionVerify           = 2
	FirstQuestionId        = 0
	TextPrintToHash        = "Write your text to hash / check"
	TextPrintDbHash        = "Write your hashed password"
	OptionNotExist         = "Option not exists"
)

var CustomOptionsQuestions = map[int]string{
	0: "Do you need another options?",
	1: "[1] Add salt",
	2: "[2] Key to decode / encode",
	3: "[3] No",
}

var MainQuestions = map[int]string{
	0: "Choose what do you want to do:",
	1: "[1] Generate password and hash",
	2: "[2] Verify password",
}
